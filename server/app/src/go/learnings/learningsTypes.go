package learnings

import (
	"errors"
	"regexp"
)

const (
	Languages    = "Languages"
	Technologies = "Technologies"
	Concepts     = "Concepts"
	Projects     = "Projects"
	Other        = "Other"
)

var titleValidator = regexp.MustCompile(`^.{1,100}$`)
var categoriesMap = map[string]struct{}{ Languages: {}, Technologies: {}, Concepts: {}, Projects: {}, Other: {} }

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

/*
 * Validate the CreateLearningRequest
 * @param createLearningRequest: the CreateLearningRequest to validate
 * @return error: an error if the CreateLearningRequest is invalid
 */
func validateCreateLearningRequest(createLearningRequest CreateLearningRequest) error {
	if _, ok := categoriesMap[createLearningRequest.Category]; !ok {
		return errors.New("category")
	}
	if ok := titleValidator.MatchString(createLearningRequest.Title); !ok {
		return errors.New("title")
	}
	return nil
}

