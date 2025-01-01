package learnings

const (
	Languages    = "Languages"
	Technologies = "Technologies"
	Concepts     = "Concepts"
	Projects     = "Projects"
	Other        = "Other"
)

type LearningBase struct {
	Title string    `json:"title"`
	Category string `json:"category"`
}

type LearningDB struct {
	ID     int `json:"id"`
	UserID int `json:"userId"`
	LearningBase
}

type CreateLearningRequest struct {
	LearningBase
}

type GetLearningResponse struct {
	ID int `json:"id"`
	LearningBase
}

