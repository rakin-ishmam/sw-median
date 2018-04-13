package fread

import (
	"encoding/csv"
	"os"
)

type rcsv struct {
	f      *os.File
	reader *csv.Reader
	res    []string
	err    error
}

func (r *rcsv) Next() {
	r.res, r.err = r.reader.Read()
}

func (r *rcsv) Data() []string {
	return r.res
}

func (r *rcsv) Err() error {
	return r.err
}

func (r *rcsv) Close() error {
	return r.f.Close()
}

// NewCsv returns instance for csv reader
func NewCsv(floc string) (FileReader, error) {
	f, err := os.Open(floc)
	if err != nil {
		return nil, err
	}

	return &rcsv{
		f:      f,
		reader: csv.NewReader(f),
	}, nil
}
