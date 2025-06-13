package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// User represents a user with their relevant data for airdrop eligibility.
type User struct {
	ID              string  `json:"id"`
	TokenBalance    float64 `json:"token_balance"`
	Transactions    int     `json:"transactions"`
	IsRegistered    bool    `json:"is_registered"`
	LastTransaction string  `json:"last_transaction"` // Format: YYYY-MM-DD
}

// EligibilityCriteria defines the rules for airdrop eligibility.
type EligibilityCriteria struct {
	MinTokenBalance             float64 `json:"min_token_balance"`
	MinTransactions             int     `json:"min_transactions"`
	MustBeRegistered            bool    `json:"must_be_registered"`
	MaxDaysSinceLastTransaction int     `json:"max_days_since_last_transaction"`
}

// LoadUsers loads user data from a JSON file.
func LoadUsers(filePath string) ([]User, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var users []User
	if err := json.Unmarshal(bytes, &users); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return users, nil
}

// IsEligible checks if a user meets the defined eligibility criteria.
func IsEligible(user User, criteria EligibilityCriteria) bool {
	// Check minimum token balance
	if user.TokenBalance < criteria.MinTokenBalance {
		return false
	}

	// Check minimum transactions
	if user.Transactions < criteria.MinTransactions {
		return false
	}

	// Check registration status
	if criteria.MustBeRegistered && !user.IsRegistered {
		return false
	}

	// Check last transaction date (simplified for this example, assuming date parsing and comparison)
	// In a real-world scenario, you'd parse LastTransaction into a time.Time object
	// and compare it with the current date or a cutoff date.
	// For simplicity, we'll just check if it's not empty if a max days criteria is set.
	if criteria.MaxDaysSinceLastTransaction > 0 && user.LastTransaction == "" {
		return false
	}

	return true
}

func main() {
	// Define eligibility criteria
	criteria := EligibilityCriteria{
		MinTokenBalance:             100.0,
		MinTransactions:             5,
		MustBeRegistered:            true,
		MaxDaysSinceLastTransaction: 30, // Example: last transaction within 30 days
	}

	// Load user data
	users, err := LoadUsers("users.json")
	if err != nil {
		fmt.Printf("Error loading users: %v\n", err)
		return
	}

	fmt.Println("Checking airdrop eligibility...")
	fmt.Printf("Criteria: %+v\n", criteria)

	eligibleUsers := []User{}
	for _, user := range users {
		if IsEligible(user, criteria) {
			eligibleUsers = append(eligibleUsers, user)
		}
	}

	fmt.Println("\nEligible Users:")
	if len(eligibleUsers) == 0 {
		fmt.Println("No users are eligible based on the current criteria.")
	} else {
		for _, user := range eligibleUsers {
			fmt.Printf("- ID: %s, Token Balance: %.2f, Transactions: %d, Registered: %t\n",
				user.ID, user.TokenBalance, user.Transactions, user.IsRegistered)
		}
	}
}
