package models

import (
	"gorm.io/gorm"
	"time"
)

// 比较简单的函数,节约字段

type MiniGlobal struct {
	Model
	CreateBy  int            `json:"createBy" gorm:"index;comment:创建者"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

// 公共字段更丰富的

type RichGlobal struct {
	Model
	ControlBy
	ModelTime
	Desc string `json:"desc" gorm:"size:35;comment:描述信息"` //描述
}
