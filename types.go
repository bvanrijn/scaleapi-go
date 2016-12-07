package scaleapi

import "time"

// Task represents a single task that you create with Scale and is completed by a worker.
type Task struct {
	TaskID      string    `json:"task_id"`
	CompletedAt time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
	CallbackURL string    `json:"callback_url"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Instruction string    `json:"instruction"`
	Urgency     string    `json:"urgency"`
	Params      struct {
		AttachmentType string   `json:"attachment_type"`
		Attachment     string   `json:"attachment"`
		Categories     []string `json:"categories"`
	} `json:"params"`
	Response struct {
		Category string `json:"category"`
	} `json:"response"`
	Metadata struct {
	} `json:"metadata"`
}
