package models

import "gorm.io/gorm"

// Task represents a task entity in the application.
//
// Fields:
//
// - ID: A unique identifier for the task, automatically managed by GORM.
// - Description: A brief description of the task. This field is stored as text.
// - Note: Additional notes related to the task, also stored as text.
type Task struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Description string `json:"description" gorm:"text"`
	Note        string `json:"note" gorm:"text"`
}
