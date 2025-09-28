# ESEP Grade Calculator

This library is a simple utility written in Go to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

# Files
- `go.mod` contains information about this module, `esep/grade-calculator`.
- `grade_calculator.go` contains all business logic for the app, such as:
  - taking in new grade postings (`AddGrade`)
  - averaging grades by weights (`calculateNumericalGrade`)
  - converting numeric average to letter grade (`GetFinalGrade`)
- `grade_calculator_test.go` contains all test cases for `grade_calculator.go`.