package domain

type PersonFilter struct {
	name        string
	surname     string
	patronymic  string
	age         *int
	gender      string
	nationality string
}

func NewPersonFilter(name, surname, patronymic, gender, nationality string, age *int) PersonFilter {
	return PersonFilter{
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		age:         age,
		gender:      gender,
		nationality: nationality,
	}
}

func (f PersonFilter) GetName() string {
	return f.name
}

func (f PersonFilter) GetSurname() string {
	return f.surname
}

func (f PersonFilter) GetPatronymic() string {
	return f.patronymic
}

func (f PersonFilter) GetAge() *int {
	return f.age
}

func (f PersonFilter) GetGender() string {
	return f.gender
}

func (f PersonFilter) GetNationality() string {
	return f.nationality
}
