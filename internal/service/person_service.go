package service

import (
	"context"
	"fmt"
	"log/slog"

	"test-people/internal/domain"
	"test-people/internal/ports/in"
	"test-people/internal/ports/out"
)

type personService struct {
	repo     out.PersonRepository
	enricher out.EnrichmentClient
}

func NewPersonService(repo out.PersonRepository, enricher out.EnrichmentClient) in.PersonService {
	return &personService{
		repo:     repo,
		enricher: enricher,
	}
}

func (s *personService) AddPerson(ctx context.Context, name, surname string, patronymic *string) error {
	slog.Info("Adding person", slog.String("name", name), slog.String("surname", surname))

	age, err := s.enricher.GetAge(name)
	if err != nil {
		slog.Error("Failed to get age", slog.String("name", name), slog.String("error", err.Error()))
		return fmt.Errorf("failed to get age: %w", err)
	}

	gender, err := s.enricher.GetGender(name)
	if err != nil {
		slog.Error("Failed to get gender", slog.String("name", name), slog.String("error", err.Error()))
		return fmt.Errorf("failed to get gender: %w", err)
	}

	nationality, err := s.enricher.GetNationality(name)
	if err != nil {
		slog.Error("Failed to get nationality", slog.String("name", name), slog.String("error", err.Error()))
		return fmt.Errorf("failed to get nationality: %w", err)
	}

	p := domain.NewPerson(0, name, surname, patronymic, gender, nationality, age)

	if err := s.repo.Save(ctx, p); err != nil {
		slog.Error("Failed to save person", slog.String("name", name), slog.String("error", err.Error()))
		return fmt.Errorf("failed to save person: %w", err)
	}

	slog.Info("Person saved successfully", slog.String("name", name))
	return nil
}

func (s *personService) UpdatePerson(ctx context.Context, id int, person *domain.Person) error {
	slog.Info("Updating person", slog.Int("id", id))

	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		slog.Error("Person not found", slog.Int("id", id), slog.String("error", err.Error()))
		return fmt.Errorf("person not found: %w", err)
	}

	p.SetName(person.GetName())
	p.SetSurname(person.GetSurname())
	p.SetAge(person.GetAge())
	p.SetPatronymic(person.GetPatronymic())
	p.SetGender(person.GetGender())
	p.SetNationality(person.GetNationality())

	if err := s.repo.Update(ctx, p); err != nil {
		slog.Error("Failed to update person", slog.Int("id", id), slog.String("error", err.Error()))
		return err
	}

	slog.Info("Person updated", slog.Int("id", id))
	return nil
}

func (s *personService) DeletePerson(ctx context.Context, id int) error {
	slog.Info("Deleting person", slog.Int("id", id))
	err := s.repo.Delete(ctx, id)
	if err != nil {
		slog.Error("Failed to delete person", slog.Int("id", id), slog.String("error", err.Error()))
	}
	return err
}

func (s *personService) GetPersonByID(ctx context.Context, id int) (*domain.Person, error) {
	slog.Info("Getting person by ID", slog.Int("id", id))
	return s.repo.FindByID(ctx, id)
}

func (s *personService) GetByFilter(ctx context.Context, filter domain.PersonFilter) ([]*domain.Person, error) {
	slog.Info("Getting persons by filter")
	return s.repo.GetByFilter(ctx, filter)
}
