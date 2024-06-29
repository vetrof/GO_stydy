package internal

import "errors"

const DEFAULT_DISCOUNT = 500

type Customer struct {
	Name     string
	Age      int
	Balance  int
	Debt     int
	Discount bool
}

func (c *Customer) WrOffDebt() error {
	if c.Debt >= c.Balance {
		return errors.New("not possible write off")
	}
	c.Balance -= c.Debt
	c.Debt = 0
	return nil
}

func (c *Customer) CalcDiscount() (int, error) {
	if !c.Discount {
		return 0, errors.New("discount not available")
	}
	result := DEFAULT_DISCOUNT - c.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil

}
