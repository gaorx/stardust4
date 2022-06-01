package sdjson

import (
	"fmt"
	"io"
)

func Print(v any, pretty bool) {
	if pretty {
		fmt.Print(MarshalPretty(v))
	} else {
		fmt.Print(MarshalStringDef(v, ""))
	}
}

func Println(v any, pretty bool) {
	if pretty {
		fmt.Println(MarshalPretty(v))
	} else {
		fmt.Println(MarshalStringDef(v, ""))
	}
}

func Fprint(w io.Writer, v any, pretty bool) (int, error) {
	if pretty {
		return fmt.Fprint(w, MarshalPretty(v))
	} else {
		return fmt.Fprint(w, MarshalStringDef(v, ""))
	}
}

func Fprintln(w io.Writer, v any, pretty bool) (int, error) {
	if pretty {
		return fmt.Fprintln(w, MarshalPretty(v))
	} else {
		return fmt.Fprintln(w, MarshalStringDef(v, ""))
	}
}
