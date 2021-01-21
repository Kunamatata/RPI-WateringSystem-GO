package wateringsystem

import (
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
)

type History struct {
	Id                   string        `json:"id"`
	Startdate            time.Time     `json:"startDate"`
	Enddate              time.Time     `json:"endDate"`
	Area                 string        `json:"area"`
	TimeWateredInSeconds time.Duration `json:"timeWatered"`
}

type HistoryModel struct {
	Id        string
	Startdate time.Time
	Enddate   time.Time
	Area      string
}

// type HistoryRepository interface {
// 	All() []History
// 	Insert() error
// 	// GetLatest() History
// 	// SaveHistory()
// }

type HistoryService struct {
	db *pgx.ConnPool
}

func NewHistoryService(db *pgx.ConnPool) *HistoryService {
	return &HistoryService{
		db: db,
	}
}

func (hr HistoryService) All() ([]History, error) {
	rows, err := hr.db.Query(`select * from history`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var historyList []History

	for rows.Next() {
		var history HistoryModel
		err := rows.Scan(&history.Id, &history.Startdate, &history.Enddate, &history.Area)
		if err != nil {
			log.Fatal(err)
		}

		historyList = append(historyList, History{
			Id:                   history.Id,
			Startdate:            history.Startdate,
			Enddate:              history.Enddate,
			Area:                 history.Area,
			TimeWateredInSeconds: time.Duration(history.Enddate.Sub(history.Startdate).Seconds())})
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
