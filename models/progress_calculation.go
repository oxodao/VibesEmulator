package models

// Not sure if this is required server-side since the client seems to have a calculation method & so on
type ProgressCalculation struct {
	A    float64 `json:"a"`
	B    float64 `json:"b"`
	Loss int64   `json:"loss"`
}
