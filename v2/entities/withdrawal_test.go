package entities_test

import (
	"fmt"
	"github.com/Jpgomes06/bank-notes/entities"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestNewWithdrawal(t *testing.T) {
	tests := []struct {
		name            string
		amount          int
		expectedIsValid bool
	}{
		{
			name:            "Valid amount 100",
			amount:          100,
			expectedIsValid: true,
		},
		{
			name:            "Unwithdrawable amount 1",
			amount:          1,
			expectedIsValid: false,
		},
		{
			name:            "Restricted amount 6",
			amount:          6,
			expectedIsValid: true,
		},
		{
			name:            "Zero amount",
			amount:          0,
			expectedIsValid: false,
		},
		{
			name:            "Negative amount",
			amount:          -50,
			expectedIsValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := entities.NewWithdrawal(tt.amount)
			assert.Equal(t, tt.expectedIsValid, w.IsValid())
		})
	}
}

func TestCalculateBankNotes(t *testing.T) {
	tests := []struct {
		name             string
		amount           int
		expectedResponse map[int]int
		expectedError    error
	}{
		{
			name:             "Valid withdrawal request with amount 100",
			amount:           100,
			expectedResponse: map[int]int{100: 1},
			expectedError:    nil,
		},
		{
			name:             "Valid withdrawal request with amount 2222",
			amount:           2222,
			expectedResponse: map[int]int{200: 11, 20: 1, 2: 1},
			expectedError:    nil,
		},
		{
			name:             "Invalid withdrawal request with amount 0",
			amount:           0,
			expectedResponse: nil,
			expectedError:    fmt.Errorf("invalid withdrawal amount"),
		},
		{
			name:             "Invalid withdrawal request with negative amount",
			amount:           -100,
			expectedResponse: nil,
			expectedError:    fmt.Errorf("invalid withdrawal amount"),
		},
		{
			name:             "Restricted amount 8",
			amount:           8,
			expectedResponse: map[int]int{2: 4},
			expectedError:    fmt.Errorf("invalid withdrawal amount"),
		},
		{
			name:             "Unwithdrawable amount 3",
			amount:           3,
			expectedResponse: nil,
			expectedError:    fmt.Errorf("invalid withdrawal amount"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := entities.NewWithdrawal(tt.amount)
			if !w.IsValid() {
				assert.Equal(t, tt.expectedError.Error(), fmt.Sprintf("invalid withdrawal amount"))
				return
			}
			response := w.CalculateBankNotes()
			assert.Equal(t, tt.expectedResponse, response)
		})
	}
}
