package godt

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"
	"time"
)

// Before testing, update the current date first.
var (
	thisYear             = 2022
	thisMonth            = 1
	thisMonday           = 17
	thisMondayDate       = time.Date(thisYear, time.Month(thisMonth), thisMonday, 0, 0, 0, 0, time.Local)
	thisMondayDateStr    = thisMondayDate.Format(DefaultDateFormat)
	thisMondayUtcDateStr = thisMondayDate.Add(time.Hour * -8).Format(DefaultDateFormat)
	thisMondayDtStr      = thisMondayDate.Format(DefaultDtFormat)
	thisMondayUtcDtStr   = thisMondayDate.Add(time.Hour * -8).Format(DefaultDtFormat)
	thisSundayDateStr    = thisMondayDate.AddDate(0, 0, 6).Format(DefaultDateFormat)
	thisSundayUtcDateStr = thisMondayDate.AddDate(0, 0, 6).Add(time.Hour * -8).Format(DefaultDateFormat)
	thisSundayDtStr      = thisMondayDate.AddDate(0, 0, 6).Format(DefaultDtFormat)
	thisSundayUtcDtStr   = thisMondayDate.AddDate(0, 0, 6).Add(time.Hour * -8).Format(DefaultDtFormat)
	lastMondayDateStr    = thisMondayDate.AddDate(0, 0, -7).Format(DefaultDateFormat)
	lastMondayUtcDateStr = thisMondayDate.AddDate(0, 0, -7).Add(time.Hour * -8).Format(DefaultDateFormat)
	lastMondayDtStr      = thisMondayDate.AddDate(0, 0, -7).Format(DefaultDtFormat)
	lastMondayUtcDtStr   = thisMondayDate.AddDate(0, 0, -7).Add(time.Hour * -8).Format(DefaultDtFormat)
	lastSundayDateStr    = thisMondayDate.AddDate(0, 0, -1).Format(DefaultDateFormat)
	lastSundayUtcDateStr = thisMondayDate.AddDate(0, 0, -1).Add(time.Hour * -8).Format(DefaultDateFormat)
	lastSundayDtStr      = thisMondayDate.AddDate(0, 0, -1).Format(DefaultDtFormat)
	lastSundayUtcDtStr   = thisMondayDate.AddDate(0, 0, -1).Add(time.Hour * -8).Format(DefaultDtFormat)

	thisMondayUnix = thisMondayDate.Unix()
	thisSundayUnix = thisMondayUnix + 6*24*60*60
	lastMondayUnix = thisMondayUnix - 7*24*60*60
	lastSundayUnix = thisMondayUnix - 1*24*60*60
)

func TestGetMondayOfThisWeek(t *testing.T) {
	got := GetMondayOfThisWeek()
	want := thisMondayDate
	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted:%#v, got:%#v", want, got)
	}
}

func TestGetSundayOfThisWeek(t *testing.T) {
	got := GetSundayOfThisWeek()
	want := thisMondayDate.AddDate(0, 0, 6)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted:%#v, got:%#v", want, got)
	}
}

func TestGetMondayDateStrOfThisWeek(t *testing.T) {
	got := GetMondayStrOfThisWeek(DefaultDateFormat)
	if !reflect.DeepEqual(got, thisMondayDateStr) {
		t.Errorf("excepted:%#v, got:%#v", thisMondayDateStr, got)
	}
}

func TestGetMondayDatetimeStrOfThisWeek(t *testing.T) {
	got := GetMondayStrOfThisWeek(DefaultDtFormat)
	if !reflect.DeepEqual(got, thisMondayDtStr) {
		t.Errorf("excepted:%#v, got:%#v", thisMondayDtStr, got)
	}
}

func TestGetSundayDateStrOfThisWeek(t *testing.T) {
	got := GetSundayStrOfThisWeek(DefaultDateFormat)
	if !reflect.DeepEqual(got, thisSundayDateStr) {
		t.Errorf("excepted:%#v, got:%#v", thisSundayDateStr, got)
	}
}

func TestGetSundayDatetimeStrOfThisWeek(t *testing.T) {
	got := GetSundayStrOfThisWeek(DefaultDtFormat)
	if !reflect.DeepEqual(got, thisSundayDtStr) {
		t.Errorf("excepted:%#v, got:%#v", thisSundayDtStr, got)
	}
}

func TestGetMondayUnixOfThisWeek(t *testing.T) {
	got := GetMondayUnixOfThisWeek()
	if !reflect.DeepEqual(got, thisMondayUnix) {
		t.Errorf("excepted:%#v, got:%#v", thisMondayUnix, got)
	}
}

func TestGetSundayUnixOfThisWeek(t *testing.T) {
	got := GetSundayUnixOfThisWeek()
	if !reflect.DeepEqual(got, thisSundayUnix) {
		t.Errorf("excepted:%#v, got:%#v", thisSundayUnix, got)
	}
}

func TestGetMondayDateStrOfLastWeek(t *testing.T) {
	got := GetMondayStrOfLastWeek(DefaultDateFormat)
	if !reflect.DeepEqual(got, lastMondayDateStr) {
		t.Errorf("excepted:%#v, got:%#v", lastMondayDateStr, got)
	}
}

func TestGetMondayDatetimeStrOfLastWeek(t *testing.T) {
	got := GetMondayStrOfLastWeek(DefaultDtFormat)
	if !reflect.DeepEqual(got, lastMondayDtStr) {
		t.Errorf("excepted:%#v, got:%#v", lastMondayDtStr, got)
	}
}

