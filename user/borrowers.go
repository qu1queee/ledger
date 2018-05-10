package user

// Borrower type with fields
type Borrower struct {
	Loan   int    `yaml:"loan"`
	Type   string `yaml:"type"` // per year, month, per week, etc
	Person string `yaml:"person"`
}
