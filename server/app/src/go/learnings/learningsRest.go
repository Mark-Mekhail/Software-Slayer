package learnings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"software-slayer/auth"
	"software-slayer/utils"
	"strconv"
	"strings"
)

var learningsService *LearningsService
var tokenService *auth.TokenServiceImpl

// @Summary Create a new learning item
// @Description Add a new learning item for a user
// @Tags Learning Items
// @Accept json
// @Param Authorization header string false "Bearer token"
// @Param learning_item body CreateLearningRequest true "Learning item to add"
// @Success 201
// @Router /learning [post]
func createLearningItem(w http.ResponseWriter, r *http.Request) {
	var createLearningRequest CreateLearningRequest
	if err := utils.Decode(w, r, &createLearningRequest); err != nil {
		return
	}

	userId, err := tokenService.AuthorizeUser(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if err := validateCreateLearningRequest(createLearningRequest); err != nil {
		http.Error(w, fmt.Sprintf("Invalid %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = learningsService.CreateLearning(userId, createLearningRequest.Title, createLearningRequest.Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// @Summary Delete a learning item
// @Description Delete a learning item for a user
// @Tags Learning Items
// @Accept json
// @Param Authorization header string false "Bearer token"
// @Param id path int true "ID of the learning item to delete"
// @Success 204
// @Router /learning/{id} [delete]
func deleteLearningItem(w http.ResponseWriter, r *http.Request) {
	learningId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/learning/"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	userId, err := tokenService.AuthorizeUser(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	learningItemUserId, err := learningsService.GetUserByLearningId(learningId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if userId != learningItemUserId {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = learningsService.DeleteLearning(learningId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Get learning items by user id
// @Description Get all learning items for a user
// @Tags Learning Items
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} GetLearningResponse
// @Router /learning/{user_id} [get]
func getLearningItemsByUserId(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/learning/"))
	if err != nil {
		http.Error(w, "Invalid user id parameter", http.StatusBadRequest)
		return
	}

	skills, err := learningsService.GetLearningsByUserId(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(skills)
}

// @Summary Get learning item categories
// @Description Get all learning item categories
// @Tags Learning Items
// @Produce json
// @Success 200 {array} string
// @Router /learning/categories [get]
func getLearningItemCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(categoriesList)
}

func InitLearningsRest(_learningsService *LearningsService, _tokenService *auth.TokenServiceImpl) {
	learningsService = _learningsService
	tokenService = _tokenService

	http.HandleFunc("GET /learning/", getLearningItemsByUserId)
	http.HandleFunc("GET /learning/categories", getLearningItemCategories)
	http.HandleFunc("POST /learning", createLearningItem)
	http.HandleFunc("DELETE /learning/", deleteLearningItem)
}
