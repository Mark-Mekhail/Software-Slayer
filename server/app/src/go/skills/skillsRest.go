package skills

import (
	"encoding/json"
	"net/http"
	"software-slayer/auth"
	"software-slayer/utils"
	"strconv"
)

// @Summary Create a new skill
// @Description Add a new skill for a user
// @Tags Skills
// @Accept json
// @Param skill body string true "Skill topic to add"
// @Success 201
// @Router /skill [post]
func createSkill(w http.ResponseWriter, r *http.Request) {
	var topic string
	if err := utils.Decode(w, r, &topic); err != nil {
		return
	}

	userId, err := auth.AuthorizeUser(r);
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = createSkillDB(userId, topic)
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
// @Param skill body string true "Skill topic to delete"
// @Success 204
// @Router /skill [delete]
func deleteSkill(w http.ResponseWriter, r *http.Request) {
	var topic string
	if err := utils.Decode(w, r, &topic); err != nil {
		return
	}

	userId, err := auth.AuthorizeUser(r);
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
// @Param skill body UpdateSkillRequest true "Old and new skill topics"
// @Success 204
// @Router /skill [put]
func updateSkill (w http.ResponseWriter, r *http.Request) {
	var skill UpdateSkillRequest
	if err := utils.Decode(w, r, &skill); err != nil {
		return
	}

	userId, err := auth.AuthorizeUser(r);
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
// @Param user_id query int true "User ID"
// @Success 200 {array} string
// @Router /skill [get]
func getSkillsByUserId(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
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

func handleSkills(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createSkill(w, r)
	case http.MethodDelete:
		deleteSkill(w, r)
	case http.MethodGet:
		getSkillsByUserId(w, r)
	case http.MethodPut:
		updateSkill(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func InitSkillRoutes() {
	http.HandleFunc("/skill", handleSkills)
}