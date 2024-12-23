package model

import (
	"database/sql"
	"time"

	"github.com/DragonPow/todo_project/app/todo/api"
)

type Todo struct {
	ID          int            `json:"id" gorm:"column:id;primaryKey;not null"`
	Title       string         `json:"title" gorm:"column:title;not null"`
	Description sql.NullString `json:"description" gorm:"column:description;"`
	UserId      int            `json:"user_id" gorm:"column:user_id;not null"`
	IsDone      bool           `json:"is_done" gorm:"column:is_done;not null;default:false"`
	DoneAt      sql.NullTime   `json:"done_at" gorm:"column:done_at;autoCreateTime:false"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at;not null;default:now()"`
	UpdateAt    time.Time      `json:"update_at" gorm:"column:update_at;not null;default:now()"`
	Items       []Item         `json:"items" gorm:"foreignKey:TodoID;references:ID"`
}

func (Todo) TableName() string {
	return "todo"
}

func (t *Todo) FromApiProto(todo *api.Todo) {
	t.Title = todo.Title
	t.Description = sql.NullString{String: todo.Description, Valid: true}
	t.IsDone = todo.Completed
	// if createdAt, err := time.Parse(time.RFC3339, todo.CreatedAt); err == nil {
	// 	t.CreatedAt = createdAt
	// }
	// if updatedAt, err := time.Parse(time.RFC3339, todo.UpdatedAt); err == nil {
	// 	t.UpdateAt = updatedAt
	// }
	// if doneAt, err := time.Parse(time.RFC3339, todo.CompletedAt); err == nil {
	// 	t.DoneAt = sql.NullTime{Time: doneAt, Valid: true}
	// }

	for _, apiItem := range todo.Items {
		item := Item{}
		item.FromApiProto(apiItem)
		t.Items = append(t.Items, item)
	}
}

func (t *Todo) ToApiProto() *api.Todo {
	items := make([]*api.Item, len(t.Items))
	for i, item := range t.Items {
		items[i] = item.ToApiProto()
	}

	var createdAt, updatedAt, completedAt string
	if !t.CreatedAt.IsZero() {
		createdAt = t.CreatedAt.Format(time.RFC3339)
	}
	if !t.UpdateAt.IsZero() {
		updatedAt = t.UpdateAt.Format(time.RFC3339)
	}
	if t.DoneAt.Valid {
		completedAt = t.DoneAt.Time.Format(time.RFC3339)
	}

	return &api.Todo{
		Id:          int32(t.ID),
		Title:       t.Title,
		Description: t.Description.String,
		Completed:   t.IsDone,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		CompletedAt: completedAt,
		Items:       items,
	}
}
