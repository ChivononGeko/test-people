package transport

import (
	"test-people/internal/domain"
)

type PersonDTO struct {
	ID          int     `json:"id" example:"1"`
	Name        string  `json:"name" example:"Ivan"`
	Surname     string  `json:"surname" example:"Petrov"`
	Patronymic  string  `json:"patronymic" example:"Ivanovich"`
	Age         *int    `json:"age,omitempty" example:"30"`
	Gender      *string `json:"gender,omitempty" example:"male"`
	Nationality *string `json:"nationality,omitempty" example:"Russian"`
}

func (dto *PersonDTO) ToDomain() *domain.Person {
	person := domain.NewPerson(dto.ID, dto.Name, dto.Surname, &dto.Patronymic, dto.Gender, dto.Nationality, dto.Age)
	return person
}

func FromDomain(p *domain.Person) *PersonDTO {
	return &PersonDTO{
		ID:          p.GetID(),
		Name:        p.GetName(),
		Surname:     p.GetSurname(),
		Patronymic:  *p.GetPatronymic(),
		Age:         p.GetAge(),
		Gender:      p.GetGender(),
		Nationality: p.GetNationality(),
	}
}
