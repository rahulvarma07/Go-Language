package models

type ToDoModel struct {
	ToDoID        string `json:"task_id" bson:"task_id"`
	Task          string `json:"task" bson:"task"`
	TaskCompleted bool   `json:"task_completed" bson:"task_completed"`
	TimeStamp     string `json:"time_stamp" bson:"time_stamp"`
}
