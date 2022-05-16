package sdexec

import (
	"fmt"
	"time"

	"github.com/gaorx/stardust4/sderr"
	"github.com/mattn/go-shellwords"
)

type Cmd struct {
	Name string
	Args []string
	Env
	Timeout time.Duration
}

func Parse(line string) (*Cmd, error) {
	l, err := shellwords.Parse(line)
	if err != nil {
		return nil, sderr.Wrap(err, "sdexec parse: parse command error")
	}
	c := &Cmd{}
	if len(l) > 0 {
		c.Name = l[0]
		c.Args = l[1:]
	}
	return c, nil
}

func Parsef(format string, a ...any) (*Cmd, error) {
	line := fmt.Sprintf(format, a...)
	return Parse(line)
}

func MustParse(line string) *Cmd {
	cmd, err := Parse(line)
	if err != nil {
		panic(err)
	}
	return cmd
}

func MustParsef(format string, a ...any) *Cmd {
	line := fmt.Sprintf(format, a...)
	return MustParse(line)
}

func (cmd *Cmd) SetDir(wd string) *Cmd {
	cmd.Dir = wd
	return cmd
}

func (cmd *Cmd) SetVar(name, val string) *Cmd {
	cmd.Env.ensure()
	cmd.Vars[name] = val
	return cmd
}

func (cmd *Cmd) AddVars(vars map[string]string) *Cmd {
	if len(vars) == 0 {
		return cmd
	}
	cmd.Env.ensure()
	for name, val := range vars {
		cmd.Vars[name] = val
	}
	return cmd
}

func (cmd *Cmd) SetVars(vars map[string]string) *Cmd {
	cmd.Vars = map[string]string{}
	for name, val := range vars {
		cmd.Vars[name] = val
	}
	return cmd
}

func (cmd *Cmd) SetTimeout(timeout time.Duration) *Cmd {
	cmd.Timeout = timeout
	return cmd
}
