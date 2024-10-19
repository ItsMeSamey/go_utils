//! Save key value pairs to a file
//! The keys MUST not have `\n` or `=` in it
//!
//! Expect file to be empty or cursor to be on a new line
//! KvpFile should be initialized with os.File, and when done, os.File should be manually closed

package utils

import (
  "encoding/json"
  "errors"
  "os"
  "strings"
  "unsafe"
)

type KvpFile struct { os.File }

var (
  ErrorNewlineInKey = errors.New("`\\n` found in key")
  ErrorEqualsInKey = errors.New("`=` found in key")
)

func (file KvpFile) Write(k, v string) error {
  if strings.IndexByte(k, '\n') != -1 {
    return ErrorNewlineInKey
  }

  if strings.IndexByte(k, '=') != -1 {
    return ErrorNewlineInKey
  }

  data, err := json.Marshal(v)
  if err != nil { return err }

  _, err = file.File.Write(unsafe.Slice(unsafe.StringData(k), len(k)))
  if err != nil { return err }

  _, err = file.File.Write([]byte{'='})
  if err != nil { return err }

  _, err = file.File.Write(data)
  if err != nil { return err }

  return nil
}

