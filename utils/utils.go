package utils

import "time"

// ParseDate converte uma string no formato "2006-01-02" para um ponteiro de time.Time
func ParseDate(date string) *time.Time {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil
	}
	return &t
}
