package usecase_test

import (
	"fmt"
	"github.com/Jpgomes06/bank-notes/usecase"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCalculateBankNotes(t *testing.T) {
	tests := []struct {
		name             string
		withdrawalAmount int
		expectedResponse map[int]int
		expectedError    error
	}{
		{
			name:             "Valid withdrawal request with amount 100",
			withdrawalAmount: 100,
			expectedResponse: map[int]int{100: 1},
			expectedError:    nil,
		},
		{
			name:             "Valid withdrawal request with amount 2222",
			withdrawalAmount: 2222,
			expectedResponse: map[int]int{2: 1, 20: 1, 200: 11},
			expectedError:    nil,
		},
		{
			name:             "Invalid withdrawal request with amount 0",
			withdrawalAmount: 0,
			expectedResponse: nil,
			expectedError:    fmt.Errorf("invalid withdrawal amount"),
		},
		{
			name:             "Invalid withdrawal request with negative amount",
			withdrawalAmount: -100,
			expectedResponse: nil,
			expectedError:    fmt.Errorf("invalid withdrawal amount"),
		},
	}

	bankNotesCalculator := usecase.NewBankNotesCalculator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := bankNotesCalculator.CalculateBankNotes(tt.withdrawalAmount)
			if tt.expectedError != nil {
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				return
			}
			assert.Equal(t, tt.expectedResponse, response)
			assert.Equal(t, nil, err)
		})
	}
}
