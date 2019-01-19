package ledger

// TransactionBook contents
type TransactionBook struct {
	Amount float64
	Place  string
	Date   string
	Id     int
}

// Ledger where all transaction should be recorded
type Ledger struct {
	User         string
	Transactions []TransactionBook
}
