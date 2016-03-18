package lazy

import (
	"html/template"
	"bytes"
	"testing"
)

func fty() *template.Template {
	return template.Must(template.New("").Parse("{{ .Name }}"))
}

var lazyTemplate = NewLazyTemplate(fty)

func renderTemplate(t *template.Template) string {
	if t == nil {
		panic("gotcha")
	}
	writer := bytes.Buffer{}
	t.Execute(&writer, struct{ Name string }{"x"})
	value := string(writer.Bytes())
	return value
}

func render(lazy *LazyTemplate) string {
	return renderTemplate(lazy.Value())
}

func Test_Render(t *testing.T) {
	tt := fty()
	x := renderTemplate(tt)
	if x != "x" {
		t.Error("nope")
	} // else { t.Logf("temaplet: %v", x) }
	x = renderTemplate(tt)
	if x != "x" {
		t.Error("nope")
	}// else { t.Logf("temaplet: %v", x)}
}

func Test_LazyTemplate(t *testing.T) {

	if value := render(lazyTemplate); value != "x" {
		t.Error("No Good")
	} // else {	t.Logf("Ok: %v", value) }
	if value := render(lazyTemplate); value != "x" {
		t.Error("No Good 2")
	} // else {	t.Logf("Ok: %v", value) }
}
