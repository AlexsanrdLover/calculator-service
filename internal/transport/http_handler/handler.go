package http_handler

import (
	"encoding/json"
	"net/http"

	"calculator-service/internal/app"
	"calculator-service/internal/domain"
)

// Handler handles HTTP requests for calculator
type Handler struct {
	calculator *app.Calculator
}

// NewHandler creates new HTTP handler
func NewHandler(calculator *app.Calculator) *Handler {
	return &Handler{calculator: calculator}
}

// Calculate processes instructions
// @Summary Process calculation instructions
// @Description Accepts JSON array of instructions (calc/print) and returns results
// @Tags calculator
// @Accept json
// @Produce json
// @Param request body []domain.Instruction true "Calculation instructions"
// @Success 200 {object} domain.Response "Calculation results"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /calculate [post]
func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	var instructions []domain.Instruction
	if err := json.NewDecoder(r.Body).Decode(&instructions); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.calculator.Calculate(instructions)
	
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}