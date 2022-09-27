package entity

import (
	"fmt"
	"time"
)

type Person struct {
	name     string
	age      uint8
	orgName  string
	birthday time.Time
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Age() uint8 {
	return p.age
}

func (p *Person) SetAge(age uint8) {
	p.age = age
}

func (p *Person) OrgName() string {
	return p.orgName
}

func (p *Person) SetOrgName(orgName string) {
	p.orgName = orgName
}

func (p *Person) Birthday() time.Time {
	return p.birthday
}

func (p *Person) SetBirthday(birthday time.Time) {
	p.birthday = birthday
}

func (p *Person) ShowInfo() {
	fmt.Println("this man is:", p.name, "age is:", p.age)
}

func NewPerson() *Person {
	p := Person{
		"asd", 23, "123", time.Now(),
	}
	//p.name = "王飞"
	//p.age = 90
	return &p
}
