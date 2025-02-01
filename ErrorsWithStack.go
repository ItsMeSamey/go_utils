//! To use this, wrap every error return with `WithStack`
//! To enable stack traces set `SHOW_ERROR_STACK_TRACES` environment variable to `true`

package utils

import (
  "unsafe"
  "runtime"
)

// when this is true, errors will have stack trace
// changing this to `const show = false` should eliminate any performance penalty
var show = false

// Max size of the stak trace
const maxTraceLen = 1024 * 16

type ErrorWithStack struct {
  current     string
  originalLen int
  original    error
}

func (e ErrorWithStack) Error() string { return e.current }
func (e ErrorWithStack) Unwrap() error { return e.original }
func (e ErrorWithStack) OriginalError() string { return e.current[:e.originalLen] }


func noStack(err error) error { return err }
func withStack(err error) error {
  if err == nil { return nil }
  if _, ok := err.(ErrorWithStack); ok { return err }

  out := make([]byte, maxTraceLen, maxTraceLen)
  originalString := err.Error()
  return ErrorWithStack{
    current: originalString + "\n##-STACK-##\n" + unsafe.String(unsafe.SliceData(out), runtime.Stack(out, false)),
    originalLen: len(originalString),
    original: err,
  }
}

var WithStack = noStack

func SetErrorStackTrace(showTrace bool) {
  if show == showTrace { return }

  show = showTrace

  if show {
    runtime.StartTrace()
    WithStack = withStack
  } else {
    WithStack = noStack
    runtime.StopTrace()
  }
}

