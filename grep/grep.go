package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flgs, files []string) []string {
	f := Newflags(flgs, pattern)
	if len(files) > 1 {
		f.multi = true
	}

	res := make([]string, 0)
	for i := range files {
		res = append(res, f.search(files[i])...)
	}

	return res
}


type flags struct {
	tmp 	map[string]bool
	pattern string
	multi   bool
}

func Newflags(flgs []string, pattern string) *flags {
	tmp := make(map[string]bool)
	for i := range flgs {
		tmp[flgs[i]] = true
	}

	return &flags{
		tmp: tmp,
		pattern: pattern,
	}
}

func (f *flags)search(filename string) ( res []string ){
	file, err := os.Open(filename)
	if err != nil {
		panic("file open failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := f.pattern
	for i:=1; scanner.Scan(); i++ {
		line := scanner.Text()

		linecopy := line

		if _,ok := f.tmp["-i"]; ok {
			pattern = strings.ToLower(pattern)
			linecopy = strings.ToLower(line)
		}

		match := strings.Contains(linecopy , pattern)

		if _,ok := f.tmp["-x"]; ok {
			match = linecopy == pattern
		}

		if _,ok := f.tmp["-v"]; ok {
			match = !match
		}
		// -l -n
		if match {
			if _,ok := f.tmp["-n"]; ok {
				line = fmt.Sprintf("%d:%s", i, line)
			}
			if f.multi {
				line = fmt.Sprintf("%s:%s", filename, line)
			}
			if _,ok := f.tmp["-l"]; ok {
				res = append(res, filename)
				break
			}

			res = append(res, line)
		}

	}

	return
}