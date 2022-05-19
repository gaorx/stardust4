package sdtemplate

import (
	"bytes"
	"text/template"

	"github.com/gaorx/stardust4/sderr"
)

type textExecutor struct {
}

func (te textExecutor) Exec(tmpl string, data any) (string, error) {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return "", sderr.Wrap(err, "sdtemplate parse text template error")
	}
	buff := bytes.NewBufferString("")
	err = t.Execute(buff, data)
	if err != nil {
		return "", sderr.Wrap(err, "sdtemplate execute text template error")
	}
	return buff.String(), nil
}

func (te textExecutor) MustExec(template string, data any) string {
	r, err := te.Exec(template, data)
	if err != nil {
		panic(sderr.WithStack(err))
	}
	return r
}
