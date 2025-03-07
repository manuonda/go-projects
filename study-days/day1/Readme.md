# Day 1: Go Study Topics

## Topics to Study

### 1. Arrays, Slices, and Maps
#### Concepts
- **Arrays**: Fixed-size sequences of elements of the same type.
- **Slices**: Dynamic-size, flexible view into the elements of an array.
- **Maps**: Unordered collections of key-value pairs.

#### Examples to Practice
- Create an array of integers and iterate over it.
- Create a slice from an array and append elements to it.
- Create a map to store key-value pairs and perform CRUD operations.

### 2. Pointers and Structs
#### Concepts
- **Pointers**: Variables that store the memory address of another variable.
- **Structs**: Composite data types that group together variables under a single name.

#### Examples to Practice
- Create a pointer to an integer and modify its value.
- Define a struct to represent a `Person` with fields like `Name` and `Age`.
- Create instances of the struct and access their fields using pointers.

### 3. Interfaces and Error Handling
#### Concepts
- **Interfaces**: Types that specify a set of method signatures.
- **Error Handling**: Mechanism to handle runtime errors gracefully.

#### Examples to Practice
- Define an interface and implement it with different types.
- Create a function that returns an error and handle it using `if` statements.

### 4. Concurrency Basics: Goroutines and Channels
#### Concepts
- **Goroutines**: Lightweight threads managed by the Go runtime.
- **Channels**: Conduits for communication between goroutines.

#### Examples to Practice
- Launch a goroutine to perform a task concurrently.
- Use channels to send and receive data between goroutines.

## Combined Problem to Solve
Create a program that simulates a simple library system. The system should:
- Use a map to store book information (title, author, and availability).
- Use structs to represent books and users.
- Implement an interface for basic library operations (borrow, return, and search).
- Handle errors such as trying to borrow a book that is not available.
- Use goroutines and channels to handle multiple users borrowing and returning books concurrently.
