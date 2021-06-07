package db

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"sync"
	"time"

	"git.neds.sh/matty/entain/sports/proto/sports"
)

// SportsRepo provides repository access to sports.
type SportsRepo interface {
	// Init will initialise our sports repository.
	Init() error

	// List will return a list of sports events.
	List() ([]*sports.SportEvent, error)
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepo creates a new sports repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the sports repository dummy data.
func (r *sportsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy sports.
		err = r.seed()
	})

	return err
}

func (r *sportsRepo) List() ([]*sports.SportEvent, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportEventsQueries()[eventList]

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanSportEvents(rows)
}

func (m *sportsRepo) scanSportEvents(
	rows *sql.Rows,
) ([]*sports.SportEvent, error) {
	var sportEvents []*sports.SportEvent

	for rows.Next() {
		var sportEvent sports.SportEvent
		var advertisedStart time.Time

		if err := rows.Scan(&sportEvent.Id, &sportEvent.Name, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sportEvent.AdvertisedStartTime = ts

		sportEvents = append(sportEvents, &sportEvent)
	}

	return sportEvents, nil
}
