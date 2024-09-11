# Golang Fullstack - Technical Test

Welcome to the Golang Fullstack Technical Test repository.

## Project Structure

The project is organized as follows:

- **backend**: Contains the Go server and related code.
- **frontend**: Contains the React code for the client-side application.

## Running the Project

### Case 1

1. Navigate to the backend directory:
    ```sh
    cd case1
    ```
2. Install the dependencies:
    ```sh
    go mod tidy
    ```
3. Run the server:
    ```sh
    go run cmd/app/main.go
    ```

### Case 2

1. Navigate to the frontend directory:
    ```sh
    cd case2
    ```
2. Install the dependencies:
    ```sh
    go mod tidy
    ```
3. Start the development server:
    ```sh
    go run cmd/app/main.go
    ```

### Case 3:

1. **Setting Up Backend:**
    - Navigate to the backend directory:
        ```sh
        cd case3/backend
        ```
    - Install the dependencies:
        ```sh
        go mod tidy
        ```
    - Run database migrations:
        ```sh
        ./database/database.sql
        ```
    - Setup env file
    - Run the server:
        ```sh
        go run main.go
        ```

2. **Setting Up Frontend:**
    - Open a new terminal window or tab.
    - Navigate to the frontend directory:
        ```sh
        cd case3/frontend
        ```
    - Install the dependencies:
        ```sh
        npm install
        ```
    - Setup env file
    - Start the development server:
        ```sh
        npm dev
        ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.