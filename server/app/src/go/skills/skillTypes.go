package skills

// SkillDB is the struct for the skills table in the database
type SkillDB struct {
	UserID int    `json:"userId"`
	Topic  string `json:"topic"`
}

// CreateSkillRequest is the struct for the request body when creating a new skill
type CreateSkillRequest struct {
	Topic string `json:"topic"`
}

// UpdateSkillRequest is the struct for the request body when updating a skill
type UpdateSkillRequest struct {
	OldTopic     string `json:"oldTopic"`
	UpdatedTopic string `json:"updatedTopic"`
}
