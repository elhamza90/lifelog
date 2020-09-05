package rest_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/elhamza90/lifelog/pkg/domain"
	"github.com/labstack/echo/v4"
)

func TestGetAllTags(t *testing.T) {
	expectedCode := http.StatusOK
	path := "/tags"
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	ctx := router.NewContext(req, rec)
	ctx.SetPath(path)
	hnd.GetAllTags(ctx)
	if rec.Code != expectedCode {
		t.Fatalf("\nExpected Code: %d\nReturned Code: %d", expectedCode, rec.Code)
	}
}

func TestGetTagExpenses(t *testing.T) {
	repo.Tags = map[domain.TagID]domain.Tag{
		8987: {ID: 8987, Name: "tag-with-nothing"},
		5555: {ID: 5555, Name: "tag-with-expense"},
	}
	repo.Expenses = map[domain.ExpenseID]domain.Expense{
		1: {
			ID:    1,
			Label: "Expense for tag 5555",
			Value: 55,
			Unit:  "Eu",
			Time:  time.Now().AddDate(0, 0, -2),
			Tags:  []domain.Tag{{ID: 5555, Name: "tag-with-expense"}},
		},
	}
	// Sub-tests definition
	tests := map[string]struct {
		idStr        string
		expectedCode int
	}{
		"Tag with expense": {
			idStr:        "5555",
			expectedCode: http.StatusOK,
		},
		"Tag without expense": {
			idStr:        "8987",
			expectedCode: http.StatusOK,
		},
		"Non-Existing Tag": {
			idStr:        "234234243",
			expectedCode: http.StatusNotFound,
		},
		"Wrong ID": {
			idStr:        "sdfsdf",
			expectedCode: http.StatusBadRequest,
		},
	}
	// Sub-tests execution
	const path string = "/tags/:id/expenses"
	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
	)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req = httptest.NewRequest(http.MethodGet, path, nil)
			rec = httptest.NewRecorder()
			ctx = router.NewContext(req, rec)
			ctx.SetPath(path)
			ctx.SetParamNames("id")
			ctx.SetParamValues(test.idStr)
			hnd.GetTagExpenses(ctx)
			if rec.Code != test.expectedCode {
				body := rec.Body.String()
				t.Fatalf("\nExpected Code: %d\nReturned Code: %v\nReturned Body: %s", test.expectedCode, rec.Code, body)
			}
		})
	}
}

