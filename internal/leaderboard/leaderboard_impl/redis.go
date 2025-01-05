package leaderboard_impl

import (
	"context"
	"errors"
	"game_leaderboard/internal/leaderboard"
	"github.com/go-redis/redis/v8"
)

type RedisLeaderboard struct {
	key    string
	client *redis.Client
}

func NewRedisLeaderboard(redisAddr, key string) *RedisLeaderboard {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &RedisLeaderboard{client: rdb, key: key}
}

func (r *RedisLeaderboard) UpdateScore(ctx context.Context, playerId string, score float64) error {
	if playerId == "" {
		return errors.New("playerId is empty")
	}
	_, err := r.client.ZAdd(ctx, r.key, &redis.Z{
		Score:  score,
		Member: playerId,
	}).Result()
	return err
}

func (r *RedisLeaderboard) GetPlayerRank(ctx context.Context, playerId string) (*leaderboard.RankInfo, error) {
	if playerId == "" {
		return nil, errors.New("playerId is empty")
	}
	rank, err := r.client.ZRevRank(ctx, r.key, playerId).Result()
	if err != nil {
		return nil, err
	}

	score, err := r.client.ZScore(ctx, r.key, playerId).Result()
	if err != nil {
		return nil, err
	}

	return &leaderboard.RankInfo{PlayerId: playerId, Score: int64(score), Rank: int32(rank) + 1}, nil
}

func (r *RedisLeaderboard) GetTopN(ctx context.Context, n int32) ([]*leaderboard.RankInfo, error) {
	if n <= 0 {
		return nil, errors.New("arg errors")
	}
	members, err := r.client.ZRevRangeWithScores(ctx, r.key, 0, int64(n)-1).Result()
	if err != nil {
		return nil, err
	}

	var ranks []*leaderboard.RankInfo
	for i, member := range members {
		ranks = append(ranks, &leaderboard.RankInfo{PlayerId: member.Member.(string), Score: int64(member.Score), Rank: int32(i) + 1})
	}

	return ranks, nil
}

func (r *RedisLeaderboard) GetPlayerRankRange(ctx context.Context, playerId string, rangeSize int32) ([]*leaderboard.RankInfo, error) {
	if playerId == "" {
		return nil, errors.New("playerId is empty")
	}
	rank, err := r.client.ZRevRank(ctx, r.key, playerId).Result()
	if err != nil {
		return nil, err
	}
	start := int64(rank) - int64(rangeSize)
	end := int64(rank) + int64(rangeSize)

	if start < 0 {
		start = 0
	}

	members, err := r.client.ZRevRangeWithScores(ctx, r.key, start, end).Result()
	if err != nil {
		return nil, err
	}

	var ranks []*leaderboard.RankInfo
	for i, member := range members {
		ranks = append(ranks, &leaderboard.RankInfo{PlayerId: member.Member.(string), Score: int64(member.Score), Rank: int32(i) + 1})
	}

	return ranks, nil
}
