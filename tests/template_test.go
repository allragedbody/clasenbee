package test

import (
	"os"
	"testing"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

// TestMain is a sample to run an endpoint test
func TestTemplate(t *testing.T) {
	sweaters := Inventory{"wool", 17}
	muban := "{{.Count}} items are made of {{.Material}}"
	tmpl, err := template.New("test").Parse(muban) //建立一个模板
	if err != nil {
		panic(err)
	}
	// tmpl.Execute(wr, data)
	err = tmpl.Execute(os.Stdout, sweaters) //将struct与模板合成，合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}
