package templates

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"testing"
)

func TestRendingSnippets(t *testing.T) {
	template, err := template.New("test").ParseFiles("showsnip.html.tmpl")
	if err == nil {
		var buffer bytes.Buffer
		err := template.ExecuteTemplate(&buffer, "content", "foobar")
		if err == nil {
			file, err := ioutil.ReadFile("test_expect_showsnip.html")
			if err == nil {
				if buffer.String() != string(file) {
					t.Fail()
				}
			}
		}
	}

	if err != nil {
		t.Error(err)
	}
}
