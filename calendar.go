package calendar

import (
	"math"
	"strconv"
	"time"
)

type Calendar struct {
	Year      int
	Month     int
	Today     time.Time
	StartDay  int
	TotalDays int
	Rows      int
	Cols      int
}

func Create(year int, month int) Calendar {
	var c = Calendar{}

	c.Today = time.Now()

	if year == 0 {
		c.Year = time.Now().Year()
	}

	if month == 0 {
		c.Month = int(time.Now().Month())
	}

	date := time.Date(c.Year, time.Month(c.Month), 1, 0, 0, 0, 0, time.UTC)

	c.StartDay = int(date.Weekday())
	c.TotalDays = daysInMonth(&c.Year, &c.Month)

	c.Rows = int(math.Ceil(float64(c.TotalDays+c.StartDay) / 7))
	c.Cols = 7

	return c
}

func daysInMonth(y *int, m *int) int {

	// moth validation
	if *m < 1 || *m > 12 {
		*m = int(time.Now().Month())
	}
	// year validation
	if *y <= 1970 && len(strconv.Itoa(*y)) != 4 {
		*y = time.Now().Year()
	}
	// Leap year check
	if *m == 2 {
		if *y%400 == 0 || (*y%4 == 0 && *y%100 != 0) {
			return 29
		}
	}
	daysInMonth := [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	return daysInMonth[*m-1]
}

func GetHead() [7]string {
	head := [7]string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
	}

	return head
}

func (cal Calendar) GetBody() [][]int {
	counter := 1
	table := make([][]int, cal.Rows)

	for i := 0; i < cal.Rows; i++ {
		table[i] = make([]int, cal.Cols)

		for j := 0; j < cal.Cols; j++ {
			if i == 0 && j < cal.StartDay-1 || counter > cal.TotalDays {
				table[i][j] = 0
			} else {
				table[i][j] = counter
				counter++
			}
		}
	}

	return table
}
