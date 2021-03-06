package db

import (
	"fmt"
	"time"

	"github.com/elhamza90/lifelog/internal/domain"
)

// Tag Model
type Tag struct {
	ID         domain.TagID
	Name       string
	Expenses   []*Expense  `gorm:"many2many:expense_tags;"`
	Activities []*Activity `gorm:"many2many:activity_tags;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// String returns a one line string representation of a Tag
func (t Tag) String() string { return fmt.Sprintf("[ %d | %s ]", t.ID, t.Name) }

// TableName specifies the name of the table for the tag model
func (t Tag) TableName() string { return "tags" }

// ToDomain converts calling Tag to Domain Tag
func (t Tag) ToDomain() domain.Tag {
	return domain.Tag{
		ID:   t.ID,
		Name: t.Name,
	}
}

// Expense Model
type Expense struct {
	ID         domain.ExpenseID
	Label      string
	Time       time.Time
	Value      float32
	Unit       string
	ActivityID domain.ActivityID // Foreign Key
	Tags       []Tag             `gorm:"many2many:expense_tags;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// ToDomain converts calling Expense to Domain Expense
func (exp Expense) ToDomain() domain.Expense {
	tags := []domain.Tag{}
	for _, t := range exp.Tags {
		tags = append(tags, t.ToDomain())
	}
	return domain.Expense{
		ID:         exp.ID,
		Label:      exp.Label,
		Time:       exp.Time,
		Value:      exp.Value,
		Unit:       exp.Unit,
		ActivityID: exp.ActivityID,
		Tags:       tags,
	}
}

// String returns a one line string representation of a Expense
func (exp Expense) String() string {
	return fmt.Sprintf("[ %d | %s ( %.2f %s) | %s | (%d tags) ]", exp.ID, exp.Label, exp.Value, exp.Unit, exp.Time.Format("2006-01-02 15:04"), len(exp.Tags))
}

// TableName specifies the name of the table for the expense model
func (exp Expense) TableName() string { return "expenses" }

// Activity Model
type Activity struct {
	ID        domain.ActivityID
	Label     string
	Place     string
	Desc      string
	Time      time.Time
	Duration  time.Duration
	Tags      []Tag `gorm:"many2many:activity_tags;"`
	Expenses  []Expense
	CreatedAt time.Time
	UpdatedAt time.Time
}

// String returns a one line string representation of a Activity
func (act Activity) String() string {
	return fmt.Sprintf("[ %d | %s | %s | %s (%s) | (%d tags) ]", act.ID, act.Label, act.Place, act.Time.Format("2006-01-02 15:04"), act.Duration, len(act.Tags))
}

// TableName specifies the name of the table for the activity model
func (act Activity) TableName() string { return "activities" }

// ToDomain converts calling Activity to Domain Activity
func (act Activity) ToDomain() domain.Activity {
	tags := []domain.Tag{}
	for _, t := range act.Tags {
		tags = append(tags, t.ToDomain())
	}
	return domain.Activity{
		ID:       act.ID,
		Label:    act.Label,
		Place:    act.Place,
		Desc:     act.Desc,
		Time:     act.Time,
		Duration: act.Duration,
		Tags:     tags,
	}
}
