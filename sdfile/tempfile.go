package sdfile

import (
	"io/ioutil"
	"os"

	"github.com/gaorx/stardust4/sderr"
)

func UseTempFile(dir, pattern string, action func(*os.File)) error {
	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return sderr.Wrap(err, "sdfile.UseTempFile: create temp file error")
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()
	action(f)
	return nil
}

func UseTempDir(dir, pattern string, action func(string)) error {
	name, err := ioutil.TempDir(dir, pattern)
	if err != nil {
		return sderr.Wrap(err, "sdfile.UseTempDir: create temp dir error")
	}
	defer func() {
		_ = os.RemoveAll(name)
	}()
	action(name)
	return nil
}

func UseTempFileForResult[R any](dir, pattern string, action func(*os.File) (R, error)) (R, error) {
	var empty R
	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return empty, sderr.Wrap(err, "sdfile.UseTempFileForResult: create temp file error")
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()
	if r, err := action(f); err != nil {
		return empty, sderr.Wrap(err, "sdfile.UseTempFileForResult: call action error")
	} else {
		return r, nil
	}
}

func UseTempDirForResult[R any](dir, pattern string, action func(string) (R, error)) (R, error) {
	var empty R
	name, err := ioutil.TempDir(dir, pattern)
	if err != nil {
		return empty, sderr.Wrap(err, "sdfile.UseTempDirForResult: create temp dir error")
	}
	defer func() {
		_ = os.RemoveAll(name)
	}()

	if r, err := action(name); err != nil {
		return empty, sderr.Wrap(err, "sdfile.UseTempDirForResult: call action error")
	} else {
		return r, nil
	}
}
