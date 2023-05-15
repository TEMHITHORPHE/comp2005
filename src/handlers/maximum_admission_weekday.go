package handlers

import (
	"net/http"
	HttpReq "phxlabs/m/comp2005/automated_testing/src/externalHTTP"
	model "phxlabs/m/comp2005/automated_testing/src/models"
	"time"
)

var WEEKDAY_TO_SEQUENCE_MAP = map[string]int{
	"Sunday":    0,
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
}
var SEQUENCE_TO_WEEKDAY_MAP = map[int]string{
	0: "Sunday",
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
}

func WeekDayOfMaximumAdmissions(w http.ResponseWriter, r *http.Request) {

	// Retrieve all Admissions
	var admissions []model.Admission
	var statusCode int = HttpReq.GetAllAdmissions_HttpReq(&admissions)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return
	}

	// Where each index represents a week day
	weeekdayOccurranceFrequencyTracker := []int{0, 0, 0, 0, 0, 0, 0}

	// Populate the occurance frequency tracker array.
	for _, admission := range admissions {
		dateAdmitted, err := time.Parse(ADMISSION_DATE_LAYOUT, admission.AdmissionDate)
		if err != nil {
			pr(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		weekDay := dateAdmitted.Weekday().String()
		pr(weekDay)
		var weekdaySequenceAsIndex int = WEEKDAY_TO_SEQUENCE_MAP[weekDay]
		weeekdayOccurranceFrequencyTracker[weekdaySequenceAsIndex] = weeekdayOccurranceFrequencyTracker[weekdaySequenceAsIndex] + 1
	}

	pr(weeekdayOccurranceFrequencyTracker)

	maxStartIndex := 0
	maxFreq := 0
	for i := 0; i < 7; i++ {
		if weeekdayOccurranceFrequencyTracker[i] > maxFreq {
			maxStartIndex = i
			maxFreq = weeekdayOccurranceFrequencyTracker[i]
		}
	}

	// extract array subset containing max values.
	weeekdayOccurranceFrequencyTracker = weeekdayOccurranceFrequencyTracker[maxStartIndex:]

	// Extract maximum week days and build JSON response.
	maxWeekdays := make(map[string]int, maxStartIndex)
	for index := 0; index < len(weeekdayOccurranceFrequencyTracker); index++ {
		freq := weeekdayOccurranceFrequencyTracker[index]
		if freq == maxFreq {
			weekDay := SEQUENCE_TO_WEEKDAY_MAP[index+maxStartIndex]
			maxWeekdays[weekDay] = maxFreq
		}
	}

	sendJSONRespose(w, &maxWeekdays)
}
