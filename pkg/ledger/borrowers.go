package ledger

// Borrower struct that defines entities that borrowed money from your account
type Borrower struct {
	Loan   int    `yaml:"loan"`
	Type   string `yaml:"type"` // per year, month, per week, etc
	Person string `yaml:"person"`
}
