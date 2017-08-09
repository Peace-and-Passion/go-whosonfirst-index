package index

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type IndexerFunc func(path string, args ...interface{}) error

type Indexer struct {
	Mode string
	Func IndexerFunc
}

func NewIndexer(mode string, f IndexerFunc) (*Indexer, error) {

	i := Indexer{
		Mode: mode,
		Func: f,
	}

	return &i, nil
}

func (i *Indexer) IndexPath(path string, args ...interface{}) error {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return err
	}

	if i.Mode == "directory" {

		return i.Func(abs_path, args)

	} else if i.Mode == "repo" {

		data := filepath.Join(abs_path, "data")

		_, err = os.Stat(data)

		if err != nil {
			return err
		}

		return i.Func(abs_path, args)

	} else if i.Mode == "filelist" {

		return i.Func(abs_path, args)

	} else if i.Mode == "meta" {

		parts := strings.Split(path, ":")

		if len(parts) != 2 {
			return errors.New("Invalid path declaration for a meta file")
		}

		for _, p := range parts {

			_, err := os.Stat(p)

			if os.IsNotExist(err) {
				return errors.New("Path does not exist")
			}
		}

		meta_file := parts[0]

		// TO DO: append data_root to args...
		// data_root := parts[1]

		return i.Func(meta_file, args)

	} else {

		return i.Func(abs_path, args)
	}

}