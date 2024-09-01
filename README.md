# Books

This project is a basic CRUD (Create, Read, Update, Delete) server for managing a collection of books. It is designed to demonstrate the use of Golang for building a RESTful API, utilizing various libraries and tools to handle routing, database interactions, and environment configurations. The server allows users to perform operations such as adding new books, retrieving details of existing books, updating book information, and deleting books from the database.

<!--toc:start-->
- [Technologies Used](#technologies-used)
- [Getting started](#getting-started)
- [APIs](#apis)
- [Project Structure](#project-structure)
<!--toc:end-->

## Technologies Used

- Golang - A statically typed, compiled programming language designed for simplicity and efficiency.
- Gorilla Mux - A powerful URL router and dispatcher for Golang, used for handling incoming HTTP requests and routing them to the appropriate handlers.
- GORM - An Object-Relational Mapping (ORM) library for Golang, which simplifies database interactions by allowing developers to work with database records as if they were regular Go objects.
- MySQL - A widely-used relational database management system, chosen for its reliability and performance.
- Godotenv - A library for loading environment variables from a .env file, making it easier to manage configuration settings for different environments (development, testing, production).

## Getting started

1. Clone the repository

```bash
git clone https://github.com/haroonalbar/book-store.github
cd book-store
```

2. Set up MySQL and create a database.

3. Create a .env file in root directory and specify the database URL and port of the application in a .env file.There is a .env.example file for reference.

4. Build the project to get the binary for execution:

```bash
go build -o ./bin/book ./cmd/main
```

5. Run the program:

```bash
./bin/book
```

## APIs

- Get all books:

```
    GET {path}/book/  
```

- Get book by ID:

```
    GET {path}/book/{id}  
```

- Create Book:

```
    POST {path}/book/
```

```json
{
  "name": "Book Name",
  "author": "Author Name",
  "publication": "Publisher Name"
}
```

- Delete Book:

```
    DELETE {path}/book/{id}
```

- Update Book:

```
    PUT {path}/book/{id}
```

```json
{
  "name": "Updated Book Name",
  "author": "Updated Author Name",
  "publication": "Updated Publisher Name"
}
```

## Project Structure

- cmd/main: Entry point of the application.
- pkg/config: Configuration files and database connection setup.
- pkg/controllers: Handlers for the API endpoints.
- pkg/models: Database models and methods for interacting with the database.
- pkg/routes: Route definitions for the API.
- pkg/utils: Utility functions.
