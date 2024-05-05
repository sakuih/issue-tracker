package models

import (

)


type Issue struct {
  ID            uint64 `gorm:"primary_key;auto_increment" json:"id"`
  Title         string
  Description   string
  Severity      int8
}


