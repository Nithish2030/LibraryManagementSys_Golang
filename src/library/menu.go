package library

import "fmt"

func Menu(lib *Library) {
	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Search Book by Title")
		fmt.Println("4. List All Books")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title, author, isbn string
			fmt.Println("Enter Title: ")
			fmt.Scanln(&title)
			fmt.Println("Enter Author: ")
			fmt.Scanln(&author)
			fmt.Println("Enter ISBN: ")
			fmt.Scanln(&isbn)
			book := Book{Title: title, Author: author, ISBN: isbn}
			lib.AddBook(book)
		case 2:
			var isbn string
			fmt.Println("Enter ISBN of book to remove: ")
			fmt.Scanln(&isbn)
			lib.RemoveBook(isbn)
		case 3:
			var title string
			fmt.Println("Enter Title: ")
			fmt.Scanln(&title)
			book := lib.SearchBookByTitle(title)
			if book != nil {
				fmt.Printf("Found Book - Title: %s, Author: %s, ISBN: %s\n", book.Title, book.Author, book.ISBN)
			} else {
				fmt.Println("Book not found!")
			}
		case 4:
			lib.ListBooks()
		case 5:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
