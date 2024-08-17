// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGame = "game"

// Game mapped from table <game>
type Game struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`
	UserID     int64     `gorm:"column:user_id;not null;comment:与user表的主键id关联" json:"user_id"`
	Type       int32     `gorm:"column:type;not null;comment:游戏类型，0-选择题，1-判断题，2-填空题，3-滑动选择题，4-图钉答案，5-排序题，6-选择题+音频" json:"type"`
	Question   string    `gorm:"column:question;comment:问题" json:"question"`
	Choice     string    `gorm:"column:choice;comment:选项，json存储" json:"choice"`
	Answer     string    `gorm:"column:answer;comment:问题，json存储" json:"answer"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`
}

// TableName Game's table name
func (*Game) TableName() string {
	return TableNameGame
}
