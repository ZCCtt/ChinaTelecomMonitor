package models

import (
	"github.com/golang-module/carbon/v2"
)

type Summary struct {
    // ...
    Use          float64         `gorm:"type:float" json:"use"`
    Total        float64         `gorm:"type:float" json:"total"`
    GeneralUse   float64         `gorm:"type:float" json:"generalUse"`
    GeneralTotal float64         `gorm:"type:float" json:"generalTotal"`
    SpecialUse   float64         `gorm:"type:float" json:"specialUse"`
    SpecialTotal float64         `gorm:"type:float" json:"specialTotal"`
    // ...
    Items        []SummaryItems  `json:"items"`
}

type SummaryItems struct {
    Name  string  `gorm:"type:varchar(1024)" json:"name"`
    Use   float64 `gorm:"type:float" json:"use"`
    Total float64 `gorm:"type:float" json:"total"`
}
