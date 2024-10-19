package utils

import (
  "os"
  "errors"
  "runtime"
)

// This works because "hosmas/utils/env"'s init is called first
var show = os.Getenv("SHOW_ERROR_STACK_TRACES") == "true"

// Max size of the stak trace
const maxTraceLen = 1024 * 16

// error interface
type errorWithStack struct { error }

// init funcction that is called automatically
func init() {
  if !show { runtime.StartTrace() }
}

// Done this way to reduce cost when `show` is false

// Adds stack trace to errors when in dev mode i.e. `ENV="dev"`
var WithStack = func () func (err error) error {
  if !show { return func (err error) error { return err } }

  return func (err error) error {
    if _, ok := err.(errorWithStack); err == nil || ok { return err }
    out := make([]byte, maxTraceLen, maxTraceLen)
    // runtime.Stack is called with all = false to prevent world stopping!
    err = errorWithStack{ errors.New(err.Error() + "\n-- STACK --\n" + string(out[:runtime.Stack(out, false)])) }
    return err
  }
}()

