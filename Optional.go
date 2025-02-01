package utils

import (
  "encoding/json"
)

type Optional[T any] struct {
  Val T
  Exists bool
}

func (self Optional[T]) MarshalJSON() ([]byte, error) {
  if self.Exists {
    return json.Marshal(self.Val)
  } else {
    return json.Marshal(nil)
  }
}

func (self *Optional[T]) UnmarshalJSON(data []byte) error {
  self.Exists = !(len(data) == 4 && B2S(data) == "null")

  return json.Unmarshal(data, &self.Val)
}

