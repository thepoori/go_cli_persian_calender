package main

import (
	"fmt"
	"time"

	"github.com/wsxiaoys/terminal/color"
)

var cprint = color.Println
var not_leap = [12]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
var leap = [12]int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}
var Weekdays = map[int]string{
	1: "Sa",
	2: "Su",
	3: "Mo",
	4: "Tu",
	5: "We",
	6: "Th",
	7: "Fr",
}
var persian_month = map[int]string{
	1: "farvardin", 2: "ordibehesht", 3: "khordad", 4: "tir", 5: "mordad", 6: "shahrivar", 7: "mehr", 8: "aban", 9: "azar", 10: "dey", 11: "bahman", 12: "esfand",
}

func print_str_time() (string) {
	var month_days int

	y, m, d := time.Now().Date()

	var year, month, day int
	if is_leap(y) {
		year, month, day = calc(y, int(m), d, true)
	} else {
		year, month, day = calc(y, int(m), d, false)
	}

	if month < 6 {
		month_days = 31
	} else if month > 6 {
		month_days = 30
	} else if month == 12 {
		if is_leap(y) {
			month_days = 29
		} else {
			month_days = 30
		}
	}

	print_str := fmt.Sprintf("    %v %v\n", persian_month[month], year)



	firstDayOfMonth := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
	weekDayOfFirst := int(firstDayOfMonth.Weekday())

	print_str += color.Sprintf("Sa Su Mo Tu We Th @rFr\n")
	for i := 0; i <= weekDayOfFirst; i++ {
		print_str += "   "
	}

	for i := 1; i <= month_days; i++ {
		if i == day  {
			print_str += color.Sprintf("@{Wk}%2d ", i)
		} else {
			print_str += fmt.Sprintf("%2d ", i)
		}
		if ((i+weekDayOfFirst)+1)%7 == 0 {
			print_str += "\n"
		}
	}

	return print_str
}

func calc(m_year int, m_month int, m_day int, is_leap bool) (int, int, int) {
	var month int
	var year int
	var date int
	var diff_between_days int
	if !is_leap {
		date = not_leap[m_month-1] + m_day
		if date > 79 {
			date = date - 79
			if date <= 186 {
				month = (date-1)/31 + 1
				date = (date-1)%31 + 1
				year = m_year - 621
			} else {
				date = date - 186
				month = (date-1)/30 + 7
				date = (date-1)%30 + 1
				year = m_year - 621
			}
		} else {
			if (m_year > 1996) && (m_year%4) == 1 {
				diff_between_days = 11
			} else {
				diff_between_days = 10
			}
			date = date + diff_between_days
			month = (date-1)/30 + 10
			date = (date-1)%30 + 1
			year = m_year - 622
		}
	} else {
		date = leap[m_month-1] + m_day
		if m_year >= 1996 {
			diff_between_days = 79
		} else {
			diff_between_days = 80
		}
		if date > diff_between_days {
			date = date - diff_between_days
			if date <= 186 {
				month = (date-1)/31 + 1
				date = (date-1)%31 + 1
				year = m_year - 621
			} else {
				date = date - 186
				month = (date-1)/30 + 7
				date = (date-1)%30 + 1
				year = m_year - 621
			}
		} else {
			date = date + 10
			month = (date-1)/30 + 10
			date = (date-1)%30 + 1
			year = m_year - 622
		}
	}
	return year, month, date
}

func is_leap(year int) bool {
	if year%4 == 0 {
		return true
	} else {
		return false
	}
}

func parse(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}
