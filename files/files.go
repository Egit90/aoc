package files

import (
	"bufio"
	"os"
)

type File struct {
	name string
}

func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)

	if err == nil {
		return !info.IsDir(), nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func NewFile(filePath string) *File {

	ok, err := fileExists(filePath)
	if err != nil {
		panic(err)
	}
	if ok {
		return &File{
			name: filePath,
		}
	} else {
		return nil
	}
}
func (file *File) ReadFile() []string {
	var lines []string

	f, err := os.Open(file.name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
