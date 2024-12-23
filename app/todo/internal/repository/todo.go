package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/DragonPow/todo_project/app/todo/internal/model"
)

func (r *repo) CreateTodo(ctx context.Context, req *model.Todo) (int, error) {
	if req.IsDone {
		req.DoneAt = sql.NullTime{Time: time.Now(), Valid: true}
	}
	for _, item := range req.Items {
		if item.IsDone {
			item.DoneAt = sql.NullTime{Time: time.Now(), Valid: true}
		}
	}
	tx := r.db.WithContext(ctx)
	return req.ID, tx.Create(req).Error
}

func (r *repo) GetTodoByID(ctx context.Context, id int) (*model.Todo, error) {
	var todo *model.Todo
	return todo, r.db.Preload("Items").First(&todo, id).Error
}

func (r *repo) SearchTodoByTitle(ctx context.Context, title string) ([]*model.Todo, error) {
	var todos []*model.Todo
	return todos, r.db.Where("title LIKE ?", "%"+title+"%").Find(&todos).Error
}

func (r *repo) UpdateTodoByID(ctx context.Context, id int, req *model.Todo) (err error) {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("panic: %v", r)
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit().Error
	}()

	fieldUpdateds := map[string]interface{}{
		"title":       req.Title,
		"description": req.Description,
		"is_done":     req.IsDone,
	}
	if req.IsDone {
		fieldUpdateds["done_at"] = sql.NullTime{Time: time.Now(), Valid: true}
	}

	err = tx.Model(&model.Todo{}).Where("id = ?", id).Updates(fieldUpdateds).Error
	if err != nil {
		return err
	}

	// Update items
	for _, item := range req.Items {
		fieldUpdateds := map[string]interface{}{
			"title":   item.Title,
			"is_done": item.IsDone,
		}
		if item.IsDone {
			fieldUpdateds["done_at"] = sql.NullTime{Time: time.Now(), Valid: true}
		}
		err = tx.Model(&model.Item{}).Where("id = ?", item.ID).Updates(fieldUpdateds).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repo) DeleteTodoByID(ctx context.Context, id int) (err error) {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("panic: %v", r)
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit().Error
	}()

	// Delete items
	err = r.db.Where("todo_id = ?", id).Delete(&model.Item{}).Error
	if err != nil {
		return err
	}

	// Delete todo
	return tx.Where("id = ?", id).Delete(&model.Todo{}).Error
}
