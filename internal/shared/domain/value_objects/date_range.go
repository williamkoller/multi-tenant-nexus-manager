package value_objects

import (
	"fmt"
	"time"
)

// DateRange - Value Object para período de datas
type DateRange struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func NewDateRange(start, end time.Time) (DateRange, error) {
	if start.After(end) {
		return DateRange{}, fmt.Errorf("start date cannot be after end date")
	}

	return DateRange{
		StartDate: start,
		EndDate:   end,
	}, nil
}

func (dr DateRange) Duration() time.Duration {
	return dr.EndDate.Sub(dr.StartDate)
}

func (dr DateRange) DurationInDays() int {
	return int(dr.Duration().Hours() / 24)
}

func (dr DateRange) DurationInMonths() int {
	years := dr.EndDate.Year() - dr.StartDate.Year()
	months := int(dr.EndDate.Month()) - int(dr.StartDate.Month())
	return years*12 + months
}

func (dr DateRange) Contains(date time.Time) bool {
	return (date.Equal(dr.StartDate) || date.After(dr.StartDate)) &&
		(date.Equal(dr.EndDate) || date.Before(dr.EndDate))
}

func (dr DateRange) Overlaps(other DateRange) bool {
	return dr.StartDate.Before(other.EndDate) && dr.EndDate.After(other.StartDate)
}

func (dr DateRange) IsValid() bool {
	return !dr.StartDate.After(dr.EndDate)
}

func (dr DateRange) String() string {
	return fmt.Sprintf("%s to %s",
		dr.StartDate.Format("2006-01-02"),
		dr.EndDate.Format("2006-01-02"))
}
