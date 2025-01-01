package learnings

import (
	"software-slayer/db"
)

/*
 * createLearningDB creates a new learning item in the database
 * @param userId: the user id of the user who is creating the learning item
 * @param title: the title of the learning item
 * @param category: the category of the learning item
 * @return error: an error if the query fails
 */
func createLearningDB(userId int, title string, category string) error {
	_, err := db.Exec("INSERT INTO user_learning_list (user_id, title, category) VALUES (?, ?, ?)", userId, title, category)
	return err
}

/*
 * deleteLearningDB deletes a learning item from the database
 * @param id: the id of the learning item to delete
 * @return error: an error if the query fails
 */
func deleteLearningDB(id int) error {
	_, err := db.Exec("DELETE FROM user_learning_list WHERE id = ?", id)
	return err
}

/*
 * getLearningsByUserIdDB gets all the learning items of a user from the database
 * @param userID: the user id of the user whose learning items are being fetched
 * @return []GetLearningResponse: a list of learning items
 * @return error: an error if the query fails
 */
func getLearningsByUserIdDB(userID int) ([]GetLearningResponse, error) {
	rows, err := db.Query("SELECT id, category, title FROM user_learning_list WHERE user_id = ?", userID)
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

/*
 * getUserByLearningIdDB gets the user id of the user who created a learning item
 * @param learningId: the id of the learning item
 * @return int: the user id of the user who created the learning item
 * @return error: an error if the query fails
 */
func getUserByLearningIdDB(learningId int) (int, error) {
	var userId int
	err := db.QueryRow("SELECT user_id FROM user_learning_list WHERE id = ?", learningId).Scan(&userId)
	return userId, err
}