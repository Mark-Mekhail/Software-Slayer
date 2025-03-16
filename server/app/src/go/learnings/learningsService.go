package learnings

import (
	"software-slayer/db"
)

type LearningsService interface {
	CreateLearning(userId int, title string, category string) error
	DeleteLearning(id int) error
	GetLearningsByUserId(userID int) ([]GetLearningResponse, error)
	GetUserByLearningId(learningId int) (int, error)
}

type LearningsServiceImpl struct {
	db *db.Database
}

func NewLearningsService(db *db.Database) *LearningsServiceImpl {
	return &LearningsServiceImpl{db: db}
}

func (s *LearningsServiceImpl) CreateLearning(userId int, title string, category string) error {
	_, err := s.db.Exec("INSERT INTO user_learning_list (user_id, title, category) VALUES (?, ?, ?)", userId, title, category)
	return err
}

func (s *LearningsServiceImpl) DeleteLearning(id int) error {
	_, err := s.db.Exec("DELETE FROM user_learning_list WHERE id = ?", id)
	return err
}

func (s *LearningsServiceImpl) GetLearningsByUserId(userID int) ([]GetLearningResponse, error) {
	rows, err := s.db.Query("SELECT id, category, title FROM user_learning_list WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	learnings := make([]GetLearningResponse, 0)
	for rows.Next() {
		var learning GetLearningResponse
		if err := rows.Scan(&learning.ID, &learning.Category, &learning.Title); err != nil {
			return nil, err
		}
		learnings = append(learnings, learning)
	}

	return learnings, nil
}

func (s *LearningsServiceImpl) GetUserByLearningId(learningId int) (int, error) {
	var userId int
	err := s.db.QueryRow("SELECT user_id FROM user_learning_list WHERE id = ?", learningId).Scan(&userId)
	return userId, err
}
