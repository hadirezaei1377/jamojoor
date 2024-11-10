package banki

import (
	"fmt"
	"sync"
)

// Account struct for holding account details
type Account struct {
	ID      string
	Balance int
	mutex   sync.Mutex
}

// Bank struct for holding multiple accounts
type Bank struct {
	accounts map[string]*Account
	sync.RWMutex
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]*Account),
	}
}

func (b *Bank) CreateAccount(id string, initialBalance int) {
	b.Lock()
	defer b.Unlock()
	b.accounts[id] = &Account{
		ID:      id,
		Balance: initialBalance,
	}
}

// money transfer between two accounts
func (b *Bank) Transfer(fromID, toID string, amount int) error {
	// Lock the accounts to prevent race conditions
	fromAccount, toAccount := b.accounts[fromID], b.accounts[toID]
	fromAccount.mutex.Lock()
	defer fromAccount.mutex.Unlock()
	toAccount.mutex.Lock()
	defer toAccount.mutex.Unlock()

	// Check if there is enough balance
	if fromAccount.Balance < amount {
		return fmt.Errorf("insufficient funds in account %s", fromID)
	}

	// Perform the transfer
	fromAccount.Balance -= amount
	toAccount.Balance += amount

	return nil
}

func (b *Bank) GetBalance(id string) (int, error) {
	b.RLock()
	defer b.RUnlock()
	account, exists := b.accounts[id]
	if !exists {
		return 0, fmt.Errorf("account %s does not exist", id)
	}
	return account.Balance, nil
}

func main() {
	bank := NewBank()

	// Create accounts
	bank.CreateAccount("user1", 1000)
	bank.CreateAccount("user2", 500)

	// Simulate multiple transactions
	var wg sync.WaitGroup
	numTransactions := 10000
	for i := 0; i < numTransactions; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := bank.Transfer("user1", "user2", 1)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	// Wait for all transactions to complete
	wg.Wait()

	// Print final balances
	balance1, _ := bank.GetBalance("user1")
	balance2, _ := bank.GetBalance("user2")
	fmt.Printf("Final Balance of User1: %d\n", balance1)
	fmt.Printf("Final Balance of User2: %d\n", balance2)
}
