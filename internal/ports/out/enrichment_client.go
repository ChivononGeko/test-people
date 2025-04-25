package out

type EnrichmentClient interface {
	GetAge(name string) (*int, error)
	GetGender(name string) (*string, error)
	GetNationality(name string) (*string, error)
}
