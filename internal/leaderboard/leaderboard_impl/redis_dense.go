package leaderboard_impl

import (
	"context"
	"errors"
	"fmt"
	"game_leaderboard/internal/leaderboard"
	"github.com/go-redis/redis/v8"
	"strconv"
)

const (
	playerScoreKey     = "player:score"
	playerScoreListKey = "player:score:%d"
	defaultPlayerId    = "score:player:%d" // 没有实际用处,就是用来计算排名
)

type RedisDenseLeaderboard struct {
	key    string
	client *redis.Client
}

func NewRedisDenseLeaderboard(redisAddr, key string) *RedisDenseLeaderboard {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &RedisDenseLeaderboard{client: rdb, key: key}
}

func (r *RedisDenseLeaderboard) UpdateScore(ctx context.Context, playerId string, score float64) error {
	if playerId == "" {
		return errors.New("playerId is empty")
	}

	// 开始一个事务
	pipe := r.client.TxPipeline()

	// 获取当前得分
	currentScore, err := r.client.HGet(ctx, playerScoreKey, playerId).Float64()
	if err == redis.Nil {
		currentScore = 0
	} else if err != nil {
		return err
	}

	// 更新哈希表中的得分
	pipe.HSet(ctx, playerScoreKey, playerId, score)

	// 如果分数发生变化
	if currentScore != score {
		// 从列表中移除旧分数的用户
		if currentScore > 0 {
			pipe.LRem(ctx, fmt.Sprintf(playerScoreListKey, int64(currentScore)), 0, playerId)
		}
		// 将用户添加到新分数列表中
		pipe.RPush(ctx, fmt.Sprintf(playerScoreListKey, int64(score)), playerId)
	}

	// 检查分数是否已存在于有序集合中
	members, err := r.client.ZRangeByScore(ctx, r.key, &redis.ZRangeBy{Min: fmt.Sprintf("%f", score), Max: fmt.Sprintf("%f", score)}).Result()
	if err != nil {
		return err
	}

	if len(members) > 0 {
		return nil
	}

	// 将新得分加入有序集合
	pipe.ZAdd(ctx, r.key, &redis.Z{Score: score, Member: fmt.Sprintf(defaultPlayerId, int64(score))})

	// 执行事务
	_, err = pipe.Exec(ctx)
	return err
}

func (r *RedisDenseLeaderboard) GetPlayerRank(ctx context.Context, playerId string) (*leaderboard.RankInfo, error) {
	if playerId == "" {
		return nil, errors.New("playerId is empty")
	}

	// 获取分数
	score, err := r.client.HGet(ctx, playerScoreKey, playerId).Float64()
	if err != nil {
		return nil, err
	}

	// 计算用户的排名
	rank := r.client.ZCount(ctx, r.key, strconv.FormatInt(int64(score), 10), "+inf").Val()

	return &leaderboard.RankInfo{PlayerId: playerId, Score: int64(score), Rank: int32(rank)}, nil
}

func (r *RedisDenseLeaderboard) GetTopN(ctx context.Context, n int32) ([]*leaderboard.RankInfo, error) {
	if n <= 0 {
		return nil, errors.New("arg errors")
	}
	members, err := r.client.ZRevRangeWithScores(ctx, r.key, 0, int64(n)-1).Result()
	if err != nil {
		return nil, err
	}

	var ranks []*leaderboard.RankInfo
	for i, member := range members {
		playerIds, err := r.client.LRange(ctx, fmt.Sprintf(playerScoreListKey, int64(member.Score)), 0, -1).Result()
		if err != nil {
			continue
		}
		for _, playerId := range playerIds {
			ranks = append(ranks, &leaderboard.RankInfo{PlayerId: playerId, Score: int64(member.Score), Rank: int32(i + 1)})
		}
	}

	return ranks, nil
}

func (r *RedisDenseLeaderboard) GetPlayerRankRange(ctx context.Context, playerId string, rangeSize int32) ([]*leaderboard.RankInfo, error) {
	if playerId == "" {
		return nil, errors.New("playerId is empty")
	}

	// 获取分数
	score, err := r.client.HGet(ctx, playerScoreKey, playerId).Float64()
	if err != nil {
		return nil, err
	}

	// 计算用户的排名
	rank := r.client.ZCount(ctx, r.key, strconv.FormatInt(int64(score), 10), "+inf").Val()

	start := rank - int64(rangeSize)
	end := rank + int64(rangeSize) - 1

	if start < 1 {
		start = 0
	} else {
		start -= 1
	}

	members, err := r.client.ZRevRangeWithScores(ctx, r.key, start, end).Result()
	if err != nil {
		return nil, err
	}

	var ranks []*leaderboard.RankInfo
	var memberRank int32
	for i, member := range members {
		playerIds, err := r.client.LRange(ctx, fmt.Sprintf(playerScoreListKey, int64(member.Score)), 0, -1).Result()
		if err != nil {
			continue
		}
		if start == 0 {
			memberRank = int32(rank) + int32(i+1) - int32(rank-start)
		} else {
			memberRank = int32(rank) - int32(i+1) - int32(rank-start)
		}
		for _, playerId := range playerIds {
			ranks = append(ranks, &leaderboard.RankInfo{PlayerId: playerId, Score: int64(member.Score), Rank: memberRank})
		}
	}

	return ranks, nil
}
