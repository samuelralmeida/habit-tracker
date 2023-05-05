package entity

import "time"

type Tracks []time.Time

type Habit struct {
	ID                int
	Description       string
	WeekFrequencyGoal int
}

func (h *Habit) IsValidHabit() bool {
	return h.ID > 0
}

type HabitTrack struct {
	Habit
	Tracks
}
