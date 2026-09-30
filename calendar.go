package calendar

import (
	"time"
)

type Calendar struct {
	Year      int
	Month     int
	Today     time.Time
	StartDay  int // 0 = Mon, 1 = Tue, ..., 6 = Sun
	TotalDays int
	Rows      int
	Cols      int
}

// Create builds a new Calendar. Passing 0 for year or month defaults to the current date.
func Create(year, month int) Calendar {
	now := time.Now()

	if year <= 0 {
		year = now.Year()
	}
	if month < 1 || month > 12 {
		month = int(now.Month())
	}

	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	// Shift Go's Sunday (0) to 6, and Monday (1) to 0
	startOffset := (int(firstOfMonth.Weekday()) + 6) % 7

	// Passing day 0 of month+1 yields the last day of month
	totalDays := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()

	cols := 7
	rows := (totalDays + startOffset + cols - 1) / cols

	return Calendar{
		Year:      year,
		Month:     month,
		Today:     now,
		StartDay:  startOffset,
		TotalDays: totalDays,
		Rows:      rows,
		Cols:      cols,
	}
}

func Head() [7]string {
	return [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
}

func (cal Calendar) Body() [][]int {
	dayCounter := 1
	table := make([][]int, cal.Rows)

	for i := 0; i < cal.Rows; i++ {
		table[i] = make([]int, cal.Cols)

		for j := 0; j < cal.Cols; j++ {
			if (i == 0 && j < cal.StartDay) || dayCounter > cal.TotalDays {
				table[i][j] = 0
			} else {
				table[i][j] = dayCounter
				dayCounter++
			}
		}
	}

	return table
}
