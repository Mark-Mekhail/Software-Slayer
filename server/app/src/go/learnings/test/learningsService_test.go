package learnings_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"software-slayer/db"
	"software-slayer/learnings"
)

func setup(t *testing.T) (sqlmock.Sqlmock, *learnings.LearningsServiceImpl) {
	// Create a mock sql.DB object using sqlmock
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	service := learnings.NewLearningsService(db.NewDB(database))

	return mock, service
}

// CreateLearning tests

func TestCreateLearning_Success(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	userId := 1
	title := "Go Programming"
	category := "Languages"

	dbMock.ExpectExec("INSERT INTO user_learning_list").
		WithArgs(userId, title, category).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute
	err := service.CreateLearning(userId, title, category)

	// Verify
	assert.NoError(t, err)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestCreateLearning_DBError(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	userId := 1
	title := "Go Programming"
	category := "Languages"

	dbMock.ExpectExec("INSERT INTO user_learning_list").
		WithArgs(userId, title, category).
		WillReturnError(errors.New("database error"))

	// Execute
	err := service.CreateLearning(userId, title, category)

	// Verify
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestCreateLearning_DuplicateEntry(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	userId := 1
	title := "Go Programming"
	category := "Languages"

	dbMock.ExpectExec("INSERT INTO user_learning_list").
		WithArgs(userId, title, category).
		WillReturnError(errors.New("Error 1062: Duplicate entry"))

	// Execute
	err := service.CreateLearning(userId, title, category)

	// Verify
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Duplicate entry")
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

// DeleteLearning tests

func TestDeleteLearning_Success(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	learningId := 1

	dbMock.ExpectExec("DELETE FROM user_learning_list").
		WithArgs(learningId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Execute
	err := service.DeleteLearning(learningId)

	// Verify
	assert.NoError(t, err)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestDeleteLearning_NotFound(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	learningId := 999

	dbMock.ExpectExec("DELETE FROM user_learning_list").
		WithArgs(learningId).
		WillReturnResult(sqlmock.NewResult(0, 0))

	// Execute
	err := service.DeleteLearning(learningId)

	// Verify - Note: Current implementation doesn't return error for no rows affected
	// This is a design choice that could be changed if needed
	assert.NoError(t, err)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestDeleteLearning_DBError(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	learningId := 1

	dbMock.ExpectExec("DELETE FROM user_learning_list").
		WithArgs(learningId).
		WillReturnError(errors.New("database error"))

	// Execute
	err := service.DeleteLearning(learningId)

	// Verify
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

// GetLearningsByUserId tests

func TestGetLearningsByUserId_Success(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	userId := 1
	expectedLearnings := []learnings.GetLearningResponse{
		{
			ID: 1,
			LearningBase: learnings.LearningBase{
				Title:    "Go Programming",
				Category: "Languages",
			},
		},
		{
			ID: 2,
			LearningBase: learnings.LearningBase{
				Title:    "Docker",
				Category: "Technologies",
			},
		},
	}

	rows := sqlmock.NewRows([]string{"id", "category", "title"})
	for _, learning := range expectedLearnings {
		rows.AddRow(learning.ID, learning.Category, learning.Title)
	}

	dbMock.ExpectQuery("SELECT id, category, title FROM user_learning_list").
		WithArgs(userId).
		WillReturnRows(rows)

	// Execute
	learningItems, err := service.GetLearningsByUserId(userId)

	// Verify
	assert.NoError(t, err)
	assert.Len(t, learningItems, 2)
	assert.Equal(t, expectedLearnings, learningItems)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestGetLearningsByUserId_NoItems(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	userId := 1

	rows := sqlmock.NewRows([]string{"id", "category", "title"})

	dbMock.ExpectQuery("SELECT id, category, title FROM user_learning_list").
		WithArgs(userId).
		WillReturnRows(rows)

	// Execute
	learningItems, err := service.GetLearningsByUserId(userId)

	// Verify
	assert.NoError(t, err)
	assert.Empty(t, learningItems)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestGetLearningsByUserId_DBError(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	userId := 1

	dbMock.ExpectQuery("SELECT id, category, title FROM user_learning_list").
		WithArgs(userId).
		WillReturnError(errors.New("database error"))

	// Execute
	learningItems, err := service.GetLearningsByUserId(userId)

	// Verify
	assert.Error(t, err)
	assert.Nil(t, learningItems)
	assert.Equal(t, "database error", err.Error())
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestGetLearningsByUserId_ScanError(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	userId := 1

	// Create a row with wrong types to cause a scan error
	rows := sqlmock.NewRows([]string{"id", "category", "title"}).
		AddRow("not an int", 123, 456) // ID should be int, not string

	dbMock.ExpectQuery("SELECT id, category, title FROM user_learning_list").
		WithArgs(userId).
		WillReturnRows(rows)

	// Execute
	learningItems, err := service.GetLearningsByUserId(userId)

	// Verify
	assert.Error(t, err)
	assert.Nil(t, learningItems)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

// GetUserByLearningId tests

func TestGetUserByLearningId_Success(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	learningId := 1
	expectedUserId := 5

	rows := sqlmock.NewRows([]string{"user_id"}).
		AddRow(expectedUserId)

	dbMock.ExpectQuery("SELECT user_id FROM user_learning_list").
		WithArgs(learningId).
		WillReturnRows(rows)

	// Execute
	userId, err := service.GetUserByLearningId(learningId)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, expectedUserId, userId)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestGetUserByLearningId_NotFound(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	learningId := 999

	dbMock.ExpectQuery("SELECT user_id FROM user_learning_list").
		WithArgs(learningId).
		WillReturnError(sql.ErrNoRows)

	// Execute
	userId, err := service.GetUserByLearningId(learningId)

	// Verify
	assert.Error(t, err)
	assert.Equal(t, 0, userId)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.NoError(t, dbMock.ExpectationsWereMet())
}

func TestGetUserByLearningId_DBError(t *testing.T) {
	// Setup
	dbMock, service := setup(t)

	learningId := 1

	dbMock.ExpectQuery("SELECT user_id FROM user_learning_list").
		WithArgs(learningId).
		WillReturnError(errors.New("database error"))

	// Execute
	userId, err := service.GetUserByLearningId(learningId)

	// Verify
	assert.Error(t, err)
	assert.Equal(t, 0, userId)
	assert.Equal(t, "database error", err.Error())
	assert.NoError(t, dbMock.ExpectationsWereMet())
}
