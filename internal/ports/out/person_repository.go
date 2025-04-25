package out

import (
	"context"

	"test-people/internal/domain"
)

type PersonRepository interface {
	Save(ctx context.Context, p *domain.Person) error
	Update(ctx context.Context, p *domain.Person) error
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (*domain.Person, error)
	GetByFilter(ctx context.Context, filter domain.PersonFilter) ([]*domain.Person, error)
}
