package transport

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"test-people/internal/domain"
)

var (
	ErrMissingID = errors.New("missing id parameter")
	ErrInvalidID = errors.New("invalid id")
)

func ParseIDFromQuery(r *http.Request) (int, error) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		return 0, ErrMissingID
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, ErrInvalidID
	}
	return id, nil
}

func ParsePersonFilterFromQuery(q url.Values) (domain.PersonFilter, error) {
	var age *int
	if ageStr := q.Get("age"); ageStr != "" {
		parsedAge, err := strconv.Atoi(ageStr)
		if err != nil {
			return domain.PersonFilter{}, err
		}
		age = &parsedAge
	}

	filters := domain.NewPersonFilter(q.Get("name"), q.Get("surname"), q.Get("patronymic"), q.Get("gender"), q.Get("nationality"), age)

	return filters, nil
}
