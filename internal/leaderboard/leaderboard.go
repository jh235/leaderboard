package leaderboard

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

const LeaderboardStartTime = 1735660800 // 2025-01-01 00:00:00 时间

type Leaderboard interface {
	UpdateScore(ctx context.Context, playerId string, score float64) error
	GetPlayerRank(ctx context.Context, playerId string) (*RankInfo, error)
	GetTopN(ctx context.Context, n int32) ([]*RankInfo, error)
	GetPlayerRankRange(ctx context.Context, playerId string, rangeSize int32) ([]*RankInfo, error)
}

type LeaderboardService struct {
	db Leaderboard
	UnimplementedLeaderboardServiceServer
}

func NewLeaderboardService(db Leaderboard) *LeaderboardService {
	return &LeaderboardService{db: db}
}

// CalculateScoreByTimestamp 根据时间计算分数
func (l LeaderboardService) CalculateScoreByTimestamp(score, timestamp int64) float64 {
	if timestamp == 0 {
		return float64(score)
	}
	return float64(score) + float64(LeaderboardStartTime)/float64(timestamp)
}

func (l LeaderboardService) UpdateScore(ctx context.Context, request *UpdateScoreRequest) (*emptypb.Empty, error) {
	err := l.db.UpdateScore(ctx, request.PlayerId, l.CalculateScoreByTimestamp(request.Score, request.Timestamp))
	return nil, err
}

func (l LeaderboardService) GetPlayerRank(ctx context.Context, request *GetPlayerRankRequest) (*RankInfo, error) {
	return l.db.GetPlayerRank(ctx, request.PlayerId)
}

func (l LeaderboardService) GetTopN(ctx context.Context, request *GetTopNRequest) (*PlayerRankRangeResponse, error) {
	rank, err := l.db.GetTopN(ctx, request.N)
	return &PlayerRankRangeResponse{Ranks: rank}, err
}

func (l LeaderboardService) GetPlayerRankRange(ctx context.Context, request *GetPlayerRankRangeRequest) (*PlayerRankRangeResponse, error) {
	rankInfo, err := l.db.GetPlayerRankRange(ctx, request.PlayerId, request.Range)
	return &PlayerRankRangeResponse{Ranks: rankInfo}, err
}
