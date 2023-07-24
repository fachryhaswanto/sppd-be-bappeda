package helper

import "time"

func ConvertToTimeMonth(bulan string) time.Month {
	var month time.Month

	if bulan == "1" {
		month = time.January
	} else if bulan == "2" {
		month = time.February
	} else if bulan == "3" {
		month = time.March
	} else if bulan == "4" {
		month = time.April
	} else if bulan == "5" {
		month = time.May
	} else if bulan == "6" {
		month = time.June
	} else if bulan == "7" {
		month = time.July
	} else if bulan == "8" {
		month = time.August
	} else if bulan == "9" {
		month = time.September
	} else if bulan == "10" {
		month = time.October
	} else if bulan == "11" {
		month = time.November
	} else if bulan == "12" {
		month = time.December
	}

	return month
}
