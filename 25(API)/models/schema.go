package models

type ToDoModel struct {
	ToDoID        string `json:"task_id,omitempty" bson:"task_id"`
	Task          string `json:"task,omitempty" bson:"task"`
	TaskCompleted bool   `json:"task_completed,omitempty" bson:"task_completed"`
	TimeStamp     string `json:"time_stamp,omitempty" bson:"time_stamp"`
}
