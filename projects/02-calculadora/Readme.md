# Calculator in Go

This project is a simple calculator implemented in Go. It supports basic arithmetic operations such as addition, subtraction, multiplication, and division.

## Installation

To install the calculator, you need to have Go installed on your machine. You can download and install Go from [here](https://golang.org/dl/).

Clone the repository:

```sh
git clone https://github.com/manuonda/go-projects.git 
cd tutorials/calculator-go
```

## Usage
To use the calculator, you can follow these steps:

1. Build the project:

    ```sh
    go build -o calc
    ```

2. Run the calculator with an expression:

    ```sh
    ./calc "3 + 4"
    ```

    This will output:

    ```
    Resultado: 7
    ```

3. If you run the calculator without any arguments, it will start in REPL mode:

    ```sh
    ./calc
    ```

    In REPL mode, you can enter expressions interactively:

    ```
    > 5 * 6
    Resultado: 30
    > 8 / 2
    Resultado: 4
    ```

The calculator supports the following operations:
- Addition (`+`)
- Subtraction (`-`)
- Multiplication (`*`)
- Division (`/`)
- And others