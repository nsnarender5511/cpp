package core

import (
	"testing"
)

func TestRemoveComments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name: "No comments",
			input: `package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}`,
			expected: `package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}`,
		},
		{
			name: "Single-line comment",
			input: `package main // comment

import "fmt" // another comment

// main function
func main() {
	fmt.Println("Hello, world!") // print greeting
}`,
			expected: `package main 

import "fmt" 


func main() {
	fmt.Println("Hello, world!") 
}`,
		},
		{
			name: "Multi-line comment",
			input: `package main

/* import "fmt"
func another() {}
*/

func main() {
	/* This is a
	   multi-line comment */
	fmt.Println("Hello, world!")
}`,
			expected: `package main



func main() {
	
	fmt.Println("Hello, world!")
}`,
		},
		{
			name: "Mixed comments",
			input: `package main // package declaration

/* This is a block comment
   about imports */
import "fmt" // importing fmt

// main is the entry point
func main() {
	fmt.Println("Hello") /* Print Hello */
}`,
			expected: `package main 


import "fmt" 


func main() {
	fmt.Println("Hello") 
}`,
		},
		{
			name: "Comment in string literal",
			input: `package main

func main() {
	myString := "// This is not a comment"
	fmt.Println(myString)
}`,
			expected: `package main

func main() {
	myString := "// This is not a comment"
	fmt.Println(myString)
}`,
		},
		{
			name: "Comment in rune literal",
			input: `package main

import "fmt"

func main() {
	r := '//'
	fmt.Printf("%c\n", r)
}`,
			expected: `package main

import "fmt"

func main() {
	r := '//'
	fmt.Printf("%c\n", r)
}`,
		},
		{
			name:     "Empty file",
			input:    "",
			expected: "",
		},
		{
			name: "File with only comments",
			input: `// This is a file with only comments.
/* And another one. */`,
			expected: ` `,
		},
		{
			name: "Shebang line",
			input: `#!/usr/bin/env gorun
package main // comment
func main(){}`,
			expected: `#!/usr/bin/env gorun
package main 
func main(){}`,
		},
		{
			name:     "Complex case with CRLF",
			input:    "package main\r\n// comment\r\nimport \"fmt\"\r\n/* block\r\ncomment */\r\nfunc main() {\r\nfmt.Println(\"Hello\") // EOL comment\r\n}",
			expected: "package main\r\n\r\nimport \"fmt\"\r\n\r\nfunc main() {\r\nfmt.Println(\"Hello\") \r\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveComments([]byte(tt.input))
			if (err != nil) != tt.hasError {
				t.Errorf("RemoveComments() error = %v, wantErr %v", err, tt.hasError)
				return
			}
			if string(got) != tt.expected {
				t.Errorf("RemoveComments() got = %q, want %q", string(got), tt.expected)
			}
		})
	}
}
