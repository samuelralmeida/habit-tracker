package localstore

import (
	"fmt"
	"time"

	"github.com/samuelralmeida/habit-tracker/database"
	"github.com/samuelralmeida/habit-tracker/entity"
)

var habits = map[int]entity.Habit{}
var tracks = map[int]entity.Tracks{}

type localstore struct{}

func New() database.Databaser {
	return &localstore{}
}

func (l *localstore) SaveHabit(habit *entity.Habit) error {
	if habit.ID == 0 {
		habit.ID = len(habits) + 1
	}
	habits[habit.ID] = *habit
	return nil
}

func (l *localstore) TrackHabit(habitID int) error {
	if !getHabit(habitID).IsValidHabit() {
		return fmt.Errorf("the habit ID %d not exists", habitID)
	}

	times := tracks[habitID]
	times = append(times, time.Now())
	tracks[habitID] = times
	return nil
}

func (l *localstore) FetchHabitTrack(habitID int) (*entity.HabitTrack, error) {
	habit := getHabit(habitID)
	tracks := tracks[habitID]
	return &entity.HabitTrack{Habit: *habit, Tracks: tracks}, nil
}

func (l *localstore) FetchHabits() ([]*entity.Habit, error) {
	resp := make([]*entity.Habit, len(habits))
	count := 0
	for _, habit := range habits {
		resp[count] = &habit
		count++
	}
	return resp, nil
}

func getHabit(habitID int) *entity.Habit {
	habit := habits[habitID]
	return &habit
}
