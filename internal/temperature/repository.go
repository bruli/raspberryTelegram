package temperature

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	Get() (Temperature, error)
}
