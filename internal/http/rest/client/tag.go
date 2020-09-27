package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/elhamza90/lifelog/internal/domain"
	"github.com/elhamza90/lifelog/internal/http/rest/server"
)

// PostTag sends a POST request with refresh token and
// gets a new Jwt Access Token
func PostTag(payload server.JSONReqTag, token string) (domain.TagID, error) {
	// Marshall Tag to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}
	// Send HTTP Request
	const path string = url + "/tags"
	requestBody := bytes.NewBuffer(jsonPayload)
	req, err := http.NewRequest("POST", path, requestBody)
	if err != nil {
		return 0, err
	}
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	// Read Response
	responseCode := resp.StatusCode
	responseBody, err := readResponseBody(resp.Body)
	if err != nil {
		return 0, err
	}
	// Check Response Code
	if responseCode != http.StatusCreated {
		return 0, fmt.Errorf("error posting new tag:\n\t- code: %d\n\t- body: %s", responseCode, responseBody)
	}
	// Extract ID created Tag
	respObj := struct {
		ID domain.TagID `json:"id"`
	}{}
	if err := json.Unmarshal(responseBody, &respObj); err != nil {
		return 0, err
	}
	return respObj.ID, nil
}

// UpdateTag sends a POST request with refresh token and
// gets a new Jwt Access Token
func UpdateTag(payload server.JSONReqTag, token string) error {
	// Marshall Tag to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	// Send HTTP Request
	path := url + "/tags/" + strconv.Itoa(int(payload.ID))
	requestBody := bytes.NewBuffer(jsonPayload)
	req, err := http.NewRequest("PUT", path, requestBody)
	if err != nil {
		return err
	}
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// Read Response
	responseCode := resp.StatusCode
	responseBody, err := readResponseBody(resp.Body)
	if err != nil {
		return err
	}
	// Check Response Code
	if responseCode != http.StatusOK {
		return fmt.Errorf("error updating tag:\n\t- code: %d\n\t- body: %s", responseCode, responseBody)
	}
	return nil
}

// DeleteTag sends a POST request with refresh token and
// gets a new Jwt Access Token
func DeleteTag(id domain.TagID, token string) error {
	// Send HTTP Request
	path := url + "/tags/" + strconv.Itoa(int(id))
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// Read Response
	responseCode := resp.StatusCode
	responseBody, err := readResponseBody(resp.Body)
	if err != nil {
		return err
	}
	// Check Response Code
	if responseCode != http.StatusNoContent {
		return fmt.Errorf("error deleting tag:\n\t- code: %d\n\t- body: %s", responseCode, responseBody)
	}
	return nil
}

// FetchTags sends a GET request to fetch all tags
func FetchTags(token string) ([]server.JSONRespListTag, error) {
	// Send HTTP Request
	const path string = url + "/tags"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return []server.JSONRespListTag{}, err
	}
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []server.JSONRespListTag{}, err
	}
	// Read Response
	responseCode := resp.StatusCode
	responseBody, err := readResponseBody(resp.Body)
	if err != nil {
		return []server.JSONRespListTag{}, err
	}
	// Check Response Code
	if responseCode != http.StatusOK {
		return []server.JSONRespListTag{}, fmt.Errorf("error posting new tag:\n\t- code: %d\n\t- body: %s", responseCode, responseBody)
	}
	// Extract Tags
	var tags []server.JSONRespListTag
	if err := json.Unmarshal(responseBody, &tags); err != nil {
		return []server.JSONRespListTag{}, err
	}
	return tags, nil
}

// FetchTagExpenses sends a GET request to fetch expenses with given tag id
func FetchTagExpenses(id domain.TagID, token string) ([]server.JSONRespListExpense, error) {
	// Send HTTP Request
	path := url + "/tags/" + strconv.Itoa(int(id)) + "/expenses"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return []server.JSONRespListExpense{}, err
	}
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []server.JSONRespListExpense{}, err
	}
	// Read Response
	responseCode := resp.StatusCode
	responseBody, err := readResponseBody(resp.Body)
	if err != nil {
		return []server.JSONRespListExpense{}, err
	}
	// Check Response Code
	if responseCode != http.StatusOK {
		return []server.JSONRespListExpense{}, fmt.Errorf("error fetching expenses:\n\t- code: %d\n\t- body: %s", responseCode, responseBody)
	}
	// Extract Expenses
	var expenses []server.JSONRespListExpense
	if err := json.Unmarshal(responseBody, &expenses); err != nil {
		return []server.JSONRespListExpense{}, err
	}
	return expenses, nil
}

// FetchTagActivities sends a GET request to fetch activities with given tag id
func FetchTagActivities(id domain.TagID, token string) ([]server.JSONRespListActivity, error) {
	// Send HTTP Request
	path := url + "/tags/" + strconv.Itoa(int(id)) + "/activities"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return []server.JSONRespListActivity{}, err
	}
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []server.JSONRespListActivity{}, err
	}
	// Read Response
	responseCode := resp.StatusCode
	responseBody, err := readResponseBody(resp.Body)
	if err != nil {
		return []server.JSONRespListActivity{}, err
	}
	// Check Response Code
	if responseCode != http.StatusOK {
		return []server.JSONRespListActivity{}, fmt.Errorf("error fetching activities:\n\t- code: %d\n\t- body: %s", responseCode, responseBody)
	}
	// Extract Activities
	var activities []server.JSONRespListActivity
	if err := json.Unmarshal(responseBody, &activities); err != nil {
		return []server.JSONRespListActivity{}, err
	}
	return activities, nil
}
