package db

import (
	"errors"
	"time"

	"github.com/elhamza90/lifelog/internal/domain"
	"github.com/elhamza90/lifelog/internal/store"
	"gorm.io/gorm"
)

// FindActivityByID returns activity with given ID.
// If none is found, returns error
func (repo Repository) FindActivityByID(id domain.ActivityID) (domain.Activity, error) {
	var act Activity
	err := repo.db.First(&act, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = store.ErrActivityNotFound
	}
	return act.ToDomain(), err
}

// SaveActivity stores the given activity in memory and returns created activity
func (repo Repository) SaveActivity(act domain.Activity) (domain.ActivityID, error) {
	tags := []Tag{}
	for _, t := range act.Tags {
		tags = append(tags, Tag{ID: t.ID, Name: t.Name})
	}
	dbAct := Activity{
		ID:    act.ID,
		Label: act.Label,
		Place: act.Place,
		Desc:  act.Desc,
		Time:  act.Time,
		Tags:  tags,
	}
	res := repo.db.Create(&dbAct)
	return domain.ActivityID(dbAct.ID), res.Error
}

// FindActivitiesByTime returns activities
// with Time field greater than or equal to the given time
func (repo Repository) FindActivitiesByTime(t time.Time) ([]domain.Activity, error) {
	res := []Activity{}
	if err := repo.db.Where("time >= ?", t).Order("time DESC").Find(&res).Error; err != nil {
		return []domain.Activity{}, err
	}
	activities := make([]domain.Activity, len(res))
	for i, exp := range res {
		activities[i] = exp.ToDomain()
	}
	return activities, nil
}

// FindActivitiesByTag returns actenses that have the provided tag in their Tags field
func (repo Repository) FindActivitiesByTag(tid domain.TagID) ([]domain.Activity, error) {
	res := []domain.Activity{}
	return res, errNotImplemented
}

// DeleteActivity removes activity with provided ID from memory
func (repo Repository) DeleteActivity(id domain.ActivityID) error {
	return errNotImplemented
}

// EditActivity edits given activity in memory
func (repo Repository) EditActivity(act domain.Activity) error {
	return errNotImplemented
}
