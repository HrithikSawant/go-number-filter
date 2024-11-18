
# Number Filtering Program

The **Number Filtering Program** is a Go-based application designed to filter numbers from a list based on various user-defined conditions. It provides a modular and extensible solution to apply different filtering rules such as even, odd, prime, and custom conditions.

---

## **Features**
- Filter numbers based on predefined conditions like:
  - Even numbers
  - Odd numbers
  - Prime numbers
  - Odd prime numbers
  - Numbers matching custom conditions
- Combine conditions to filter numbers that match:
  - **All specified conditions**
  - **Any specified condition**
---

## **Usage Examples**

### **1. Filter Even Numbers**
This will filter all even numbers from the input list.

```go
evenNumbers := filterAll(numbers, isEven)
fmt.Println("Even Numbers:", evenNumbers)
```

**Example Input:** `[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]`  
**Example Output:** `[2, 4, 6, 8, 10]`

---

### **2. Filter Odd Numbers**
Filters all odd numbers from the list.

```go
oddNumbers := filterAll(numbers, isOdd)
fmt.Println("Odd Numbers:", oddNumbers)
```

**Example Input:** `[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]`  
**Example Output:** `[1, 3, 5, 7, 9]`

---

### **3. Filter Prime Numbers**
Filters all prime numbers from the list.

```go
primeNumbers := filterAll(numbers, isPrime)
fmt.Println("Prime Numbers:", primeNumbers)
```

**Example Input:** `[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]`  
**Example Output:** `[2, 3, 5, 7]`

---

### **4. Filter Odd Prime Numbers**
Filters numbers that are both odd and prime.

```go
oddPrimeNumbers := filterAll(numbers, isOdd, isPrime)
fmt.Println("Odd Prime Numbers:", oddPrimeNumbers)
```

**Example Input:** `[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]`  
**Example Output:** `[3, 5, 7]`

---

### **5. Filter Even Numbers or Multiples of 5**
Filters numbers that are either even or multiples of 5.

```go
evenOrMultiplesOf5 := filterAny(numbers, isEven, isMultipleOf(5))
fmt.Println("Even or Multiples of 5:", evenOrMultiplesOf5)
```

**Example Input:** `[1, 2, 3, ..., 20]`  
**Example Output:** `[10, 20]`

---

### **6. Custom Combined Filters**
Filters numbers that match all or any conditions.

#### Match All Conditions:
```go
result := filterAll(numbers, isOdd, greaterThan(5), isMultipleOf(3))
fmt.Println("Matching All Conditions (odd, >5, multiple of 3):", result)
```

**Example Input:** `[1, 2, ..., 20]`  
**Example Output:** `[9, 15]`

#### Match Any Conditions:
```go
result := filterAny(numbers, isPrime, greaterThan(15), isMultipleOf(5))
fmt.Println("Matching Any Conditions (prime, >15, multiple of 5):", result)
```

**Example Input:** `[1, 2, ..., 20]`  
**Example Output:** `[2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20]`

---

## **Installation**

1. Clone the repository:
   ```bash
   $ git clone https://github.com/HrithikSawant/go-number-filtering.git
   $ cd go-number-filtering
   ```

2. Build the program:
   ```bash
   $ go build
   ```

3. Run the program:
   ```bash
   $ ./go-number-filtering
   ```

---

## **Testing**

To ensure functionality, the program is fully tested with predefined and custom cases. Run tests using:

```bash
$ go test ./...
ok  	github.com/HrithikSawant/go-number-filter	(cached)

```

---

## **Extending Functionality**

You can easily add custom filtering logic using Go’s function parameters. Example: Filter numbers greater than a specified value:

```go
func greaterThan(threshold int) func(int) bool {
    return func(number int) bool {
        return number > threshold
    }
}
```

---

## **Conclusion**

The Number Filtering Program is a flexible, modular, and efficient solution for filtering numbers based on multiple conditions. It is designed with extensibility and performance in mind, allowing users to create custom filters and combine them as needed.