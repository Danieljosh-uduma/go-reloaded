# Go Reloaded Formatter

Go Reloaded Formatter is a simple yet powerful command-line tool written in Go that processes and formats text based on specific tags and rules. It handles numerical base conversions, case adjustments, punctuation placement, and article correcting.

## Features

### Text Modifiers
The formatter identifies tags within parentheses and applies transformations to the preceding word or a specified number of preceding words:

- **Numerical Conversions**:
  - `(hex)`: Converts the preceding hexadecimal string to its decimal equivalent.
  - `(bin)`: Converts the preceding binary string to its decimal equivalent.
- **Case Adjustments**:
  - `(up)`: Converts the preceding word to uppercase.
  - `(low)`: Converts the preceding word to lowercase.
  - `(cap)`: Capitalizes the preceding word (Title Case).
- **Batch Adjustments**:
  - `(up, n)`: Converts the preceding `n` words to uppercase.
  - `(low, n)`: Converts the preceding `n` words to lowercase.
  - `(cap, n)`: Capitalizes the preceding `n` words.

### Punctuation Handling
The tool automatically formats punctuation marks (`.`, `,`, `!`, `?`, `:`, `;`):
- Removes leading spaces before punctuation.
- Ensures a trailing space after punctuation (unless followed by another punctuation mark).

### Article Correction
- Automatically changes "a" to "an" (and "A" to "An") if the following word begins with a vowel (`a`, `e`, `i`, `o`, `u`) or the letter `h`.

### Quote Management
- Properly formats single quotes around text by removing unnecessary spaces.

## Usage

To use the formatter, you need to have Go installed.

```bash
go version 1.22.0
```

Run the program by providing an input file and an output file name as arguments.

```bash
cd main
go run main.go sample.txt result.txt
```

### Example

**Input (`sample.txt`):**
```text
it (cap) was the best of times, it was the worst of times (up) , a orange is ripe .
```

**Output (`result.txt`):**
```text
It was the best of times, it was the worst of TIMES, an orange is ripe.
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://acad.learn2earn.ng/git/duduma/go-reloaded.git
   ```
2. Navigate to the project directory:
   ```bash
   cd go-reloaded
   ```

## Testing

The project includes a suite of unit tests. Run them using:
```bash
go test -v
```