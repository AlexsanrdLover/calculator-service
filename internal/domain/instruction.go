package domain

const (
    TypeCalc  = "calc"
    TypePrint = "print"
    OpAdd     = "+"
    OpSub     = "-"
    OpMul     = "*"
    OpDiv     = "/"
    OpPow     = "^"
    OpSqrt    = "V"
)

// Instruction represents a calculation instruction
// swagger:model Instruction
// x-order: ["type", "op", "var", "left", "right"]
type Instruction struct {
    // Type of instruction
    // required: true
    // enum: calc,print
    // example: calc
    // @x-order type,op,var,left,right
    Type  string `json:"type"`
    
    // Operation type (for calc instructions)
    // enum: +,-,*,/,^,√
    // example: +
    Op    string `json:"op,omitempty"`
    
    // Variable name
    // required: true
    // example: x
    Var   string `json:"var"`
    
    // Left operand (number or variable name)
    // required: true
    // example: 10
    Left  interface{} `json:"left"`
    
    // Right operand (number or variable name, omit for unary operations)
    // example: 5
    Right interface{} `json:"right,omitempty"`
}

// Response contains calculation results
type Response struct {
	// List of print results
	Items []PrintResult `json:"items"`
}

// PrintResult represents single print output
type PrintResult struct {
	// Variable name
	Var   string `json:"var"`
	
	// Calculated value
	Value float64  `json:"value"`
}