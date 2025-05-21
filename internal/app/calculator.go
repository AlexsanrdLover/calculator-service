package app

import (
	"fmt"
	"math"
	"sync"
	"time"
	"calculator-service/internal/domain"
)

type Calculator struct {
	vars      map[string]float64
	varsMutex sync.RWMutex
}

func NewCalculator() *Calculator {
	return &Calculator{
		vars: make(map[string]float64),
	}
}

func (c *Calculator) Calculate(instructions []domain.Instruction) domain.Response {
	var (
		res    domain.Response
		wg     sync.WaitGroup
		resMux sync.Mutex
	)

	for _, instr := range instructions {
		if instr.Type == domain.TypeCalc {
			wg.Add(1)
			go func(instr domain.Instruction) {
				defer wg.Done()
				c.processCalcInstruction(instr)
			}(instr)
		}
	}
	wg.Wait()

	for _, instr := range instructions {
		if instr.Type == domain.TypePrint {
			c.processPrintInstruction(instr, &res, &resMux)
		}
	}

	return res
}

func (c *Calculator) processCalcInstruction(instr domain.Instruction) {
	var result float64
	var err error

	switch instr.Op {
	case domain.OpAdd, domain.OpSub, domain.OpMul, domain.OpDiv, domain.OpPow:
		left, err1 := c.parseOperand(instr.Left)
		right, err2 := c.parseOperand(instr.Right)
		if err1 != nil || err2 != nil {
			return
		}
		result, err = c.binaryOperation(instr.Op, left, right)

	case domain.OpSqrt:
		val, err1 := c.parseOperand(instr.Left)
		if err1 != nil {
			return
		}
		result, err = c.unaryOperation(val)

	default:
		return
	}

	if err == nil {
		c.varsMutex.Lock()
		c.vars[instr.Var] = result
		c.varsMutex.Unlock()
		fmt.Printf("Calculation: %s = %v\n", instr.Var, result)
	} else {
		fmt.Printf("Error in calculation %s: %v\n", instr.Var, err)
	}
}

func (c *Calculator) processPrintInstruction(instr domain.Instruction, res *domain.Response, resMux *sync.Mutex) {
	c.varsMutex.RLock()
	val, ok := c.vars[instr.Var]
	c.varsMutex.RUnlock()

	if ok {
		resMux.Lock()
		res.Items = append(res.Items, domain.PrintResult{
			Var:   instr.Var,
			Value: val,
		})
		resMux.Unlock()
		fmt.Printf("Print: %s = %v\n", instr.Var, val)
	} else {
		fmt.Printf("Print: variable %s not found\n", instr.Var)
	}
}

func (c *Calculator) parseOperand(op interface{}) (float64, error) {
	c.varsMutex.RLock()
	defer c.varsMutex.RUnlock()

	switch v := op.(type) {
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		if val, ok := c.vars[v]; ok {
			return val, nil
		}
		return 0, fmt.Errorf("variable %s not found", v)
	default:
		return 0, fmt.Errorf("unsupported type: %T", op)
	}
}

func (c *Calculator) binaryOperation(op string, a, b float64) (float64, error) {
	time.Sleep(50 * time.Millisecond)

	switch op {
	case domain.OpAdd:
		return a + b, nil
	case domain.OpSub:
		return a - b, nil
	case domain.OpMul:
		return a * b, nil
	case domain.OpDiv:
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	case domain.OpPow:
		return math.Pow(a, b), nil
	default:
		return 0, fmt.Errorf("unknown binary operation")
	}
}

func (c *Calculator) unaryOperation(a float64) (float64, error) {
	time.Sleep(50 * time.Millisecond)

	if a < 0 {
		return 0, fmt.Errorf("square root of negative number")
	}
	return math.Sqrt(a), nil
}