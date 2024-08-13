package job

import "time"

type Job struct {
	ID          string    `json:"id"`
	Cron        string    `json:"cron"`
	Command     string    `json:"command"`
	RunRightNow bool      `json:"run_right_now"`
	CreatedAt   time.Time `json:"created_at"`
}
