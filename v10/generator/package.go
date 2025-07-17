// Utility methods for managing and writing generated code
package generator

import (
	"fmt"
	gofmt "go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

type File struct {
	Directory string
	Name      string
}

// Package represents the output package
type Package struct {
	name   string
	header string
	files  map[File]string
}

func NewPackage(name, header string) *Package {
	return &Package{name: name, header: header, files: make(map[File]string)}
}

func (p *Package) WriteFiles(targetDir string) error {
	for name, body := range p.files {
		targetFile := ""
		packageName := p.name
		if name.Directory == "" {
			targetFile = filepath.Join(targetDir, name.Name)
		} else {
			packageName = name.Directory
			targetFile = filepath.Join(targetDir, name.Directory, name.Name)
			err := os.MkdirAll(filepath.Join(targetDir, name.Directory), 0755)
			if err != nil {
				return fmt.Errorf("Error creating directory %v - %v", filepath.Join(targetDir, name.Directory), err)
			}
		}
		fileContent, err := gofmt.Source([]byte(fmt.Sprintf("%v\npackage %v\n%v", p.header, packageName, body)))
		if err != nil {
			return fmt.Errorf("Error writing file %v - %v", targetFile, err)
		}

		err = ioutil.WriteFile(targetFile, fileContent, 0640)
		if err != nil {
			return fmt.Errorf("Error writing file %v - %v", targetFile, err)
		}
	}
	return nil
}

func (p *Package) Files() []string {
	files := make([]string, 0)
	for file, _ := range p.files {
		files = append(files, file.Directory+"/"+file.Name)
	}
	sort.Strings(files)
	return files
}

func (p *Package) HasFile(directory, name string) bool {
	_, ok := p.files[File{Directory: directory, Name: name}]
	return ok
}

func (p *Package) AddFile(directory, name string, body string) {
	p.files[File{Directory: directory, Name: name}] = body
}
