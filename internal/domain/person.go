package domain

import (
	"strings"
	"time"
)

type Person struct {
	id          int
	name        string
	surname     string
	patronymic  *string
	age         *int
	gender      *string
	nationality *string
	createdAt   time.Time
}

func NewPerson(id int, name, surname string, patronymic, gender, nationality *string, age *int) *Person {
	return &Person{
		id:          id,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		age:         age,
		gender:      gender,
		nationality: nationality,
		createdAt:   time.Now(),
	}
}

func (p *Person) GetID() int              { return p.id }
func (p *Person) GetName() string         { return p.name }
func (p *Person) GetSurname() string      { return p.surname }
func (p *Person) GetPatronymic() *string  { return p.patronymic }
func (p *Person) GetAge() *int            { return p.age }
func (p *Person) GetGender() *string      { return p.gender }
func (p *Person) GetNationality() *string { return p.nationality }
func (p *Person) GetCreatedAt() time.Time { return p.createdAt }

func (p *Person) SetName(name string) {
	name = strings.TrimSpace(name)
	if name != "" {
		p.name = name
	}
}

func (p *Person) SetSurname(surname string) {
	surname = strings.TrimSpace(surname)
	if surname != "" {
		p.surname = surname
	}
}

func (p *Person) SetPatronymic(patronymic *string) {
	if patronymic != nil {
		value := strings.TrimSpace(*patronymic)
		if value != "" {
			p.patronymic = &value
		}
	}
}

func (p *Person) SetAge(age *int) {
	if age != nil && *age > 0 && *age < 150 {
		p.age = age
	}
}

func (p *Person) SetGender(gender *string) {
	if gender != nil {
		value := strings.ToLower(strings.TrimSpace(*gender))
		if value == "male" || value == "female" {
			p.gender = &value
		}
	}
}

func (p *Person) SetNationality(nationality *string) {
	if nationality != nil {
		value := strings.TrimSpace(*nationality)
		if value != "" {
			p.nationality = &value
		}
	}
}
