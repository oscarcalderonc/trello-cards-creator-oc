package models

import "errors"

type TaskCategory string

const (
	Maintenance TaskCategory = "maintenance"
	Research                 = "research"
	Test                     = "test"
)

type CardType string

const (
	Issue CardType = "issue"
	Bug            = "bug"
	Task           = "task"
)

func (ct CardType) IsValid() error {
	switch ct {
	case Issue, Bug, Task:
		return nil
	}
	return errors.New("Invalid card type")
}

type IssueCard struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type BugCard struct {
	Description string `json:"description" binding:"required"`
}

type TaskCard struct {
	Title    string       `json:"title" binding:"required"`
	Category TaskCategory `json:"category" binding:"required,oneof=maintenance research test"`
}
