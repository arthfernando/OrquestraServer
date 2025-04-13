# PainelLembretesServer

### ServerSideEvent
- To open an Event Stream:
  ```bash
    curl http://localhost:3000/event-stream
  ```
- To send an Event from the server:
  ```bash
    curl -d '{"message":"Hello, Event Stream!"}' -H "Content-Type: application/json" -X POST http://localhost:3000/event-stream
  ```