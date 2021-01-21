package history

import (
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
)

type HistoryService struct {
	db *pgx.ConnPool
}

func NewRepository(db *pgx.ConnPool) HistoryRepository {
	return HistoryService{
		db: db,
	}
}

func (hr HistoryService) All() ([]HistoryModel, error) {
	rows, err := hr.db.Query(`select * from history`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var historyList []HistoryModel

	for rows.Next() {
		var history HistoryModel
		err := rows.Scan(&history.ID, &history.Startdate, &history.Enddate, &history.Area)
		if err != nil {
			log.Fatal(err)
		}

		historyList = append(historyList, HistoryModel{
			ID:        history.ID,
			Startdate: history.Startdate,
			Enddate:   history.Enddate,
			Area:      history.Area,
			// TimeWateredInSeconds: time.Duration(history.Enddate.Sub(history.Startdate).Seconds())})
		})
	}
	return historyList, nil
}

func (hr HistoryService) Insert(startDate time.Time, endDate time.Time, area string) (*uuid.UUID, error) {
	id, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	_, err = hr.db.Exec("insert into history values($1, $2, $3, $4)", id, startDate, endDate, area)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
