package common

import (
	"fmt"
	"golang.org/x/xerrors"
	"runtime"
)

func GetTraceStr(skip int) (output string) {
	pc, _, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		output = fmt.Sprintf("%v:%v", f.Name(), line)
	}
	return
}

func Errorf(err error, format string, a ...interface{}) error {
	msg := format
	if len(a) > 0 {
		msg = fmt.Sprintf(format, a...)
	}
	traceStr := GetTraceStr(2)
	if err != nil {
		return xerrors.Errorf("%+v:\n%+v\nerr: %w", msg, traceStr, err)
	}
	return xerrors.Errorf("%+v\n%+v", msg, traceStr)
}