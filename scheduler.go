package worker

import "time"

// Scheduler is a interface, for job used to `GetScheduleTime`
type Scheduler interface {
	GetScheduleTime() *time.Time
}

// Schedule could be embedded as job argument, then the job will get run as scheduled feature
type Schedule struct {
	ScheduleTime *time.Time
}

// GetScheduleTime get scheduled time
func (schedule Schedule) GetScheduleTime() *time.Time {
	return schedule.ScheduleTime
}
