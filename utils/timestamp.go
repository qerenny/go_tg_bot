package utils

import (
	"fmt"
	"math"
	"strings"
	"time"
)

var moscow *time.Location

func init() {
	var err error
	moscow, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(fmt.Sprintf("failed to load location: %v", err))
	}
}

func currentTimestamp() int64 {
	return time.Now().In(moscow).UnixMilli()
}

func currentTime() time.Time {
	return time.Now().In(moscow)
}

func timestampToDate(timestampInMilisec int64) string {
	return time.UnixMilli(timestampInMilisec).In(moscow).Format("2006-01-02 15:04:05")
}

func dateToTimestamp(date string) int64 {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		panic(err)
	}
	timestampInMillisec := parsedTime.In(moscow).UnixMilli()
	return timestampInMillisec
}

func setDate(years, months, days int) time.Time {
	return currentTime().AddDate(years, months, days)
}

func updateDate(previousDate string, years, months, days int) time.Time {

	if dateToTimestamp(previousDate) <= currentTimestamp() {
		previousDate = timestampToDate(currentTimestamp())
	}

	previousTime := time.UnixMilli(dateToTimestamp(previousDate)).In(moscow)
	updatedDate := previousTime.AddDate(years, months, days)
	return updatedDate
}

func getWordForm(number int, forms [3]string) string {
	n := int(math.Abs(float64(number))) % 100
	n1 := n % 10

	if 10 < n && n < 20 {
		return forms[2]
	}
	if n1 == 1 {
		return forms[0]
	}
	if 2 <= n1 && n1 <= 4 {
		return forms[1]
	}
	return forms[2]
}

func calculateTimeDifference(expiryDate string) string {
	userTimestamp := dateToTimestamp(expiryDate)
	currentTimestamp := currentTimestamp()

	if userTimestamp <= currentTimestamp {
		return "муф.. твоя подписка уже закончилась😿"
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", expiryDate)
	if err != nil {
		panic(err)
	}

	years := parsedTime.Year() - currentTime().Year()
	months := int(parsedTime.Month()) - int(currentTime().Month())
	days := parsedTime.Day() - currentTime().Day()
	hours := parsedTime.Hour() - currentTime().Hour()

	var resultParts []string

	if years > 0 {
		yearsStr := fmt.Sprintf("%d %s", years, getWordForm(years, [3]string{"год", "года", "лет"}))
		resultParts = append(resultParts, yearsStr)
	}

	if months > 0 {
		monthsStr := fmt.Sprintf("%d %s", months, getWordForm(months, [3]string{"месяц", "месяца", "месяцев"}))
		resultParts = append(resultParts, monthsStr)
	}

	if days > 0 {
		daysStr := fmt.Sprintf("%d %s", days, getWordForm(days, [3]string{"день", "дня", "дней"}))
		resultParts = append(resultParts, daysStr)
	}

	if hours > 0 {
		hoursStr := fmt.Sprintf("%d %s", hours, getWordForm(hours, [3]string{"час", "часа", "часов"}))
		resultParts = append(resultParts, hoursStr)
	}

	if len(resultParts) == 0 {
		return "кошечки-божечки, у тебя осталось меньше часа🙀"
	}

	return strings.Join(resultParts, ", ")
}
