package grpc

import (
	"context"
	"calculator-service/api"
	"calculator-service/internal/app"
	"calculator-service/internal/domain"
)

type Server struct {
	api.UnimplementedCalculatorServer
	calculator *app.Calculator
}

func NewServer(calc *app.Calculator) *Server {
	return &Server{calculator: calc}
}

func (s *Server) Calculate(ctx context.Context, req *api.CalculationRequest) (*api.CalculationResponse, error) {
	instructions := make([]domain.Instruction, len(req.Instructions))
	for i, instr := range req.Instructions {
		instructions[i] = domain.Instruction{
			Type:  instr.Type,
			Op:    instr.Op,
			Var:   instr.Var,
			Left:  getOperand(instr.GetLeft()),
			Right: getOperand(instr.GetRight()),
		}
	}

	response := s.calculator.Calculate(instructions)

	grpcResponse := &api.CalculationResponse{}
	for _, item := range response.Items {
		grpcResponse.Items = append(grpcResponse.Items, &api.PrintResult{
			Var:   item.Var,
			Value: item.Value,
		})
	}

	return grpcResponse, nil
}

func getOperand(operand interface{}) interface{} {
	switch v := operand.(type) {
	case *api.Instruction_LeftNum:
		return v.LeftNum
	case *api.Instruction_LeftVar:
		return v.LeftVar
	case *api.Instruction_RightNum:
		return v.RightNum
	case *api.Instruction_RightVar:
		return v.RightVar
	default:
		return 0.0
	}
}