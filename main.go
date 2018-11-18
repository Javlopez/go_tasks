package main

import (
	"fmt"
	"time"
)

type User interface {
	PrintName()
	PrintDetails()
}

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, email: %s, Location: %s]\n", p.Dob.String(), p.Email, p.Location)
}

type Admin struct {
	Person
	Roles []string
}

func (a Admin) PrintDetails() {
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, r := range a.Roles {
		fmt.Println(r)
	}
}

type Member struct {
	Person
	Skills []string
}

func (m Member) PrintDetails() {
	m.Person.PrintDetails()
	for _, s := range m.Skills {
		fmt.Println(s)
	}
}

type Team struct {
	Name, Description string
	Users             []User
}

func (t Team) GetTeamDetails() {
	fmt.Printf("Team: %s - %s \n", t.Name, t.Description)
	fmt.Println("Details of the team members")
	for _, u := range t.Users {
		u.PrintName()
		u.PrintDetails()
	}
}

func main() {
	alex := Admin{
		Person{
			"Alex",
			"Jhon",
			time.Date(1970, time.January, 10, 0, 0, 0, 0, time.UTC),
			"alex@email.com",
			"New York"},
		[]string{"Manage Team", "Manage Tasks"},
	}

	shiju := Member{
		Person{
			"Shiju",
			"Var",
			time.Date(1979, time.March, 10, 0, 0, 0, 0, time.UTC),
			"chijju@email.com",
			"Kansas"},
		[]string{"Go", "Docker", "Kubernetes"},
	}

	javier := Member{
		Person{
			"Javier",
			"Lopez",
			time.Date(1999, time.December, 10, 0, 0, 0, 0, time.UTC),
			"senior@dev.com",
			"Berlin"},
		[]string{"Go", "Docker", "Kubernetes", "Terraform", "AWS"},
	}

	team := Team{
		"BackEnd Team",
		"Working on Backend, with GO and GRPC",
		[]User{alex, shiju, javier},
	}

	team.GetTeamDetails()
}
