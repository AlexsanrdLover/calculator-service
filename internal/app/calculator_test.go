package app

import (
	"testing"
	"calculator-service/internal/domain"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name         string
		instructions []domain.Instruction
		expected     []domain.PrintResult
	}{
		{
			"Simple addition",
			[]domain.Instruction{
				{Type: domain.TypeCalc, Op: domain.OpAdd, Var: "x", Left: 1.3, Right: 2},
				{Type: domain.TypePrint, Var: "x"},
			},
			[]domain.PrintResult{{Var: "x", Value: 3.3}},
		},
		{
			"Subtraction",
			[]domain.Instruction{
				{Type: domain.TypeCalc, Op: domain.OpSub, Var: "y", Left: 10, Right: 4},
				{Type: domain.TypePrint, Var: "y"},
			},
			[]domain.PrintResult{{Var: "y", Value: 6}},
		},
		{
			"Subtraction",
			[]domain.Instruction{
				{Type: domain.TypeCalc, Op: domain.OpSub, Var: "y", Left: 10, Right: 4},
				{Type: domain.TypePrint, Var: "y"},
			},
			[]domain.PrintResult{{Var: "y", Value: 6}},
		},
		{
			"Division",
			[]domain.Instruction{
				{Type: domain.TypeCalc, Op: domain.OpDiv, Var: "z", Left: 20, Right: 5},
				{Type: domain.TypePrint, Var: "z"},
			},
			[]domain.PrintResult{{Var: "z", Value: 4}},
		},
		{
			"Power",
			[]domain.Instruction{
				{Type: domain.TypeCalc, Op: domain.OpPow, Var: "p", Left: 2, Right: 3},
				{Type: domain.TypePrint, Var: "p"},
			},
			[]domain.PrintResult{{Var: "p", Value: 8}},
		},
		{
			"Square root",
			[]domain.Instruction{
				{Type: domain.TypeCalc, Op: domain.OpSqrt, Var: "s", Left: 25},
				{Type: domain.TypePrint, Var: "s"},
			},
			[]domain.PrintResult{{Var: "s", Value: 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := NewCalculator()
			result := calc.Calculate(tt.instructions)
			
			if len(result.Items) != len(tt.expected) {
				t.Fatalf("Expected %d items, got %d", len(tt.expected), len(result.Items))
			}
			
			for i, item := range result.Items {
				if item != tt.expected[i] {
					t.Errorf("At index %d: expected %v, got %v", i, tt.expected[i], item)
				}
			}
		})
	}
}

func TestErrorCases(t *testing.T) {
	tests := []struct {
		name         string
		instructions []domain.Instruction
	}{
		{
			"Division by zero",
			[]domain.Instruction{
				{Type: domain.TypeCalc, Op: domain.OpDiv, Var: "err", Left: int64(10), Right: int64(0)},
			},
		},
		{
			"Negative sqrt",
			[]domain.Instruction{
				{Type: domain.TypeCalc, Op: domain.OpSqrt, Var: "err", Left: int64(-9)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := NewCalculator()
			result := calc.Calculate(tt.instructions)
			
			if len(result.Items) != 0 {
				t.Errorf("Expected no results for error case, got %d", len(result.Items))
			}
		})
	}
}