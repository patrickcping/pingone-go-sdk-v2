package fsconfig

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

func ReadFile(filePath string) (filedata map[string]string, err error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return ParseFiledata(f)
}

func ParseFiledata(f io.Reader) (filedata map[string]string, err error) {
	filedata = make(map[string]string)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile(`^(PINGONE_\w+)={1}(.+)$`)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if matched := r.MatchString(line); matched {
			matches := r.FindAllStringSubmatch(line, -1)
			if len(matches) == 1 && len(matches[0]) == 3 {
				filedata[matches[0][1]] = matches[0][2]
			}
		}
	}

	return filedata, nil
}
