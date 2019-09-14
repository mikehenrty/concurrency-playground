package shared

import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Load(filename string) (files []string) {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	reader:= bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	files = make([]string, 0, 100000)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	return files
}
