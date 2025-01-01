package skills

import (
	"software-slayer/db"
)

/*
 * createSkillDB creates a new skill in the database
 * @param userId: the user id of the user who is creating the skill
 * @param topic: the topic of the skill
 * @return error: an error if the query fails
 */
func createSkillDB(userId int, topic string) error {
	_, err := db.Exec("INSERT INTO user_learning_list (user_id, topic) VALUES (?, ?)", userId, topic)
	return err
}

/*
 * deleteSkillDB deletes a skill from the database
 * @param userId: the user id of the user who is deleting the skill
 * @param topic: the topic of the skill
 * @return error: an error if the query fails
 */
func deleteSkillDB(userId int, topic string) error {
	_, err := db.Exec("DELETE FROM user_learning_list WHERE user_id = ? AND topic = ?", userId, topic)
	return err
}

/*
 * updateSkillDB updates a skill in the database
 * @param userId: the user id of the user who is updating the skill
 * @param oldTopic: the old topic of the skill
 * @param updatedTopic: the updated topic of the skill
 * @return error: an error if the query fails
 */
func updateSkillDB(userId int, oldTopic string, updatedTopic string) error {
	_, err := db.Exec("UPDATE user_learning_list SET topic = ? WHERE user_id = ? AND topic = ?", updatedTopic, userId, oldTopic)
	return err
}

/*
 * getSkillTopicsByUserIdDB gets all the skill topics of a user from the database
 * @param userID: the user id of the user whose skills are being fetched
 * @return []string: a list of skill topics
 * @return error: an error if the query fails
 */
func getSkillTopicsByUserIdDB(userID int) ([]string, error) {
	rows, err := db.Query("SELECT topic FROM user_learning_list WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []string
	for rows.Next() {
		var skill string
		if err := rows.Scan(&skill); err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	return skills, nil
}
