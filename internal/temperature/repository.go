package temperature

//go:generate moq -out RepositoryMock.go . Repository
type Repository interface {
	Get() (Temperature, error)
}
