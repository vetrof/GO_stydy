package internal

type Partner struct {
	Name    string
	Age     int
	Balance int
	Debt    int
}

func (c *Partner) WrOffDebt() error {
	c.Debt = 0
	return nil
}
