package repository

import (
	"context"
	"fmt"
	"log/slog"
	"test-people/internal/domain"
	"test-people/internal/ports/out"

	"github.com/jackc/pgx/v4"
)

type postgresPersonRepository struct {
	db *pgx.Conn
}

func NewPostgresPersonRepository(db *pgx.Conn) out.PersonRepository {
	return &postgresPersonRepository{db: db}
}

func (r *postgresPersonRepository) FindByID(ctx context.Context, id int) (*domain.Person, error) {
	slog.Info("FindByID called", "id", id)
	query := `
		SELECT p.id, p.name, p.surname, p.patronymic, p.age, g.name AS gender, n.name AS nationality, p.created_at
		FROM persons p
		LEFT JOIN genders g ON g.id = p.gender_id
		LEFT JOIN nationalities n ON n.id = p.nationality_id
		WHERE p.id = $1
	`

	row := r.db.QueryRow(ctx, query, id)

	var raw RawPerson
	err := row.Scan(&raw.ID, &raw.Name, &raw.Surname, &raw.Patronymic, &raw.Age, &raw.Gender, &raw.Nationality, &raw.CreatedAt)
	if err != nil {
		slog.Error("Failed to scan row", "error", err)
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	person := domain.NewPerson(raw.ID, raw.Name, raw.Surname, raw.Patronymic, raw.Gender, raw.Nationality, raw.Age)

	slog.Info("FindByID successful", "id", id)
	return person, nil
}

func (r *postgresPersonRepository) Save(ctx context.Context, p *domain.Person) error {
	slog.Info("Save called", "name", p.GetName(), "surname", p.GetSurname())
	query := `
		INSERT INTO persons (name, surname, patronymic, age, gender_id, nationality_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	tx, err := r.db.Begin(ctx)
	if err != nil {
		slog.Error("Failed to begin transaction", "error", err)
		return err
	}
	defer tx.Rollback(ctx)

	genderID, err := r.getOrCreateGenderID(ctx, p.GetGender())
	if err != nil {
		slog.Error("Failed to get or create gender ID", "error", err)
		return err
	}
	nationalityID, err := r.getOrCreateNationalityID(ctx, p.GetNationality())
	if err != nil {
		slog.Error("Failed to get or create nationality ID", "error", err)
		return err
	}

	_, err = tx.Exec(ctx, query, p.GetName(), p.GetSurname(), p.GetPatronymic(), p.GetAge(), genderID, nationalityID)
	if err != nil {
		slog.Error("Failed to execute query", "query", query, "error", err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		slog.Error("Failed to commit transaction", "error", err)
		return err
	}

	slog.Info("Save successful", "name", p.GetName(), "surname", p.GetSurname())
	return nil
}

func (r *postgresPersonRepository) Update(ctx context.Context, p *domain.Person) error {
	slog.Info("Update called", "id", p.GetID())
	query := `
		UPDATE persons SET name=$1, surname=$2, patronymic=$3, age=$4, gender_id=$5, nationality_id=$6 
		WHERE id=$7
	`

	genderID, err := r.getOrCreateGenderID(ctx, p.GetGender())
	if err != nil {
		slog.Error("Failed to get or create gender ID", "error", err)
		return err
	}
	nationalityID, err := r.getOrCreateNationalityID(ctx, p.GetNationality())
	if err != nil {
		slog.Error("Failed to get or create nationality ID", "error", err)
		return err
	}

	_, err = r.db.Exec(ctx, query, p.GetName(), p.GetSurname(), p.GetPatronymic(), p.GetAge(), genderID, nationalityID, p.GetID())
	if err != nil {
		slog.Error("Failed to execute query", "query", query, "error", err)
		return err
	}

	slog.Info("Update successful", "id", p.GetID())
	return nil
}

func (r *postgresPersonRepository) Delete(ctx context.Context, id int) error {
	slog.Info("Delete called", "id", id)
	query := `DELETE FROM persons WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		slog.Error("Failed to execute delete query", "error", err)
		return err
	}
	slog.Info("Delete successful", "id", id)
	return nil
}

