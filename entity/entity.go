package entity

import "time"

type frequency struct {
	Option      int
	Description string
}

func (f *frequency) IsValidFrequency() bool {
	return f.Description != ""
}

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

var frequencies = map[int]frequency{
	0: {Option: 0, Description: "daily"},
	1: {Option: 1, Description: "weekly"},
	2: {Option: 2, Description: "monthly"},
}

func OptionsFrequency() []frequency {
	return []frequency{frequencies[0], frequencies[1], frequencies[2]}
}

func ParseFrequency(option int) frequency {
	return frequencies[option]
}
