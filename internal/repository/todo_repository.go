package repository

import (
	"arno/internal/models"
	"time"
)

func (r *Repository) CreateTodo(req *models.TodoRequest, userID int) error {
	todo := models.Todo{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Done,
		CreatedAt:   time.Now(),
	}
	return r.conn.Create(&todo).Error
}

func (r *Repository) GetTodosByUserID(userID int) ([]models.Todo, error) {
	var todos []models.Todo
	if err := r.conn.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *Repository) Update(todoID int, req *models.TodoRequest) error {
	var todo models.Todo

	if err := r.conn.First(&todo, todoID).Error; err != nil {
		return err
	}

	todo.Title = req.Title
	todo.Description = req.Description
	todo.Status = req.Done
	todo.CreatedAt = time.Now()

	return r.conn.Save(&todo).Error
}

func (r *Repository) GetTaskByID(id int) (models.Todo, error) {
	var task models.Todo
	err := r.conn.Where("id = ?", id).First(&task).Error
	return task, err
}

func (r *Repository) Delete(id int) error {
	err := r.conn.Delete(&models.Todo{}, id).Error
	return err
}
