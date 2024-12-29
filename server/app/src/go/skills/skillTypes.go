package skills

type SkillDB struct {
	UserID int    `json:"userId"`
	Topic  string `json:"topic"`
}

type CreateSkillRequest struct {
	Topic string `json:"topic"`
}

type DeleteSkillRequest struct {
	Topic string `json:"topic"`
}

type UpdateSkillRequest struct {
	OldTopic    string `json:"oldTopic"`
	UpdatedTopic string `json:"updatedTopic"`
}