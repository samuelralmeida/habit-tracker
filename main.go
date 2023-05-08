package main

import (
	"fmt"

	"github.com/samuelralmeida/habit-tracker/database/localstore"
	"github.com/samuelralmeida/habit-tracker/entity"
)

func main() {

	database := localstore.New()

	err := database.SaveHabit(&entity.Habit{ID: 22, Description: "Estudar inglês 30 minutos", Frequency: entity.Daily, FrequencyGoal: 1})
	fmt.Println(err)

	h, err := database.FetchHabitTrack(18)
	fmt.Println(err, h)

	h, err = database.FetchHabitTrack(22)
	fmt.Println(err, h)

	err = database.TrackHabit(22)
	fmt.Println(err)

	h, err = database.FetchHabitTrack(22)
	fmt.Println(err, h)
}
