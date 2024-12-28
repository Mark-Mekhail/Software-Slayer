package skills

type SkillDB struct {
	UserID int    `json:"user_id"`
	Topic  string `json:"topic"`
}

type CreateSkillRequest struct {
	Topic string `json:"topic"`
}

type DeleteSkillRequest struct {
	Topic string `json:"topic"`
}

type UpdateSkillRequest struct {
	OldTopic    string `json:"old_topic"`
	UpdatedTopic string `json:"updated_topic"`
}