package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(50))
		err := wallet.Withdraw(Bitcoin(25))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(25))
	})

	t.Run("Withdraw insufficient funds.", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(25))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Error("didn't want error but got one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if want == nil {
		t.Error("wanted an error but didn't get one")
	}

	if want != want {
		t.Errorf("got %q, want %s", got, want)
	}
}
