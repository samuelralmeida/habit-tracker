package database

import "github.com/samuelralmeida/habit-tracker/entity"

type Databaser interface {
	SaveHabit(habit *entity.Habit) error
	TrackHabit(habitID int) error
	FetchHabitTrack(habitID int) (*entity.HabitTrack, error)
	FetchHabits() ([]*entity.Habit, error)
}
