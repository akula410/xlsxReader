package main

import (
	"archive/zip"
	"errors"
)

func OpenFile(path string) (xlsx *Xlsx, err error) {

	zipReader, err := zip.OpenReader(path)

	if err != nil {
		return nil, err
	}

	files := zipReader.File

	if len(files) == 0 {
		return nil, errors.New("No files in xlsx package")
	}

	xlsx = new(Xlsx)

	xlsx.Files = make(map[string]*zip.File, len(files))

	for _, file := range files {
		xlsx.Files[file.Name]=new(zip.File)
		xlsx.Files[file.Name]=file
	}

	err = xlsx.readRootRels()

	return

}
