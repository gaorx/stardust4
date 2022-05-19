package sdtemplate

import (
	"bytes"
	"html/template"

	"github.com/gaorx/stardust4/sderr"
)

type htmlExecutor struct {
}

func (te htmlExecutor) Exec(tmpl string, data any) (string, error) {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return "", sderr.Wrap(err, "sdtemplate parse html template error")
	}
	buff := bytes.NewBufferString("")
	err = t.Execute(buff, data)
	if err != nil {
		return "", sderr.Wrap(err, "sdtemplate execute html template error")
	}
	return buff.String(), nil
}

func (te htmlExecutor) MustExec(template string, data any) string {
	r, err := te.Exec(template, data)
	if err != nil {
		panic(sderr.WithStack(err))
	}
	return r
}
