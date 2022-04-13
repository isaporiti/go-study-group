package store

import "errors"

type Record struct {
	Title              string
	Artist             string
	Copies             int
	Price              int
	DiscountPercentage int
}

func (r Record) CalculatePriceWithDiscount() int {
	return r.Price - (r.Price * r.DiscountPercentage / 100)
}

func Buy(record Record) (Record, error) {
	if record.Copies == 0 {
		return Record{}, errors.New("not available stock")
	}
	record.Copies--
	return record, nil
}
