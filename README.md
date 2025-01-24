
# Assignment 4A: Simple Library Management System in Go

## Objective:
To build a simple Library Management System (LMS) in Go, demonstrating the use of basic Go features such as structs, interfaces, slices, and error handling. The system will manage both physical books and eBooks, allowing users to add, remove, search, and list books and eBooks in a library.

## Problem Statement:

### Part 1: Basic Library Management System
- **Book Struct**: Define a `Book` struct with attributes for Title, Author, ISBN, and Availability.
- **Library Struct**: Define a `Library` struct that manages a collection of `Book` objects.
- **Main Functionality**: Demonstrate basic operations such as adding, removing, searching for, and listing books.

### Part 2: Enhanced Library Management System
- **EBook Struct**: Define an `EBook` struct that embeds the `Book` struct and adds a `FileSize` attribute.
- **Interface**: Create a `BookInterface` with a `DisplayDetails` method to display book details. Both `Book` and `EBook` must implement this interface.
- **Polymorphism**: Update the `Library` struct to handle both `Book` and `EBook` instances, leveraging polymorphism via the `BookInterface`.

### Part 3: Implementing a Basic Text-Based Menu
- Add a text-based menu to allow users to interact with the system.
- Options should include adding/removing books, searching for books by title, and listing all books/eBooks.
  
## How to Run the Program

### Prerequisites:
- Install **Go** (version 1.18 or later) from [Go's official website](https://golang.org/dl/).
- Set up your Go workspace (ensure that `GOPATH` is correctly configured).

### Running the Application:

1. **Clone the repository** or copy the Go files to your local machine.
2. **Navigate to the directory** containing the Go files in your terminal.
3. Open a terminal and run the Go program using the following command:

   ```bash
   go run main.go
   ```

4. The program will display a menu with the following options:
   - Add a Book/EBook to the library.
   - Remove a Book/EBook by ISBN.
   - Search for Books by Title.
   - List all Books/EBooks.

5. Follow the prompts to interact with the library system.

### Example of Input/Output:

#### Sample Input 1: Adding a Book
```bash
Enter 1 to add a Book/EBook, or 2 to remove a Book/EBook:
1
Is this a Book or EBook? (book/ebook): book
Enter Book Title: Go Programming Basics
Enter Book Author: John Doe
Enter ISBN (unique identifier): 123456789
Is the Book available (true/false): true
Book added successfully!
```

#### Sample Input 2: Searching for a Book by Title
```bash
Enter 1 to add a Book/EBook, or 2 to remove a Book/EBook:
3
Enter the book title to search for: Go
Found Books:
1. Title: Go Programming Basics, Author: John Doe, ISBN: 123456789, Available: true
```

#### Sample Input 3: Listing All Books
```bash
Enter 1 to add a Book/EBook, or 2 to remove a Book/EBook:
4
Listing all Books/EBooks:
1. Title: Go Programming Basics, Author: John Doe, ISBN: 123456789, Available: true
```

### Expected Outputs:
- **Add Book**: When a new book or eBook is added, the program should confirm the addition.
- **Remove Book**: The book is removed if the ISBN exists in the system.
- **Search for Books**: A list of books with titles matching the search query will be displayed.
- **List All Books**: All books in the library, including their availability and eBook file sizes (if applicable), are shown.

## Code Structure

### `book.go`
Defines the `Book` struct and methods for:
- Initializing a book.
- Displaying book details.

### `ebook.go`
Defines the `EBook` struct (which embeds `Book`) and includes:
- An additional `FileSize` attribute.
- An overridden method to display book details (including file size).

### `library.go`
Defines the `Library` struct, with methods for:
- Adding a book.
- Removing a book by ISBN.
- Searching for books by title.
- Listing all books.

### `main.go`
Contains the main program logic, including:
- A text-based menu system.
- User input handling (using `fmt.Scan` or `bufio`).
- Invoking the functions from `library.go`, `book.go`, and `ebook.go`.

### `interfaces.go`
Defines the `BookInterface` interface that both `Book` and `EBook` implement.

## Error Handling
- Duplicate ISBNs are handled by checking if the ISBN exists before adding a new book.
- Invalid user input is handled with error messages and retry options.

## Sample Tasks for Participants:

1. **Add a Book**: 
   - Title: `Go Programming Basics`
   - Author: `John Doe`
   - ISBN: `123456789`
   - Available: `true`
   
2. **Add an EBook**:
   - Title: `Advanced Go Programming`
   - Author: `Jane Smith`
   - ISBN: `987654321`
   - Available: `true`
   - FileSize: `5 MB`
   
3. **Search for Books with "Go" in the Title**:  
   The search should return books with "Go" in their title (e.g., "Go Programming Basics").

4. **List All Books**:  
   The list should display all books and eBooks with their details, including file sizes for eBooks.

5. **Remove a Book by ISBN**:  
   Remove a book using its ISBN and confirm successful removal.

## Submission Requirements:
1. Submit the Go source code files.
2. Include a README file with:
   - Instructions on how to run the program.
   - Examples of input/output for various operations.
3. Provide screenshots of the program's execution demonstrating all functionalities.

## Assessment Criteria:

1. **Completeness**:
   - Does the program include all required functionalities?
   - Are the enhancements (EBook and Interface) correctly implemented?

2. **Code Quality**:
   - Is the code well-organized and easy to read?
   - Are meaningful variable and function names used?

3. **Go Features**:
   - Are Go-specific features like slices, interfaces, and error handling effectively utilized?

4. **Functionality**:
   - Does the program run without errors?
   - Are all features working as expected?

---
