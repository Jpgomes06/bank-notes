
# Bank Notes Script

Bank Notes is a script that allows users to request a withdrawal amount. The script validates the input to ensure that the amount is not negative, for example. It then calculates the smallest possible number of banknotes needed to meet the requested amount and returns the final result.

## Techs

* [Golang](https://go.dev) - Programming language
* [bufio](https://pkg.go.dev/std#bufio) - Package for buffered I/O
* [fmt](https://pkg.go.dev/std#fmt) - Package for formatted I/O
* [github.com/go-playground/validator/v10](https://pkg.go.dev/github.com/go-playground/validator/v10) - Package for data validation
* [log](https://pkg.go.dev/std#log) - Package for logging
* [os](https://pkg.go.dev/std#os) - Package for operating system functions
* [slices](https://pkg.go.dev/std#slices) - Package for slice operations (Go 1.18+)
* [strconv](https://pkg.go.dev/std#strconv) - Package for string conversions

## Approach
- Reading input via terminal
- Input value validations
- `CalculateBankNotes` function centralizing the calculation
- Log messages

1. **Clone the repository**

```bash
 git clone https://github.com/Jpgomes06/bank-notes.git
```

2. **Navigate to the project directory**
```bash
 cd bank-notes
```

3. **Install dependencies**
```bash
 go mod tidy
```

4. **Run the application**
```bash
 go run main.go
```

5. **Run Tests (Optional)**
```bash
 go test ./... -v
```