package learnings_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"software-slayer/learnings"
)

type MockLearningsService struct{}

func (m *MockLearningsService) CreateLearning(userId int, title string, category string) error {
	if title == "invalid" {
		return errors.New("invalid title")
	}
	return nil
}

func (m *MockLearningsService) DeleteLearning(id int) error {
	if id == 999 {
		return errors.New("learning item not found")
	}
	return nil
}

func (m *MockLearningsService) GetLearningsByUserId(userID int) ([]learnings.GetLearningResponse, error) {
	if userID == 999 {
		return nil, errors.New("user not found")
	}

	return []learnings.GetLearningResponse{
		{
			ID: 1,
			LearningBase: learnings.LearningBase{
				Title:    "Go Programming",
				Category: learnings.Languages,
			},
		},
		{
			ID: 2,
			LearningBase: learnings.LearningBase{
				Title:    "Docker",
				Category: learnings.Technologies,
			},
		},
	}, nil
}

func (m *MockLearningsService) GetUserByLearningId(learningId int) (int, error) {
	if learningId == 1 {
		return 1, nil
	}
	if learningId == 2 {
		return 2, nil
	}
	return 0, errors.New("learning item not found")
}

type MockTokenService struct{}

func (m *MockTokenService) GenerateToken(userID int) (string, error) {
	return "mocked_token", nil
}

func (m *MockTokenService) AuthorizeUser(token string) (int, error) {
	if token == "valid_token" {
		return 1, nil
	}
	if token == "user2_token" {
		return 2, nil
	}
	return 0, errors.New("invalid token")
}

var ts *httptest.Server

func TestMain(m *testing.M) {
	mockLearningsService := &MockLearningsService{}
	mockTokenService := &MockTokenService{}
	learnings.InitLearningsRest(mockLearningsService, mockTokenService)
	ts = httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	os.Exit(m.Run())
}

func TestCreateLearningItemSuccess(t *testing.T) {
	requestBody := learnings.CreateLearningRequest{
		LearningBase: learnings.LearningBase{
			Title:    "Learn Go",
			Category: learnings.Languages,
		},
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", ts.URL+"/learning", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "valid_token")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestCreateLearningItemUnauthorized(t *testing.T) {
	requestBody := learnings.CreateLearningRequest{
		LearningBase: learnings.LearningBase{
			Title:    "Learn Go",
			Category: learnings.Languages,
		},
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", ts.URL+"/learning", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "invalid_token")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}
}

func TestCreateLearningItemBadRequest(t *testing.T) {
	requestBody := learnings.CreateLearningRequest{
		LearningBase: learnings.LearningBase{
			Title:    "invalid",
			Category: "InvalidCategory",
		},
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", ts.URL+"/learning", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "valid_token")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestDeleteLearningItemSuccess(t *testing.T) {
	req, _ := http.NewRequest("DELETE", ts.URL+"/learning/1", nil)
	req.Header.Set("Authorization", "valid_token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected %d, got %d", http.StatusNoContent, resp.StatusCode)
	}
}

func TestDeleteLearningItemUnauthorized(t *testing.T) {
	// User 2 trying to delete User 1's item
	req, _ := http.NewRequest("DELETE", ts.URL+"/learning/1", nil)
	req.Header.Set("Authorization", "user2_token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}
}

func TestDeleteLearningItemNotFound(t *testing.T) {
	req, _ := http.NewRequest("DELETE", ts.URL+"/learning/999", nil)
	req.Header.Set("Authorization", "valid_token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}

func TestGetLearningItemsByUserIdSuccess(t *testing.T) {
	resp, err := http.Get(ts.URL + "/learning/1")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var learningItems []learnings.GetLearningResponse
	if err := json.NewDecoder(resp.Body).Decode(&learningItems); err != nil {
		t.Fatal(err)
	}

	if len(learningItems) != 2 {
		t.Errorf("expected %d items, got %d", 2, len(learningItems))
	}
}

func TestGetLearningItemsByUserIdInvalidId(t *testing.T) {
	resp, err := http.Get(ts.URL + "/learning/invalid")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestGetLearningItemsByUserIdNotFound(t *testing.T) {
	resp, err := http.Get(ts.URL + "/learning/999")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}

func TestGetLearningItemCategories(t *testing.T) {
	resp, err := http.Get(ts.URL + "/learning/categories")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var categories []string
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		t.Fatal(err)
	}

	expectedCategories := []string{
		learnings.Languages,
		learnings.Technologies,
		learnings.Concepts,
		learnings.Projects,
		learnings.Other,
	}

	if len(categories) != len(expectedCategories) {
		t.Errorf("expected %d categories, got %d", len(expectedCategories), len(categories))
	}

	// Check that all expected categories are present
	for _, expectedCategory := range expectedCategories {
		found := false
		for _, category := range categories {
			if category == expectedCategory {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected category %s not found", expectedCategory)
		}
	}
}
