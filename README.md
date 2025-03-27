# 📡 Broadcast Server

A simple WebSocket-based broadcast server built with Go and `gorilla/websocket`.

## 🚀 Features
- WebSocket communication using `gorilla/websocket`
- CLI commands powered by `spf13/cobra`
- Client and server interaction

## 📦 Installation
```sh
# Clone the repository
git clone https://github.com/Abhishek2010dev/Broadcast-Server
cd Broadcast-Server

# Install dependencies
go mod tidy
```

## 🛠 Usage
### Start the Server
```sh
go run main.go start
```
Server will start at: [http://localhost:3000](http://localhost:3000)

### Connect as a Client
```sh
go run main.go connect --username alice
```

## 📜 License
This project is licensed under the MIT License.

## 🔗 Reference
[Project Roadmap](https://roadmap.sh/projects/broadcast-server)

