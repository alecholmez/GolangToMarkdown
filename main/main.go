//Go To Markdown
package main

//###Dependencies
import (
    "fmt"
    "os"
    "alec.holmes/goToMarkdown/genMarkdown"
)

//##Main Function
/*
Initial entrance for the program

Run the program with the following command:
main main.go markdown.md
*/
func main() {
    userFile := os.Args[1]
    createFile := os.Args[2]

    lines, err := genMarkdown.ReadLines(userFile)
    if err != nil {
        fmt.Printf("readLines: %s \n", err)
    }

    if err := genMarkdown.WriteMarkdown(lines, createFile); err != nil {
        fmt.Printf("writeLines: %s \n", err)
    }
}
