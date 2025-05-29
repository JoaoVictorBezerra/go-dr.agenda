package model

import "time"

type Scheduler struct {
	Patient   Patient
	Medic     Medic
	Date      time.Time
	Status    SchedulerStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SchedulerStatus string

const (
	CONFIRMED SchedulerStatus = "CONFIRMED"
	UNMARKED  SchedulerStatus = "UNMARKED"
	DONE      SchedulerStatus = "DONE"
)
