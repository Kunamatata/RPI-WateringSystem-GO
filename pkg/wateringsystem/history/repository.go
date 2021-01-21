package history

import (
	"time"

	"github.com/gofrs/uuid"
)

// HistoryModel represent the history model
type HistoryModel struct {
	ID        string
	Startdate time.Time
	Enddate   time.Time
	Area      string
}

// HistoryRepository represents the repository for the history
type HistoryRepository interface {
	All() ([]HistoryModel, error)
	Insert(startDate time.Time, endDate time.Time, area string) (*uuid.UUID, error)
}
