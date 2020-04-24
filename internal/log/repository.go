package log

//go:generate moq -out RepositoryMock.go . Repository
type Repository interface {
	Get() (Logs, error)
}
