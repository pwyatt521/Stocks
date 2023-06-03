package yahoo

import "time"

func FormatInt64ToDate(timestamp int64) string {
	// Convert int64 to time.Time
	t := time.Unix(timestamp, 0)

	// Format the time as a readable date
	formattedDate := t.Format("2006-01-02") // Use your desired date format here

	return formattedDate
}

func FormatInt64ToTime(timestamp int64) string {
	// Convert int64 to time.Time
	t := time.Unix(timestamp, 0)

	// Format the time as a readable time
	formattedTime := t.Format("15:04:05") // Use your desired time format here

	return formattedTime
}
