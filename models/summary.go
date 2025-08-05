package models

import (
	"github.com/golang-module/carbon/v2"
)

type Summary struct {
	ID           int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string          `gorm:"type:varchar(1024)" json:"username"` // 电信账号名
	  Use          float64         `gorm:"type:float" json:"use"`
    Total        float64         `gorm:"type:float" json:"total"`
    GeneralUse   float64         `gorm:"type:float" json:"generalUse"`
    GeneralTotal float64         `gorm:"type:float" json:"generalTotal"`
    SpecialUse   float64         `gorm:"type:float" json:"specialUse"`
    SpecialTotal float64         `gorm:"type:float" json:"specialTotal"`
	Balance      int64           `gorm:"type:int" json:"balance"`            // 余额
	VoiceUsage   int64           `gorm:"type:int" json:"voiceUsage"`         // 语音使用量
	VoiceAmount  int64           `gorm:"type:int" json:"voiceAmount"`        // 语音总量
	CreateTime   carbon.DateTime `gorm:"type:TIMESTAMP" json:"createTime"`   // 查询时间
	Items        []SummaryItems  `json:"items"`
}

type SummaryItems struct {
	Name  string `gorm:"type:varchar(1024)" json:"name"`
	Use   float64 `gorm:"type:float" json:"use"`
    Total float64 `gorm:"type:float" json:"total"`
}