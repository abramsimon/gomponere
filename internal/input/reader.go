package input

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type Reader interface {
	ReadAll(root string) ([]byte, error)
}

type ReaderImpl struct {
	fs afero.Fs
}

func NewReader(fs afero.Fs) *ReaderImpl {
	return &ReaderImpl{
		fs,
	}
}

func (r ReaderImpl) ReadAll(root string) ([]byte, error) {
	var err error

	// find all files
	yamlFiles, err := r.FindFiles(root)
	if err != nil {
		return nil, err
	}

	// read all of the files
	allBytes := []byte{}
	for _, name := range yamlFiles {
		f, err := r.fs.Open(name)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		b, err := afero.ReadAll(f)
		if err != nil {
			return nil, err
		}

		allBytes = append(allBytes, b...)

		// if the last byte is not a newline, append one so the yaml can unmarshal properly
		if !bytes.HasSuffix(allBytes, []byte{'\n'}) {
			allBytes = append(allBytes, byte('\n'))
		}
	}

	return allBytes, nil
}

func (r ReaderImpl) FindFiles(root string) ([]string, error) {
	var files []string
	err := afero.Walk(r.fs, root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// no need for directories
		if info.IsDir() {
			return nil
		}

		// match by extension
		ext := filepath.Ext(info.Name())
		if ext == ".yaml" || ext == ".yml" || ext == ".comp" {
			files = append(files, path)
		}

		// no match found
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}
