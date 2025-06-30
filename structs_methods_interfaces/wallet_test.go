package structsmethodsinterfaces

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{Bitcoin(10)}
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Error("wanted an error but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("Deposit test function", func(t *testing.T) {

		wallet.Deposit(10)
		got := wallet.Balance()
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		wanted := Bitcoin(20)

		if got != wanted {
			t.Errorf("got: %s, wanted: %s", got, wanted)
		}
	})

	t.Run("Withdraw test function", func(t *testing.T) {
		wallet.Withdraw(Bitcoin(5))
		got := wallet.Balance()
		fmt.Printf("address of balance in test is %s \n", wallet.balance)
		wanted := Bitcoin(15)

		if got != wanted {
			t.Errorf("got: %s, wanted: %s", got, wanted)
		}
	})

	t.Run("Withdraw insuffficent funds test function", func(t *testing.T) {
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, Bitcoin(10))
	})

}
