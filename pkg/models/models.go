package models

import (
  "errors"
  "time"
)

type User struct {
  ID             int
  Name           string
  Email          string
  HashedPassword []byte
  Created        time.Time
  Active         bool
}

