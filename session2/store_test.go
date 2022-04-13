package store

import "testing"

func TestRecord(t *testing.T) {
	_ = Record{
		Title:  "Thriller",
		Artist: "Michael Jackson",
		Copies: 1,
	}
}

func TestBuy(t *testing.T) {
	t.Run("decrements number of copies by one", func(t *testing.T) {
		record := Record{Copies: 2}

		result, _ := Buy(record)

		want := 1
		got := result.Copies

		if want != got {
			t.Errorf("expected copies to be decremented to 1, got: %d", got)
		}
	})

	t.Run("informs when no stock available", func(t *testing.T) {
		record := Record{Copies: 0}

		_, err := Buy(record)

		if err == nil {
			t.Error("expected an insufficient stock error, got nil")
		}
	})
}

func TestCalculatePriceWithDiscount(t *testing.T) {
	record := Record{Price: 40, DiscountPercentage: 25}

	got := record.CalculatePriceWithDiscount()

	want := 30

	if got != want {
		t.Errorf("want price of %d, got %d", want, got)
	}
}