func (r *postgresPersonRepository) GetByFilter(ctx context.Context, filter domain.PersonFilter) ([]*domain.Person, error) {
	slog.Info("GetByFilter called", "filter", filter)
	query := `
		SELECT p.id, p.name, p.surname, p.patronymic, p.age, g.name AS gender, n.name AS nationality, p.created_at
		FROM persons p
		LEFT JOIN genders g ON g.id = p.gender_id
		LEFT JOIN nationalities n ON n.id = p.nationality_id
		WHERE 1=1
	`
	args := []interface{}{}
	i := 1

	if filter.GetName() != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", i)
		args = append(args, "%"+filter.GetName()+"%")
		i++
	}
	if filter.GetSurname() != "" {
		query += fmt.Sprintf(" AND surname ILIKE $%d", i)
		args = append(args, "%"+filter.GetSurname()+"%")
		i++
	}
	if filter.GetPatronymic() != "" {
		query += fmt.Sprintf(" AND patronymic ILIKE $%d", i)
		args = append(args, "%"+filter.GetPatronymic()+"%")
		i++
	}
	if filter.GetGender() != "" {
		query += fmt.Sprintf(" AND gender_id = (SELECT id FROM genders WHERE name = $%d)", i)
		args = append(args, filter.GetGender())
		i++
	}
	if filter.GetNationality() != "" {
		query += fmt.Sprintf(" AND nationality_id = (SELECT id FROM nationalities WHERE name = $%d)", i)
		args = append(args, filter.GetNationality())
		i++
	}
	if filter.GetAge() != nil {
		query += fmt.Sprintf(" AND age = $%d", i)
		args = append(args, *filter.GetAge())
		i++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		slog.Error("Failed to execute query", "query", query, "error", err)
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var result []*domain.Person
	for rows.Next() {
		var raw RawPerson
		err := rows.Scan(&raw.ID, &raw.Name, &raw.Surname, &raw.Patronymic, &raw.Age, &raw.Gender, &raw.Nationality, &raw.CreatedAt)
		if err != nil {
			slog.Error("Failed to scan row", "error", err)
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		person := domain.NewPerson(raw.ID, raw.Name, raw.Surname, raw.Patronymic, raw.Gender, raw.Nationality, raw.Age)
		result = append(result, person)
	}

	slog.Info("GetByFilter successful", "filter", filter)
	return result, nil
}

// Логирование в дополнительных методах
func (r *postgresPersonRepository) getOrCreateGenderID(ctx context.Context, gender *string) (*int, error) {
	slog.Debug("getOrCreateGenderID called", "gender", gender)
	if gender == nil {
		return nil, nil
	}
	var id int
	err := r.db.QueryRow(ctx, `SELECT id FROM genders WHERE name = $1`, *gender).Scan(&id)
	if err == pgx.ErrNoRows {
		err = r.db.QueryRow(ctx, `INSERT INTO genders(name) VALUES($1) RETURNING id`, *gender).Scan(&id)
	}
	if err != nil {
		slog.Error("Failed to get or create gender ID", "gender", gender, "error", err)
		return nil, err
	}
	slog.Debug("getOrCreateGenderID successful", "gender", gender, "id", id)
	return &id, nil
}

func (r *postgresPersonRepository) getOrCreateNationalityID(ctx context.Context, nationality *string) (*int, error) {
	slog.Debug("getOrCreateNationalityID called", "nationality", nationality)
	if nationality == nil {
		return nil, nil
	}
	var id int
	err := r.db.QueryRow(ctx, `SELECT id FROM nationalities WHERE name = $1`, *nationality).Scan(&id)
	if err == pgx.ErrNoRows {
		err = r.db.QueryRow(ctx, `INSERT INTO nationalities(name) VALUES($1) RETURNING id`, *nationality).Scan(&id)
	}
	if err != nil {
		slog.Error("Failed to get or create nationality ID", "nationality", nationality, "error", err)
		return nil, err
	}
	slog.Debug("getOrCreateNationalityID successful", "nationality", nationality, "id", id)
	return &id, nil
}
