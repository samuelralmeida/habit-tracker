package entity

import "time"

type frequency string

const (
	Monthly frequency = "monthly"
	Weekly  frequency = "weekly"
	Daily   frequency = "daily"
)

type Tracks []time.Time

type Habit struct {
	ID            int
	Description   string
	Frequency     frequency
	FrequencyGoal int
}

func (h *Habit) IsValidHabit() bool {
	return h.ID > 0
}

type HabitTrack struct {
	Habit
	Tracks
}