func TestGetMondayUnixOfLastWeek(t *testing.T) {
	got := GetMondayUnixOfLastWeek()
	if !reflect.DeepEqual(got, lastMondayUnix) {
		t.Errorf("excepted:%#v, got:%#v", lastMondayUnix, got)
	}
}

func TestGetSundayDateStrOfLastWeek(t *testing.T) {
	got := GetSundayStrOfLastWeek(DefaultDateFormat)
	if !reflect.DeepEqual(got, lastSundayDateStr) {
		t.Errorf("excepted:%#v, got:%#v", lastSundayDateStr, got)
	}
}

func TestGetSundayDatetimeStrOfLastWeek(t *testing.T) {
	got := GetSundayStrOfLastWeek(DefaultDtFormat)
	if !reflect.DeepEqual(got, lastSundayDtStr) {
		t.Errorf("excepted:%#v, got:%#v", lastSundayDtStr, got)
	}
}

func TestGetSundayUnixOfLastWeek(t *testing.T) {
	got := GetSundayUnixOfLastWeek()
	if !reflect.DeepEqual(got, lastSundayUnix) {
		t.Errorf("excepted:%#v, got:%#v", lastSundayUnix, got)
	}
}

func TestGetMondayUTCDateStrOfThisWeek(t *testing.T) {
	got := GetMondayUtcStrOfThisWeek(DefaultDateFormat)
	if !reflect.DeepEqual(got, thisMondayUtcDateStr) {
		t.Errorf("excepted:%#v, got:%#v", thisMondayUtcDateStr, got)
	}
}

func TestGetMondayUTCDatetimeStrOfThisWeek(t *testing.T) {
	got := GetMondayUtcStrOfThisWeek(DefaultDtFormat)
	if !reflect.DeepEqual(got, thisMondayUtcDtStr) {
		t.Errorf("excepted:%#v, got:%#v", thisMondayUtcDtStr, got)
	}
}

func TestGetSundayUTCDateStrOfThisWeek(t *testing.T) {
	got := GetSundayUtcStrOfThisWeek(DefaultDateFormat)
	if !reflect.DeepEqual(got, thisSundayUtcDateStr) {
		t.Errorf("excepted:%#v, got:%#v", thisSundayUtcDateStr, got)
	}
}

func TestGetSundayUTCDatetimeStrOfThisWeek(t *testing.T) {
	got := GetSundayUtcStrOfThisWeek(DefaultDtFormat)
	if !reflect.DeepEqual(got, thisSundayUtcDtStr) {
		t.Errorf("excepted:%#v, got:%#v", thisSundayUtcDtStr, got)
	}
}

func TestGetMondayUTCDateStrOfLastWeek(t *testing.T) {
	got := GetMondayUtcStrOfLastWeek(DefaultDateFormat)
	if !reflect.DeepEqual(got, lastMondayUtcDateStr) {
		t.Errorf("excepted:%#v, got:%#v", lastMondayUtcDateStr, got)
	}
}

func TestGetMondayUTCDatetimeStrOfLastWeek(t *testing.T) {
	got := GetMondayUtcStrOfLastWeek(DefaultDtFormat)
	if !reflect.DeepEqual(got, lastMondayUtcDtStr) {
		t.Errorf("excepted:%#v, got:%#v", lastMondayUtcDtStr, got)
	}
}

func TestGetSundayUTCDateStrOfLastWeek(t *testing.T) {
	got := GetSundayUtcStrOfLastWeek(DefaultDateFormat)
	if !reflect.DeepEqual(got, lastSundayUtcDateStr) {
		t.Errorf("excepted:%#v, got:%#v", lastSundayUtcDateStr, got)
	}
}

func TestGetSundayUTCDatetimeStrOfLastWeek(t *testing.T) {
	got := GetSundayUtcStrOfLastWeek(DefaultDtFormat)
	if !reflect.DeepEqual(got, lastSundayUtcDtStr) {
		t.Errorf("excepted:%#v, got:%#v", lastSundayUtcDtStr, got)
	}
}

func mapToKeySortedList(inputMap map[string]interface{}) []map[string]interface{} {
	var (
		keyList    []string
		sortedList []map[string]interface{}
	)
	for key, _ := range inputMap {
		keyList = append(keyList, key)
	}
	sort.Strings(keyList)
	for _, key := range keyList {
		sortedList = append(sortedList, map[string]interface{}{
			key: inputMap[key],
		})
	}
	return sortedList
}

func TestGetLastWeekDateMap(t *testing.T) {
	got := GetLastWeekDateMap(DefaultDateFormat, DefaultDtFormat)
	want := map[string]interface{}{
		"start_date":     lastMondayDateStr,
		"end_date":       thisMondayDateStr,
		"utc_start_date": lastMondayUtcDtStr,
		"utc_end_date":   thisMondayUtcDtStr,
		"first_date":     lastMondayDateStr,
		"last_date":      lastSundayDateStr,
		"start_date_s":   lastMondayUnix,
		"end_date_s":     thisMondayUnix,
		"start_date_ms":  lastMondayUnix * 1000,
		"end_date_ms":    thisMondayUnix * 1000,
	}

	gotListJson, _ := json.Marshal(mapToKeySortedList(got))
	wantListJson, _ := json.Marshal(mapToKeySortedList(want))
	if !reflect.DeepEqual(gotListJson, wantListJson) {
		t.Errorf("excepted: %#v, got: %#v", wantListJson, gotListJson)
	}
}
