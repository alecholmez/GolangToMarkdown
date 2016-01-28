# README

This tool is used to generate markdown from a go source code file for easy documentation of code.

Currently the best way to generate the proper markdown is to follow these guidelines:

*Single-line comments are translated to H1 tags*

*Single-line comments with ## after are translated to H2 tags*
`//##`

*Single-line comments with ### after are translated to H3 tags*
`//###`

*Single-line comments with ### after are translated to H4 tags*
`//####`

*Multi-line comments are translated to P tags*

*Function, type, and import statements are to be wrapped in code block syntax*
#### Currently this is not implemented and is soon to come

## Usage

`cd` into the package directory

`cd` into `/genMarkdown` and run `go build`

`cd` into `/main` and run `go install`

To run the program use this command:

```
main (sourceCodeFile.go) (output.md)
```
IE:
```
main main.go markdown.md
```

A markdown file will be placed in the directory of the main folder
