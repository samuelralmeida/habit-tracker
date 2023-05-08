package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/samuelralmeida/habit-tracker/database"
	"github.com/samuelralmeida/habit-tracker/database/localstore"
	"github.com/samuelralmeida/habit-tracker/entity"
)

func main() {
	db := localstore.New()

	reader := bufio.NewReader(os.Stdin)
	errorCount := 0

	for {
		if errorCount > 3 {
			break
		}

		printInitialOptions()

		text, _ := askInput(reader)

		fn, err := checkInitialOptions(text)
		if err != nil {
			errorCount++
			log.Println("error to check initial options:", err)
			continue
		}

		err = fn(reader, db)
		if err != nil {
			errorCount++
			log.Printf("error to execute option (%s): %s\n", text, err)
			continue
		}

	}
}

func options() []string {
	return []string{
		"Exit",
		"Add an habit",
		"Track an habit",
		"Check an habit",
	}
}

func printInitialOptions() {
	fmt.Println("What do you want?")

	for i, option := range options() {
		fmt.Printf("%d - %s\n", i, option)
	}

	fmt.Println("---------------------")
}

func askInput(reader *bufio.Reader) (string, error) {
	fmt.Print("-> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Replace(text, "\n", "", -1), nil
}

func checkInitialOptions(text string) (func(*bufio.Reader, database.Databaser) error, error) {
	resp, err := convertOption(text)
	if err != nil {
		return nil, err
	}

	endProgram(resp)

	switch resp {
	case 1:
		return addHabit, nil
	case 2:
		return trackHabit, nil
	case 3:
		return checkHabit, nil
	default:
		return nil, errors.New("invalid option")
	}
}

func endProgram(option int) {
	if option == 0 {
		os.Exit(0)
	}
}

func addHabit(reader *bufio.Reader, db database.Databaser) error {
	fmt.Println("Describe an habit to add:")
	description, _ := askInput(reader)

	fmt.Println("What is the interval frequency?")
	frequencies := entity.OptionsFrequency()
	for _, frequency := range frequencies {
		fmt.Printf("%d - %s\n", frequency.Option, frequency.Description)
	}
	frequency, _ := askInput(reader)

	fmt.Println("How many times do you want to do this habit in the previous interval?")
	times, _ := askInput(reader)

	habit, err := createHabit(description, frequency, times)
	if err != nil {
		return fmt.Errorf("error to create habit: %w", err)
	}

	return db.SaveHabit(habit)
}

func createHabit(description, frequency, times string) (*entity.Habit, error) {
	frequencyChoice, err := convertOption(frequency)
	if err != nil {
		return nil, fmt.Errorf("error to convert frequency: %w", err)
	}

	f := entity.ParseFrequency(frequencyChoice)
	if !f.IsValidFrequency() {
		return nil, errors.New("interval frequency invalid")
	}

	t, err := convertOption(times)
	if err != nil {
		return nil, fmt.Errorf("error to convert time: %w", err)
	}

	return &entity.Habit{Description: description, Frequency: f, FrequencyGoal: t}, nil
}

func trackHabit(reader *bufio.Reader, db database.Databaser) error {
	habits, err := db.FetchHabits()
	if err != nil {
		return err
	}

	fmt.Println("What habit do you want to track?")
	for _, habit := range habits {
		fmt.Printf("%d - %s\n", habit.ID, habit.Description)
	}

	id, _ := askInput(reader)

	habitID, err := convertOption(id)
	if err != nil {
		return fmt.Errorf("error to convert habit id: %w", err)
	}

	return db.TrackHabit(habitID)
}

func checkHabit(reader *bufio.Reader, db database.Databaser) error {
	habits, err := db.FetchHabits()
	if err != nil {
		return err
	}

	fmt.Println("What habit do you want to check?")
	for _, habit := range habits {
		fmt.Printf("%d - %s\n", habit.ID, habit.Description)
	}

	id, _ := askInput(reader)

	habitID, err := convertOption(id)
	if err != nil {
		return fmt.Errorf("error to convert habit id: %w", err)
	}

	habitTrack, err := db.FetchHabitTrack(habitID)
	if err != nil {
		return fmt.Errorf("error to fetch habit track: %w", err)
	}

	fmt.Println("Habit description:", habitTrack.Description)
	fmt.Println("Habit frequency:", habitTrack.Frequency.Description)
	fmt.Println("Habit frequency goal:", habitTrack.FrequencyGoal)

	now := time.Now()

	initDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	fmt.Println("You have already done in the current month:")
	total := 0
	for _, track := range habitTrack.Tracks {
		if track.After(initDate) && track.Before(now) {
			total++
			fmt.Printf("- %s\n", track.Format(time.RFC3339))
		}
	}

	fmt.Println("Total:", total)

	return nil
}

func convertOption(input string) (int, error) {
	return strconv.Atoi(input)
}
