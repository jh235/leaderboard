package leaderboard_impl

import (
	"context"
	"game_leaderboard/internal/leaderboard"
)

type SkipListLeaderboard struct {
}

func NewSkipListLeaderboard() *SkipListLeaderboard {
	return &SkipListLeaderboard{}
}

func (r *SkipListLeaderboard) UpdateScore(ctx context.Context, playerId string, score float64) error {
	//TODO implement me
	panic("implement me")
}

func (r *SkipListLeaderboard) GetPlayerRank(ctx context.Context, playerId string) (*leaderboard.RankInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (r *SkipListLeaderboard) GetTopN(ctx context.Context, n int32) ([]*leaderboard.RankInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (r *SkipListLeaderboard) GetPlayerRankRange(ctx context.Context, playerId string, rangeSize int32) ([]*leaderboard.RankInfo, error) {
	//TODO implement me
	panic("implement me")
}
