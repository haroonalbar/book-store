# Books

This is a basic CRUD server for Books.

## Technologies Used

- Golang - Programming Language.
- Gorilla Mux - Router for handling requests.
- GORM - ORM for golang.
- Mysql - Database.

## Getting started

1. Setup mysql and create a db.

2. Specify the db_url and port of the application in .env
there is .env.example for reference.

3. To build and get the binary for execution:

```bash
go build -o ./bin/book ./cmd/main
```

4. Run the programme:

```bash
./bin/book
```

5. Apis

- Get all books:
    GET {path}/book/  
- Get book by ID:
    GET {path}/book/{id}  
- Create Book
    POST {path}/book/ with body values of name, author and publisher.
- Delete Book
    DELETE {path}/book/{id}
- Update Book
    PUT {path}/book/{id} with body of updated values of name, author and publisher.