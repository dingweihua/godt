package godt

import (
	"time"
)

const (
	DefaultDateFormat = "2006-01-02"
	DefaultDtFormat   = "2006-01-02 15:04:05"
)

func GetMondayOfThisWeek() time.Time {
	now := time.Now()

	offset := int(now.Weekday() - time.Monday)
	// Sunday - 0, Saturday - 6
	if offset < 0 {
		offset = 6
	}

	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -offset)
}

func GetSundayOfThisWeek() time.Time {
	return GetMondayOfThisWeek().AddDate(0, 0, 6)
}

func GetMondayStrOfThisWeek(format string) string {
	return GetMondayOfThisWeek().Format(format)
}

func GetSundayStrOfThisWeek(format string) string {
	return GetMondayOfThisWeek().AddDate(0, 0, 6).Format(format)
}

func GetMondayStrOfLastWeek(format string) string {
	return GetMondayOfThisWeek().AddDate(0, 0, -7).Format(format)
}

func GetSundayStrOfLastWeek(format string) string {
	return GetMondayOfThisWeek().AddDate(0, 0, -1).Format(format)
}

func GetMondayUnixOfThisWeek() int64 {
	return GetMondayOfThisWeek().Unix()
}

func GetSundayUnixOfThisWeek() int64 {
	return GetSundayOfThisWeek().Unix()
}

func GetMondayUnixOfLastWeek() int64 {
	return GetMondayOfThisWeek().AddDate(0, 0, -7).Unix()
}

func GetSundayUnixOfLastWeek() int64 {
	return GetSundayOfThisWeek().AddDate(0, 0, -7).Unix()
}

func GetMondayUtcStrOfThisWeek(format string) string {
	return time.Unix(GetMondayUnixOfThisWeek(), 0).UTC().Format(format)
}

func GetSundayUtcStrOfThisWeek(format string) string {
	return time.Unix(GetSundayUnixOfThisWeek(), 0).UTC().Format(format)
}

func GetMondayUtcStrOfLastWeek(format string) string {
	return time.Unix(GetMondayUnixOfLastWeek(), 0).UTC().Format(format)
}

func GetSundayUtcStrOfLastWeek(format string) string {
	return time.Unix(GetSundayUnixOfLastWeek(), 0).UTC().Format(format)
}

func GetLastWeekDateMap(dateFormat string, dtFormat string) map[string]interface{} {
	return map[string]interface{}{
		"start_date":     GetMondayStrOfLastWeek(dateFormat),
		"end_date":       GetMondayStrOfThisWeek(dateFormat),
		"utc_start_date": GetMondayUtcStrOfLastWeek(dtFormat),
		"utc_end_date":   GetMondayUtcStrOfThisWeek(dtFormat),
		"first_date":     GetMondayStrOfLastWeek(dateFormat),
		"last_date":      GetSundayStrOfLastWeek(dateFormat),
		"start_date_s":   GetMondayUnixOfLastWeek(),
		"end_date_s":     GetMondayUnixOfThisWeek(),
		"start_date_ms":  GetMondayUnixOfLastWeek() * 1000,
		"end_date_ms":    GetMondayUnixOfThisWeek() * 1000,
	}
}

func GetDateMapByStartDateStr(startDateStr, dateFormat, dtFormat string) map[string]interface{} {
	startTime, err := time.ParseInLocation(dateFormat, startDateStr, time.Local)
	if err != nil {
		return map[string]interface{}{}
	}
	endTime := startTime.AddDate(0, 0, 7)
	firstTime, lastTime := startTime, startTime.AddDate(0, 0, 6)

	return map[string]interface{}{
		"start_date":     startTime.Format(dateFormat),
		"end_date":       endTime.Format(dateFormat),
		"utc_start_date": time.Unix(startTime.Unix(), 0).UTC().Format(dtFormat),
		"utc_end_date":   time.Unix(endTime.Unix(), 0).UTC().Format(dtFormat),
		"first_date":     firstTime.Format(dateFormat),
		"last_date":      lastTime.Format(dateFormat),
		"start_date_s":   startTime.Unix(),
		"end_date_s":     endTime.Unix(),
		"start_date_ms":  startTime.Unix() * 1000,
		"end_date_ms":    endTime.Unix() * 1000,
	}
}