func TestAddTag(t *testing.T) {
	// Init repo with a tag to test duplicate tag name return code.
	repo.Tags = map[domain.TagID]domain.Tag{
		8987: {ID: 8987, Name: "existing-tag"},
	}

	// Sub-tests definition
	tests := map[string]struct {
		json         string
		expectedCode int
	}{
		"Correct": {
			json:         `{"name":"new-tag"}`,
			expectedCode: http.StatusCreated,
		},
		"Duplicate": {
			json:         `{"name":"existing-tag"}`,
			expectedCode: http.StatusBadRequest,
		},
		"Invalid Chars": {
			json:         `{"name":"bad$tag"}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	// Sub-tests execution
	const path string = "/tags"
	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
	)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req = httptest.NewRequest(http.MethodPost, path, strings.NewReader(test.json))
			req.Header.Set("Content-Type", "application/json")
			rec = httptest.NewRecorder()
			ctx = router.NewContext(req, rec)
			ctx.SetPath(path)
			hnd.AddTag(ctx)
			if rec.Code != test.expectedCode {
				t.Fatalf("\nExpected Code: %d\nReturned Code: %v", test.expectedCode, rec.Code)
			}
		})
	}
}

func TestEditTag(t *testing.T) {
	// Init repo with a tag to test duplicate tag name return code.
	repo.Tags = map[domain.TagID]domain.Tag{
		8987: {ID: 8987, Name: "existing-tag"},
		8988: {ID: 8988, Name: "dup-tag"},
	}

	// Sub-tests definition
	tests := map[string]struct {
		json         string
		id           domain.TagID
		expectedCode int
	}{
		"Correct": {
			json:         `{"name":"edited-tag"}`,
			id:           8987,
			expectedCode: http.StatusOK,
		},
		"Non-Existing Tag ID": {
			json:         `{"name":"bla-tag"}`,
			id:           789987, // Random Non-Existing Tag ID!
			expectedCode: http.StatusNotFound,
		},
		"Duplicate": {
			json:         `{"name":"dup-tag"}`, // This name already exists!
			id:           8987,
			expectedCode: http.StatusBadRequest,
		},
		"Invalid Chars": {
			json:         `{"name":"bad$tag"}`,
			id:           8987,
			expectedCode: http.StatusBadRequest,
		},
	}

	// Sub-tests execution
	const path string = "/tags/:id"
	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
	)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req = httptest.NewRequest(http.MethodPut, path, strings.NewReader(test.json))
			req.Header.Set("Content-Type", "application/json")
			rec = httptest.NewRecorder()
			ctx = router.NewContext(req, rec)
			ctx.SetPath(path)
			ctx.SetParamNames("id")
			ctx.SetParamValues(strconv.Itoa(int(test.id)))
			hnd.EditTag(ctx)
			if rec.Code != test.expectedCode {
				body := rec.Body.String()
				t.Fatalf("\nExpected Code: %d\nReturned Code: %v\nReturned Body: %s", test.expectedCode, rec.Code, body)
			}
		})
	}
}

func TestDeleteTag(t *testing.T) {
	// Init repo with a tag to test return code when deleting it
	repo.Tags = map[domain.TagID]domain.Tag{
		8987: {ID: 8987, Name: "tag-with-nothing"},
		9999: {ID: 9999, Name: "tag-with-activity"},
		5555: {ID: 5555, Name: "tag-with-expense"},
	}
	repo.Expenses = map[domain.ExpenseID]domain.Expense{
		1: {
			ID:    1,
			Label: "Expense for tag 5555",
			Value: 55,
			Unit:  "Eu",
			Time:  time.Now().AddDate(0, 0, -2),
			Tags:  []domain.Tag{{ID: 5555, Name: "tag-with-expense"}},
		},
	}
	repo.Activities = map[domain.ActivityID]domain.Activity{
		1: {
			ID:       1,
			Label:    "Activity for tag 9999",
			Time:     time.Now().AddDate(0, 0, -2),
			Duration: time.Duration(time.Hour),
			Tags:     []domain.Tag{{ID: 9999, Name: "tag-with-activity"}},
		},
	}

	// Sub-tests definition
	tests := map[string]struct {
		id           domain.TagID // id of tag to delete
		expectedCode int
	}{
		"Correct": {
			id:           8987,
			expectedCode: http.StatusNoContent,
		},
		"Non-Existing Tag": {
			id:           234234, // Random non existing Tag ID
			expectedCode: http.StatusNotFound,
		},
		"Tag with Expense": {
			id:           5555,
			expectedCode: http.StatusUnprocessableEntity,
		},
		"Tag with Activity": {
			id:           9999,
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	// Sub-tests execution
	const path string = "/tags/:id"
	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
	)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req = httptest.NewRequest(http.MethodDelete, path, nil)
			req.Header.Set("Content-Type", "application/json")
			rec = httptest.NewRecorder()
			ctx = router.NewContext(req, rec)
			ctx.SetPath(path)
			ctx.SetParamNames("id")
			ctx.SetParamValues(strconv.Itoa(int(test.id)))
			hnd.DeleteTag(ctx)
			if rec.Code != test.expectedCode {
				body := rec.Body.String()
				t.Fatalf("\nExpected Code: %d\nReturned Code: %v\nReturned Body: %s", test.expectedCode, rec.Code, body)
			}
		})
	}
}
