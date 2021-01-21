package history

import "time"

//History represent the structure of a History object
type History struct {
	ID                   string        `json:"id"`
	Startdate            time.Time     `json:"startDate"`
	Enddate              time.Time     `json:"endDate"`
	Area                 string        `json:"area"`
	TimeWateredInSeconds time.Duration `json:"timeWatered"`
}
