package learnings

const (
	Languages    = "Languages"
	Technologies = "Technologies"
	Concepts     = "Concepts"
	Projects     = "Projects"
	Other        = "Other"
)

var categories = []string{Languages, Technologies, Concepts, Projects, Other}

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

