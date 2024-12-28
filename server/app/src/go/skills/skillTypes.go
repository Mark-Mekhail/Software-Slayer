package skills

type Skill struct {
	UserID int    `json:"user_id"`
	Topic  string `json:"topic"`
}

type UpdateSkillRequest struct {
	OldTopic    string `json:"old_topic"`
	UpdatedTopic string `json:"updated_topic"`
}