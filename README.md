# InfoCenter

A project that allows clients to achieve almost real-time communication between each other by sending messages, using concurrency to send or receive messages at the same time.

## How to run the application

- Clone the Git repository to your local machine using the command `git clone https://github.com/kipraskrasilnikas/info-center/`
- Open a terminal or command prompt and navigate to the directory where you cloned the repository.
- Run the command `go build -o infocenter.exe` to build the executable file.
- Run the command `./infocenter.exe` to launch the project. This will start the server on port 8080.
- Open a web browser and navigate to `http://localhost:8080/infocenter/`. This will show the default page for the InfoCenter project.

## How the project works

- The `main.go` file contains the entry point for the program. It starts an HTTP server and registers an HTTP request handler for the `/infocenter/` endpoint.
- When a client makes a `GET` request to the `/infocenter/<topic>` endpoint, the `InfocenterHandler` function in `main.go` calls the `GetHandler` function in `messageReceiver.go` to send any existing messages for the specified topic to the client.
- When a client makes a `POST` request to the `/infocenter/<topic>` endpoint, the `InfocenterHandler` function in `main.go` reads the message from the request body, assigns it a unique ID, and adds it to a slice of messages for the specified topic in the `messages` map.
- The `GetHandler` function in `messageReceiver.go` retrieves the slice of messages for the specified topic from the `messages` map and sends each message to the client using Server-Sent Events (SSE) format.
- If the client disconnects or the connection times out, the `GetHandler` function stops sending messages and returns.
- The `messages` map is protected by a `sync.RWMutex` to allow concurrent read access and exclusive write access.
