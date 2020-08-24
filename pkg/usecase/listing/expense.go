package listing

import (
	"sort"
	"time"

	"github.com/elhamza90/lifelog/pkg/domain"
)

// FindExpensesByTime returns expenses with Time field greater than or equal to the given time.
// The returned expenses are ordered from most recent to oldest
// It returns ErrExpenseTimeFuture when given time is future
func (srv Service) FindExpensesByTime(t time.Time) ([]domain.Expense, error) {
	if t.After(time.Now()) {
		return []domain.Expense{}, domain.ErrExpenseTimeFuture
	}
	res, err := srv.repo.FindExpensesByTime(t)
	if err != nil {
		return []domain.Expense{}, err
	}
	// Sort using Time field descendent
	sort.Slice(*res, func(i, j int) bool {
		elemI := (*res)[i]
		elemJ := (*res)[j]
		return elemI.Time.After(elemJ.Time)
	})
	return *res, nil
}

// FindExpensesByTag returns expenses that have the tag with given ID
// in their Tags field
// The returned expenses are ordered from most recent to oldest
// It returns an error if tag with given ID is not found
func (srv Service) FindExpensesByTag(tid domain.TagID) ([]domain.Expense, error) {
	// Check if Tag exists
	if _, err := srv.repo.FindTagByID(tid); err != nil {
		return []domain.Expense{}, err
	}
	// Get Expenses from Repo
	res, err := srv.repo.FindExpensesByTag(tid)
	if err != nil {
		return []domain.Expense{}, err
	}
	// Sort using Time field descendent
	sort.Slice(*res, func(i, j int) bool {
		elemI := (*res)[i]
		elemJ := (*res)[j]
		return elemI.Time.After(elemJ.Time)
	})
	return *res, nil

}
