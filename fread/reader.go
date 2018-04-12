package fread

type FileReader interface {
	Next()
	Err() error
	Data() []string
	Close() error
}
