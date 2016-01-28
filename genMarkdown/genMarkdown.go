//Golang to Markdown
package genMarkdown

//Dependencies
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//##readLines
/*readLines takes a file as an arguement, and reads the contents into memory*/
func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//##writeMarkdown
/*writeMarkdown takes 2 parameters: text from memory, and a path to a file that it needs to create. It will then translate the contents from the read file into markdown.*/
func WriteMarkdown(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	for _, line := range lines {
		line = CommentScanner(line)
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

//##CommentScanner
/*CommentScanner runs through a string and checks to see if the line is a comment. If so, the function will change the line to markdown.*/
func CommentScanner(line string) string {
	switch {
		case strings.Contains(line, "/*"):
			line = strings.Replace(line, "/*", "", -1)
		case strings.Contains(line, "*/"):
			line = strings.Replace(line, "*/", "", -1)
		case strings.Contains(line, "//####"):
			line = strings.Replace(line, "//####", "####", -1)
		case strings.Contains(line, "//###"):
			line = strings.Replace(line, "//###", "### ", -1)
		case strings.Contains(line, "//##"):
			line = strings.Replace(line, "//##", "## ", -1)
		case strings.Contains(line, "//"):
			line = strings.Replace(line, "//", "# ", -1)
	}
	return line
}
