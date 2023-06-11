package main

import (
	"fmt"
	"time"
)

func calc(x int, y int) int {
	return x + y
}

func main() {
	calc(1, 2)

	t := calc3(time.Now())
	fmt.Println(t)

	//時刻型で21:00を表す
	t1 := time.Date(2023, 12, 1, 21, 0, 0, 0, time.Local)
	fmt.Println(calc3(t1))

}

func calc2() (time.Time, time.Time) {
	now := time.Now()
	return now, now.Add(time.Hour)
}

func calc3(t time.Time) time.Time {
	return t.Add(time.Hour * 3)
}
