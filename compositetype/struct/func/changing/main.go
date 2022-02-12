package main

import (
	"fmt"
	"log"
	"time"
)

type CV struct {
	Person     Person
	Experience map[int]Experience
	Skills     []Skill
}

type Person struct {
	FirstName string
	LastName  string
}

type Experience struct {
	StartDate time.Time
	EndDate   time.Time
	Company   string
}

type Skill string

func NewCV(firstName, lastName string) *CV {
	return &CV{
		Person:     Person{FirstName: firstName, LastName: lastName},
		Experience: map[int]Experience{},
	}
}

// experience will be added to the original cv because
// Experience is a map which the pointer to the hmap struct
func (cv CV) AddExperience(index int, e Experience) {
	cv.Experience[index] = e
}

// Added experience will be not visible in the parent func
// because Skills is the slice and it will be copied
// during sending to the method, underlying array can
// be changed inside the func, but slice capacity and length
// won't be changed because we will change the copies of them.
func (cv CV) AddSkill(skill Skill) {
	cv.Skills = append(cv.Skills, skill)
}

func (cv *CV) AddSkillPtr(skill Skill) {
	cv.Skills = append(cv.Skills, skill)
}

func main() {
	cv := NewCV("Dmitriy", "Neustroev")

	mustParseDate := func(date string) time.Time {
		startTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Fatalf("can't parse date %v: %v", date, err)
		}

		return startTime
	}

	cv.AddExperience(1, Experience{
		StartDate: mustParseDate("2020-01-23"),
		EndDate:   mustParseDate("2021-05-01"),
		Company:   "Enterra",
	})

	fmt.Printf("%v\n", cv.Experience)

	cv.AddSkill("Golang")
	fmt.Printf("After AddSkill %v\n", cv.Skills)

	cv.AddSkillPtr("PHP")
	fmt.Printf("After AddSkillPtr %v\n", cv.Skills)
}
