package grpc

import (
	"context"
	"testing"
	"calculator-service/api"
	"calculator-service/internal/app"
)

func TestGRPCServer_Calculate(t *testing.T) {
	calc := app.NewCalculator()
	server := NewServer(calc)

	req := &api.CalculationRequest{
		Instructions: []*api.Instruction{
			{
				Type: "calc",
				Op:   "+",
				Var:  "x",
				Left: &api.Instruction_LeftNum{LeftNum: 1},
				Right: &api.Instruction_RightNum{RightNum: 2},
			},
			{
				Type: "print",
				Var:  "x",
			},
		},
	}

	resp, err := server.Calculate(context.Background(), req)
	if err != nil {
		t.Fatalf("Calculate failed: %v", err)
	}

	if len(resp.Items) != 1 || resp.Items[0].Value != 3 {
		t.Errorf("Expected x=3, got %v", resp.Items)
	}
}