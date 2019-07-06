package input

import (
	"fmt"
	"os"
	"path/filepath"

	"../model"
	"github.com/spf13/afero"
)

type records struct {
	Area []model.Area
}

func ReadAll(fs afero.Fs, root string) error {
	// find all files
	var yamlFiles []os.FileInfo

	err := afero.Walk(fs, root, func(path string, info os.FileInfo, err error) error {
		// no need for directories
		if info.IsDir() {
			return nil
		}

		// match by extension
		ext := filepath.Ext(info.Name())
		if ext == ".yaml" || ext == ".yml" || ext == ".comp" {
			yamlFiles = append(yamlFiles, info)
		}

		// no match found
		return nil
	})
	if err != nil {
		return err
	}

	// iterate over files
	for _, info := range yamlFiles {
		// decode
		fmt.Printf("%s\n", info.Name())

		// return decoded values
	}

	return nil
}
