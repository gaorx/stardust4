package sdtemplate

type TemplateExecutor interface {
	Exec(template string, data any) (string, error)
	MustExec(template string, data any) string
}
