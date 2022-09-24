package daocore

import (
  "database/sql/driver"
  "encoding/json"
  "errors"
)


type JsonType map[string]interface{}

func (j *JsonType) Scan(v interface{}) error {
  if w, ok := v.([]byte); ok {
    return json.Unmarshal(w, j)
  }
  return errors.New("invalid json format")
}

func (j JsonType) Value() (driver.Value, error) {
  return json.Marshal(j)
}