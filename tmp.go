package main

import (
	"bytes"
	"fmt"
	"text/template"
)

const docTmp = `
### {{.QuestionDetail.QuestionId}}. {{.QuestionDetail.QuestionTitle}}
{{.QuestionDetail.Content}}

<button class="section" target="solution" show="Show the solution" hide="Hide the solution"></button>

<!--sec data-title="{{.QuestionDetail.QuestionTitle}}" data-id="solution" data-show=false ces-->

[import](../../problems/{{.PadLeftId}}-{{.QuestionDetail.QuestionTitleSlug}}/{{.QuestionDetail.QuestionTitleSlug}}.go)

<!--endsec-->
`
const strSummaryTmp = `
* [{{.QuestionDetail.QuestionId}}. {{.QuestionDetail.QuestionTitle}}]({{.PadLeftId}}-{{.QuestionDetail.QuestionTitleSlug}}.md)
`

const problemTmp = `
package problem{{.PadLeftId}}

{{.CodeDefinition.DefaultCode}}
`
const problemTestTmp = `
package problem{{.PadLeftId}}

import (
	"testing"
)

func Test{{.MetaData.Name}}(t *testing.T) {
	
}
`

func getDoc(req TmpModel) string {
	t := template.Must(template.New("doc").Parse(docTmp))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, req); err != nil {
		panic(err)
	}
	s := buf.String()
	return fmt.Sprintln(s)
}

func getSummary(req TmpModel) string {
	t := template.Must(template.New("summary").Parse(strSummaryTmp))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, req); err != nil {
		panic(err)
	}
	s := buf.String()
	return s
}

func getProblem(req TmpModel) string {
	t := template.Must(template.New("problem").Parse(problemTmp))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, req); err != nil {
		panic(err)
	}
	s := buf.String()
	return s
}

func getProblemTest(req TmpModel) string {
	t := template.Must(template.New("test").Parse(problemTestTmp))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, req); err != nil {
		panic(err)
	}
	s := buf.String()
	return s
}
