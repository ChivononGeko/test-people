package transport

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"test-people/internal/ports/in"
)

type PersonHandler struct {
	service in.PersonService
}

func NewPersonHandler(service in.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

// CreatePerson godoc
// @Summary      Создать пользователя
// @Description  Добавляет нового пользователя
// @Tags         person
// @Accept       json
// @Produce      json
// @Param        person  body      transport.PersonDTO  true  "Новый пользователь"
// @Success      201
// @Failure      400  {string}  string "Invalid request"
// @Failure      500  {string}  string "Internal error"
// @Router       /person [post]
func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var dto PersonDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		slog.Warn("Invalid request body", slog.String("error", err.Error()))
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	slog.Info("Creating person", slog.String("name", dto.Name), slog.String("surname", dto.Surname))

	err := h.service.AddPerson(r.Context(), dto.Name, dto.Surname, &dto.Patronymic)
	if err != nil {
		slog.Error("Failed to add person", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	slog.Info("Person created successfully")
}

// UpdatePerson godoc
// @Summary      Обновить пользователя
// @Description  Обновляет данные пользователя по ID
// @Tags         person
// @Accept       json
// @Produce      json
// @Param        id      query     int                  true  "ID пользователя"
// @Param        person  body      transport.PersonDTO  true  "Обновлённые данные"
// @Success      200
// @Failure      400  {string}  string "Invalid input"
// @Failure      500  {string}  string "Internal error"
// @Router       /person [put]
func (h *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromQuery(r)
	if err != nil {
		slog.Warn("Invalid ID in query", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dto PersonDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		slog.Warn("Invalid body", slog.String("error", err.Error()))
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	person := dto.ToDomain()

	if err := h.service.UpdatePerson(r.Context(), id, person); err != nil {
		slog.Error("Failed to update person", slog.Int("id", id), slog.String("error", err.Error()))
		http.Error(w, "failed to update person", http.StatusInternalServerError)
		return
	}

	slog.Info("Person updated", slog.Int("id", id))
	w.WriteHeader(http.StatusOK)
}

// DeletePerson godoc
// @Summary      Удалить пользователя
// @Description  Удаляет пользователя по ID
// @Tags         person
// @Produce      json
// @Param        id  query  int  true  "ID пользователя"
// @Success      204
// @Failure      400  {string}  string "Invalid ID"
// @Failure      500  {string}  string "Internal error"
// @Router       /person [delete]
func (h *PersonHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromQuery(r)
	if err != nil {
		slog.Warn("Invalid ID in query", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.DeletePerson(r.Context(), id); err != nil {
		slog.Error("Failed to delete person", slog.Int("id", id), slog.String("error", err.Error()))
		http.Error(w, "failed to delete person", http.StatusInternalServerError)
		return
	}

	slog.Info("Person deleted", slog.Int("id", id))
	w.WriteHeader(http.StatusNoContent)
}

// GetPersonByFilters godoc
// @Summary      Получить пользователей по фильтру
// @Description  Возвращает список пользователей, подходящих под фильтр
// @Tags         person
// @Produce      json
// @Param        name       query  string  false  "Имя"
// @Param        surname    query  string  false  "Фамилия"
// @Param        patronymic query  string  false  "Отчество"
// @Success      200  {array}   transport.PersonDTO
// @Failure      400  {string}  string "Invalid query parameters"
// @Failure      500  {string}  string "Internal error"
// @Router       /person [get]
func (h *PersonHandler) GetPersonByFilters(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		slog.Warn("Method not allowed", slog.String("method", r.Method))
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filter, err := ParsePersonFilterFromQuery(r.URL.Query())
	if err != nil {
		slog.Warn("Invalid query parameters", slog.String("error", err.Error()))
		http.Error(w, "invalid query parameters", http.StatusBadRequest)
		return
	}

	slog.Info("Getting persons by filter")

	persons, err := h.service.GetByFilter(r.Context(), filter)
	if err != nil {
		slog.Error("Failed to get persons", slog.String("error", err.Error()))
		http.Error(w, "failed to get persons", http.StatusInternalServerError)
		return
	}

	var result []PersonDTO
	for _, p := range persons {
		result = append(result, *FromDomain(p))
	}

	slog.Info("Found persons", slog.Int("count", len(result)))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
