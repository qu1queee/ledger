package user

import "time"

// Ledger type with fields
type Ledger struct {
	Admin   string `yaml:"user"`
	Salary  int    `yaml:"salary"`
	Access  time.Time
	Clients []Borrower `yaml:"borrowers"`
	Bills   []Bill     `yaml:"bills"`
}
