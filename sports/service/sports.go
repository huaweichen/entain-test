package service

import (
	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"golang.org/x/net/context"
)

type Sport interface {
	// ListEvents will return a collection of sports events.
	ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsResponse, error)
}

// sportsService implements the sports interface.
type sportService struct {
	sportRepo db.SportsRepo
}

// NewSportsService instantiates and returns a new sportsService.
func NewSportsService(sportRepo db.SportsRepo) Sport {
	return &sportService{sportRepo}
}

func (s *sportService) ListEvents(ctx context.Context, in *sports.ListEventsRequest) (*sports.ListEventsResponse, error) {
	sportEvents, err := s.sportRepo.List()
	if err != nil {
		return nil, err
	}

	return &sports.ListEventsResponse{SportEvents: sportEvents}, nil
}
