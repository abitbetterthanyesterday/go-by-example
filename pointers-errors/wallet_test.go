package pointers

import (
	"testing"
)

func TestWallet(t *testing.T){
	assertBalance := func(t testing.TB, wallet Wallet, want Shitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	 assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assertError :=  func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
	
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Shitcoin(10))
		assertBalance(t, wallet, Shitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Shitcoin(20)}
		err := wallet.Withdraw(Shitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Shitcoin(10))
	})

	t.Run("insufficient funds", func(t *testing.T){
		startingBalance := Shitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Shitcoin(100))

		assertError(t, err, ErrorInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)
	})
}

