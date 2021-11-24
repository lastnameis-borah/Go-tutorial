package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("02-01-2006 Monday 15:01:05"))

	createDate := time.Date(2020, time.January, 29, 6, 59, 0, 0, time.UTC)
	// Month and Location is time.Month and time.Location, others are int type
	// (year, month, date, hour, minute, second, nanosecond, location)
	fmt.Println(createDate.Format("02-01-2006 Monday 15:01:05"))
}
