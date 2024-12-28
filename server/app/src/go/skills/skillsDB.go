package skills

import (
	"software-slayer/db"
)

func createSkillDB(userId int, topic string) error {
	_, err := db.Exec("INSERT INTO skills (user_id, topic) VALUES (?, ?)", userId, topic)
	return err
}

func deleteSkillDB(userId int, topic string) error {
	_, err := db.Exec("DELETE FROM skills WHERE user_id = ? AND topic = ?", userId, topic)
	return err
}

func updateSkillDB(userId int, oldTopic string, updatedTopic string) error {
	_, err := db.Exec("UPDATE skills SET topic = ? WHERE user_id = ? AND topic = ?", updatedTopic, userId, oldTopic)
	return err
}

func getSkillTopicsByUserIdDB(userID int) ([]string, error) {
	rows, err := db.Query("SELECT topic FROM skills WHERE user_id = ?", userID)
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