package ledger

// PaymentChronology is the type of Chronology for Bills, see the constants below:
type PaymentChronology string

// Payment chronology types
const (
	PaymentChronologyMonthly = PaymentChronology("monthly") // A bill that takes place each month
	PaymentChronologyYearly  = PaymentChronology("yearly")  // A bill that takes place each year
)

// Bill type with fields
type Bill struct {
	Company     string  `yaml:"name"`
	Type        string  `yaml:"type"`
	Amount      float32 `yaml:"amount"`
	Description string  `yaml:"description"`
}
