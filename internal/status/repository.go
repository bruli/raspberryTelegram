package status

//go:generate moq -out RepositoryMock.go . Repository
type Repository interface {
	Get() (*Status, error)
}
