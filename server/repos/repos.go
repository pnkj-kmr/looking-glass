package repos

// Model helps to keep record of avaiable model metadeta
type Model interface {
	GetAll() [][]byte
	Get(key string) ([]byte, error)
	Insert(key string, data []byte) error
	Delete(key string) error
}
