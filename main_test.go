package main_test

import (
	"github.com/Jpgomes06/bank-notes"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSuccessTransactionIntegration(t *testing.T) {
	tests := []struct {
		name              string
		withdrawalRequest main.WithdrawalRequest
		expectedResponse  map[int]int
	}{
		{
			name: "Valid withdrawal request with amount 100",
			withdrawalRequest: main.WithdrawalRequest{
				Amount: 100,
			},
			expectedResponse: map[int]int{2: 0, 5: 0, 10: 0, 20: 0, 50: 0, 100: 1, 200: 0},
		},
		{
			name: "Valid withdrawal request with amount 2222",
			withdrawalRequest: main.WithdrawalRequest{
				Amount: 2222,
			},
			expectedResponse: map[int]int{2: 1, 5: 0, 10: 0, 20: 1, 50: 0, 100: 0, 200: 11},
		},
		{
			name: "Valid withdrawal request with amount 30333",
			withdrawalRequest: main.WithdrawalRequest{
				Amount: 30333,
			},
			expectedResponse: map[int]int{2: 4, 5: 1, 10: 0, 20: 1, 50: 0, 100: 1, 200: 151},
		},
		{
			name: "Valid withdrawal request with amount 400451",
			withdrawalRequest: main.WithdrawalRequest{
				Amount: 400451,
			},
			expectedResponse: map[int]int{2: 3, 5: 1, 10: 0, 20: 2, 50: 0, 100: 0, 200: 2002},
		},
		{
			name: "Valid withdrawal request with amount 5005556",
			withdrawalRequest: main.WithdrawalRequest{
				Amount: 5005556,
			},
			expectedResponse: map[int]int{2: 3, 5: 0, 10: 0, 20: 0, 50: 1, 100: 1, 200: 25027},
		},
		{
			name: "Valid withdrawal request with amount 60006667",
			withdrawalRequest: main.WithdrawalRequest{
				Amount: 60006667,
			},
			expectedResponse: map[int]int{2: 1, 5: 1, 10: 1, 20: 0, 50: 1, 100: 0, 200: 300033},
		},
		{
			name: "Valid withdrawal request with amount 700007759",
			withdrawalRequest: main.WithdrawalRequest{
				Amount: 700007759,
			},
			expectedResponse: map[int]int{2: 2, 5: 1, 10: 0, 20: 0, 50: 1, 100: 1, 200: 3500038},
		},
		{
			name: "Valid withdrawal request with amount 8000880829",
			withdrawalRequest: main.WithdrawalRequest{
				Amount: 8000880829,
			},
			expectedResponse: map[int]int{2: 2, 5: 1, 10: 0, 20: 1, 50: 0, 100: 0, 200: 40004404},
		},
		{
			name: "Valid withdrawal request with amount 11",
			withdrawalRequest: main.WithdrawalRequest{
				Amount: 11,
			},
			expectedResponse: map[int]int{2: 3, 5: 1, 10: 0, 20: 0, 50: 0, 100: 0, 200: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := main.CalculateBankNotes(tt.withdrawalRequest.Amount)
			assert.Equal(t, tt.expectedResponse, response)
		})
	}
}
