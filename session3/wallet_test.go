package wallet

import "testing"

func TestDeposit(t *testing.T) {
	tests := []struct {
		name   string
		amount Bitcoin
	}{
		{name: "10", amount: Bitcoin(10)},
		{name: "20", amount: Bitcoin(20)},
		{name: "99", amount: Bitcoin(99)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wallet := Wallet{}

			wallet.Deposit(test.amount)

			balance := wallet.Balance()
			want := test.amount

			if balance != want {
				t.Errorf("want a balance of %d, got %d", want, balance)
			}
		})
	}
}

func TestNegativeDeposit(t *testing.T) {
	wallet := Wallet{}

	err := wallet.Deposit(Bitcoin(-5))
	balance := wallet.Balance()
	want := 0
	if balance != 0 {
		t.Errorf("want a balance of %d, got %d", want, balance)
	}
	if err != NegativeDepositError {
		t.Errorf("want error, got nil")
	}
}

func TestWithdraw(t *testing.T) {
	wallet := Wallet{balance: Bitcoin(20)}

	wallet.Withdraw(Bitcoin(15))

	balance := wallet.Balance()
	want := Bitcoin(5)

	if balance != want {
		t.Errorf("want a balance of %d, got %d", want, balance)
	}
}

func TestDisallowOverdraw(t *testing.T) {
	wallet := Wallet{balance: Bitcoin(20)}

	err := wallet.Withdraw(Bitcoin(100))
	balance := wallet.Balance()
	want := Bitcoin(20)

	if balance != want {
		t.Errorf("want a balance of %d, got %d", want, balance)
	}
	if err != OverdrawError {
		t.Errorf("want error, got nil")
	}
}
