package main

import "time"

type DailyMetrics struct {
	ID             int       `json:"id"`
	EntryDate      string    `json:"entry_date"` // Use string for simple HTML date handling
	SleepQuality   int       `json:"sleep_quality"`
	PhysicalEnergy int       `json:"physical_energy"`
	Focus          int       `json:"focus"`
	Motivation     int       `json:"motivation"`
	PastView       int       `json:"past_view"`
	SocialActivity int       `json:"social_activity"`
	CreatedAt      time.Time `json:"created_at"`
}

