package model

import (
	"database/sql"
	"time"

	"github.com/DragonPow/todo_project/app/todo/api"
)

type Item struct {
	ID        int          `json:"id" gorm:"column:id;primaryKey;not null"`
	IsDone    bool         `json:"is_done" gorm:"column:is_done;not null;default:false"`
	Title     string       `json:"title" gorm:"column:title;not null"`
	TodoID    int          `json:"todo_id" gorm:"column:todo_id;not null"`
	Todo      Todo         `json:"todo" gorm:"foreignKey:TodoID;references:ID"`
	DoneAt    sql.NullTime `json:"done_at" gorm:"column:done_at;autoCreateTime:false"`
	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at;not null;default:now()"`
	UpdateAt  time.Time    `json:"update_at" gorm:"column:update_at;not null;default:now()"`
}

func (Item) TableName() string {
	return "item"
}

func (i *Item) FromApiProto(item *api.Item) {
	i.Title = item.Title
	i.IsDone = item.Completed
	i.Title = item.Title

	// if createdAt, err := time.Parse(time.RFC3339, item.CreatedAt); err == nil {
	// 	i.CreatedAt = createdAt
	// }
	// if updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt); err == nil {
	// 	i.UpdateAt = updatedAt
	// }
	// if doneAt, err := time.Parse(time.RFC3339, item.CompletedAt); err == nil {
	// 	i.DoneAt = sql.NullTime{Time: doneAt, Valid: true}
	// }
}

func (i *Item) ToApiProto() *api.Item {
	var createdAt, updatedAt, completedAt string
	if !i.CreatedAt.IsZero() {
		createdAt = i.CreatedAt.Format(time.RFC3339)
	}
	if !i.UpdateAt.IsZero() {
		updatedAt = i.UpdateAt.Format(time.RFC3339)
	}
	if i.DoneAt.Valid {
		completedAt = i.DoneAt.Time.Format(time.RFC3339)
	}

	return &api.Item{
		Id:          int32(i.ID),
		Title:       i.Title,
		Completed:   i.IsDone,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		CompletedAt: completedAt,
	}
}
