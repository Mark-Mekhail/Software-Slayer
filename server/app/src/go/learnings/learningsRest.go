package learnings

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"software-slayer/auth"
	"software-slayer/utils"
)

var learningsService LearningsService
var tokenService auth.TokenService

// @Summary Create a new learning item
// @Description Add a new learning item for a user
// @Tags Learning Items
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param learning_item body CreateLearningRequest true "Learning item to add"
// @Success 201 {object} map[string]string "Learning item created"
// @Failure 400 {object} utils.ErrorResponse "Invalid learning item data"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Server error"
// @Router /learning [post]
func createLearningItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var createLearningRequest CreateLearningRequest
	if err := utils.Decode(w, r, &createLearningRequest); err != nil {
		return
	}

	userId, err := tokenService.AuthorizeUser(r.Header.Get("Authorization"))
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or missing authentication token")
		return
	}

	if err := validateCreateLearningRequest(createLearningRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid %s", err.Error()))
		return
	}

	log.Printf("Creating learning item '%s' in category '%s' for user ID: %d",
		createLearningRequest.Title, createLearningRequest.Category, userId)

	err = learningsService.CreateLearning(ctx, userId, createLearningRequest.Title, createLearningRequest.Category)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			utils.RespondWithError(w, http.StatusConflict, "This learning item already exists for your account")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create learning item")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Learning item created successfully"})
}

// @Summary Delete a learning item
// @Description Delete a learning item for a user
// @Tags Learning Items
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "ID of the learning item to delete"
// @Success 204 "No Content"
// @Failure 400 {object} utils.ErrorResponse "Invalid ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 404 {object} utils.ErrorResponse "Learning item not found"
// @Failure 500 {object} utils.ErrorResponse "Server error"
// @Router /learning/{id} [delete]
func deleteLearningItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	learningId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/learning/"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid learning item ID")
		return
	}

	userId, err := tokenService.AuthorizeUser(r.Header.Get("Authorization"))
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or missing authentication token")
		return
	}

	learningItemUserId, err := learningsService.GetUserByLearningId(ctx, learningId)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Learning item not found")
		return
	}

	if userId != learningItemUserId {
		utils.RespondWithError(w, http.StatusUnauthorized, "You don't have permission to delete this learning item")
		return
	}

	log.Printf("Deleting learning item ID: %d for user ID: %d", learningId, userId)

	err = learningsService.DeleteLearning(ctx, learningId)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete learning item")
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}

// @Summary Get learning items by user id
// @Description Get all learning items for a user
// @Tags Learning Items
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} GetLearningResponse
// @Failure 400 {object} utils.ErrorResponse "Invalid user ID"
// @Failure 404 {object} utils.ErrorResponse "User not found"
// @Failure 500 {object} utils.ErrorResponse "Server error"
// @Router /learning/{user_id} [get]
func getLearningItemsByUserId(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	userID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/learning/"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	log.Printf("Fetching learning items for user ID: %d", userID)

	learningItems, err := learningsService.GetLearningsByUserId(ctx, userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve learning items")
		return
	}

	log.Printf("Found %d learning items for user ID: %d", len(learningItems), userID)
	utils.RespondWithJSON(w, http.StatusOK, learningItems)
}

// @Summary Get learning item categories
// @Description Get all learning item categories
// @Tags Learning Items
// @Produce json
// @Success 200 {array} string
// @Router /learning/categories [get]
func getLearningItemCategories(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, categoriesList)
}

// InitLearningsRest initializes the learning REST endpoints
func InitLearningsRest(_learningsService LearningsService, _tokenService auth.TokenService) {
	learningsService = _learningsService
	tokenService = _tokenService

	http.HandleFunc("GET /learning/", getLearningItemsByUserId)
	http.HandleFunc("GET /learning/categories", getLearningItemCategories)
	http.HandleFunc("POST /learning", createLearningItem)
	http.HandleFunc("DELETE /learning/", deleteLearningItem)

	log.Println("Learning REST endpoints initialized")
}
