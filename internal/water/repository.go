package water

//go:generate moq -out RepositoryMock.go . Repository
type Repository interface {
	Execute(zone string, seconds uint8) error
}
