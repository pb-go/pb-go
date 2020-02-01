package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"testing"
)

type ShowSnippet struct {
	Content string
}

func TestRendingSnippets(t *testing.T) {
	template, err := template.New("test").ParseFiles("showsnip.html.tmpl")
	if err == nil {
		var buffer bytes.Buffer
		err := template.ExecuteTemplate(&buffer, "", ShowSnippet{
			Content: "%bool;\n\n &foo #bar*test <xml>",
		})
		if err == nil {
			file, err := ioutil.ReadFile("expect_showsnip_test.html")
			if err == nil {
				if buffer.String() != string(file) {
					fmt.Println("expect is:")
					fmt.Println(string(file))
					fmt.Println("but actual is:")
					fmt.Println(buffer.String())
					t.Fail()
				}
			}
		}
	}

	if err != nil {
		t.Error(err)
	}
}
