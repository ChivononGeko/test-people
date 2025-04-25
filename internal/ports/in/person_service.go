package in

import (
	"context"

	"test-people/internal/domain"
)

type PersonService interface {
	AddPerson(ctx context.Context, name, surname string, patronymic *string) error
	UpdatePerson(ctx context.Context, id int, person *domain.Person) error
	DeletePerson(ctx context.Context, id int) error
	GetPersonByID(ctx context.Context, id int) (*domain.Person, error)
	GetByFilter(ctx context.Context, filter domain.PersonFilter) ([]*domain.Person, error)
}
