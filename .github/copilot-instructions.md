# Copilot Instructions for ServerlessDQL

## Project Overview
ServerlessDQL is a distributed system for Deep Q-Learning (DQL) using serverless and edge computing paradigms. The codebase is split into two main components:
- `CloudNode/`: Handles cloud-side orchestration, aggregation, and possibly model training.
- `EdgeNode/`: Contains edge logic, including Dockerfile and `Edge.go` for edge-side data collection, inference, or local training.

## Architecture & Data Flow
- Edge nodes (see `EdgeNode/Edge.go`) interact with the cloud node, likely sending experience data or receiving model updates.
- Communication between nodes may use REST, RPC, or message queues (details in respective directories).
- Each node is designed for containerization (see `EdgeNode/Dockerfile`).

## Developer Workflows
- **Build Edge Node:**
  - Use `docker build -t edgenode .` inside `EdgeNode/` to build the edge container.
- **Run Edge Node:**
  - Use `docker run edgenode` to start the edge node.
- **Debugging:**
  - For Go code, use standard Go tools (`go build`, `go run Edge.go`) inside `EdgeNode/`.
- **Cloud Node:**
  - Refer to `CloudNode/` for cloud-side build/run instructions (not present in README, inspect files for details).

## Conventions & Patterns
- Edge logic is implemented in Go (`EdgeNode/Edge.go`).
- Containerization is a first-class concern; always update Dockerfiles when changing dependencies.
- Inter-node communication should be abstracted for easy swapping of protocols.
- Keep edge and cloud logic decoupled; avoid direct imports across node boundaries.

## Integration Points
- External dependencies (e.g., Go modules, Python packages) should be declared in respective node directories.
- Any cloud-edge communication protocol changes must be reflected in both `CloudNode/` and `EdgeNode/`.

## Key Files
- `EdgeNode/Edge.go`: Main edge node logic.
- `EdgeNode/Dockerfile`: Container spec for edge node.
- `CloudNode/`: Cloud orchestration and aggregation logic (details depend on files present).

## Example: Building and Running Edge Node
```bash
cd EdgeNode
# Build Docker image
docker build -t edgenode .
# Run container
docker run edgenode
```

---
For questions about unclear workflows or missing conventions, ask for clarification or inspect the relevant directory for scripts/configs.
