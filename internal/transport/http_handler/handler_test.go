package http_handler

import (
	"fmt"
	"bytes"
	"net/http/httptest"
	"testing"
	"calculator-service/internal/app"
)

func TestCalculate_Success(t *testing.T) {
	calc := app.NewCalculator()
	handler := NewHandler(calc)

	jsonBody := `[
		{"type":"calc","op":"+","var":"x","left":1,"right":2},
		{"type":"print","var":"x"}
	]`

	req := httptest.NewRequest("POST", "/calculate", bytes.NewBufferString(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.Calculate(rr, req)
	fmt.Println(rr.Body.String())
}