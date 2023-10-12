# GoCRUDApp

## Description
GoCRUDApp is a simple CRUD application written in Go. It allows users to create, read, update, and delete posts. Designed with simplicity in mind, it serves as an excellent starting point for beginners to familiarize themselves with web development in Go and for building more complex applications in the future.

## Features
- **Create Post**: Add a new post with a title and content.
- **Read Post**: View all existing posts.
- **Update Post**: Modify the title and content of an existing post.
- **Delete Post**: Remove an existing post.

## Prerequisites
Ensure you have the following installed on your local machine:
- [Go](https://golang.org/dl/) (version 1.16+ is recommended)

## Setup & Installation
1. **Clone the Repository**
   ```sh
   git clone [https://github.com/yourusername/GoCRUDApp.git](https://github.com/DonnovanJiles70122/GO-CRUD.git)
   cd go-crud
   ```

2. Set Up the Database using GORM

    GORM is a fantastic ORM library for Golang, providing a simple way to interact with databases from our Go code.

    First, make sure to install GORM. You can do this by running the following command:
    ```sh
    go get -u gorm.io/gorm
    ```
    Ensure you call `InitDB` in your `main.go` or wherever your application starts running.
   
3. **Run the Application**
   ```sh
   CompileDaemon -command=./go-crud"
   ```
   The application should now be running at `http://localhost:8080`.

## API Endpoints
If you want to interact with the app via API, here are the available endpoints:

- **Get All Posts**: `GET /api/posts`
- **Get a Single Post**: `GET /api/posts/:id`
- **Create a Post**: `POST /api/posts`
- **Update a Post**: `PUT /api/posts/:id`
- **Delete a Post**: `DELETE /api/posts/:id`

Ensure to include the appropriate payload where necessary (for POST and PUT requests).

## Contributing
Feel free to fork the project, create a feature branch, and submit a pull request. For bugs, questions, and discussions please use the [Github Issues](https://github.com/yourusername/GoCRUDApp/issues).

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements
- Go Echo Framework
- Go GORM
- And everyone who has contributed to the above projects.

**Note**: Don't forget to replace placeholder text such as `yourusername` with actual values suitable for your project. You may also need to tailor the endpoints or technology stack according to your app's implementation.
