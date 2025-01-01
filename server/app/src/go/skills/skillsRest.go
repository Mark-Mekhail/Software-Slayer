package skills

import (
	"encoding/json"
	"net/http"
	"software-slayer/auth"
	"software-slayer/utils"
	"strconv"
	"strings"
)

// @Summary Create a new skill
// @Description Add a new skill for a user
// @Tags Skills
// @Accept json
// @Param Authorization header string false "Bearer token"
// @Param skill body CreateSkillRequest true "Skill topic to add"
// @Success 201
// @Router /skill [post]
func createSkill(w http.ResponseWriter, r *http.Request) {
	var createSkillRequest CreateSkillRequest
	if err := utils.Decode(w, r, &createSkillRequest); err != nil {
		return
	}

	userId, err := auth.AuthorizeUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = createSkillDB(userId, createSkillRequest.Topic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// @Summary Delete a skill
// @Description Delete a skill for a user
// @Tags Skills
// @Accept json
// @Param Authorization header string false "Bearer token"
// @Param topic path string true "Skill topic to delete"
// @Success 204
// @Router /skill/:topic [delete]
func deleteSkill(w http.ResponseWriter, r *http.Request) {
	topic := strings.TrimPrefix(r.URL.Path, "/skill/")

	userId, err := auth.AuthorizeUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = deleteSkillDB(userId, topic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Update a skill
// @Description Update a skill for a user
// @Tags Skills
// @Accept json
// @Param Authorization header string false "Bearer token"
// @Param skill body UpdateSkillRequest true "Old and new skill topics"
// @Success 204
// @Router /skill [put]
func updateSkill(w http.ResponseWriter, r *http.Request) {
	var skill UpdateSkillRequest
	if err := utils.Decode(w, r, &skill); err != nil {
		return
	}

	userId, err := auth.AuthorizeUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = updateSkillDB(userId, skill.OldTopic, skill.UpdatedTopic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Get skills by user id
// @Description Get all skills for a user
// @Tags Skills
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} string
// @Router /skill/{user_id} [get]
func getSkillsByUserId(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/skill/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	skills, err := getSkillTopicsByUserIdDB(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(skills)
}

/*
 * InitSkillRoutes initializes the skill routes.
 */
func InitSkillRoutes() {
	http.HandleFunc("GET /skill/", getSkillsByUserId)
	http.HandleFunc("POST /skill", createSkill)
	http.HandleFunc("DELETE /skill/", deleteSkill)
	http.HandleFunc("PUT /skill", updateSkill)
}
