package main

import (
	"bufio"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
	"slices"
	"strconv"
)

type WithdrawalRequest struct {
	Amount int
}

func (a WithdrawalRequest) isInvalidAmount() bool {
	return a.Amount < 0 || a.Amount == 1 || a.Amount == 3
}

var validate = validator.New()

func main() {
	fmt.Print("Enter the requested withdrawal amount: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		amount, err := strconv.Atoi(input)
		if err != nil {
			log.Println("Error: invalid value. Please enter a number.")
			return
		}
		request := WithdrawalRequest{Amount: amount}
		if err = validate.Struct(&request); err != nil {
			log.Println("Validation error:", err)
			return
		}
		if request.isInvalidAmount() {
			fmt.Println("Error: the requested amount cannot be processed. Available bank notes: 2, 5, 10, 20, 50, 100, 200.")
			return
		}
		usedBankNotes := CalculateBankNotes(request.Amount)
		fmt.Printf("Requested withdrawal amount: %d\n", request.Amount)
		fmt.Println("Used bank notes:", usedBankNotes)
	}
}

var existingBankNotes = []int{200, 100, 50, 20, 10, 5, 2}

func CalculateBankNotes(requestedAmount int) map[int]int {
	var (
		usedBankNotes                    = map[int]int{200: 0, 100: 0, 50: 0, 20: 0, 10: 0, 5: 0, 2: 0}
		restrictedAmountsForBankNoteFive = []int{6, 8}
		unwithdrawableAmounts            = []int{1, 3}
	)
	for _, bankNote := range existingBankNotes {
		var (
			isAmountZero                       = requestedAmount == 0
			isRestrictedAmountsForBankNoteFive = slices.Contains(restrictedAmountsForBankNoteFive, requestedAmount) && bankNote == 5
		)
		if isAmountZero {
			break
		}
		if isRestrictedAmountsForBankNoteFive {
			continue
		}
		for requestedAmount-bankNote >= 0 {
			if slices.Contains(unwithdrawableAmounts, requestedAmount-bankNote) {
				break
			}
			requestedAmount -= bankNote
			usedBankNotes[bankNote] += 1
		}
	}
	return usedBankNotes
}
