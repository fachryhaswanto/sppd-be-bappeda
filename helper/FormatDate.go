package helper

import "time"

func FormatDate(inputDateString string) string {
	var inputDate, err = time.Parse("2006/1/02", inputDateString)
	if err != nil {
		return inputDateString
	}

	var outputDateString = inputDate.Format("02/1/2006")

	return outputDateString
}
