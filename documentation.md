# Network Lookup Tool (ns) Documentation 🌐

Welcome to the documentation for the **ns** CLI tool. This project is a modular networking utility built in Go, designed to perform DNS lookups and port scanning efficiently.

---

## 📁 Directory Structure & File Descriptions

The project follows a clean, modular structure to separate CLI configuration from the core networking logic.

### Root Directory
- **`main.go`**: The application heart. It initializes the CLI framework, defines the user interface (commands and flags), and routes execution to the internal engine.
- **`go.mod` / `go.sum`**: Dependency management files.
- **`documentation.md`**: You are here! Complete guide to the codebase.

### `cmd/` Sub-directory
This package contains the specific "engines" for our networking tasks:
- **`lookup.go`**: Handles all DNS-related operations, including talking to upstream resolvers.
- **`port.go`**: Contains the logic for testing TCP port connectivity.

---

## 🛠️ Detailed File Logic

### 1. `main.go`
This file uses the `urfave/cli/v3` library to create a professional terminal interface.
- **Logic**: It defines a `rootCmd` with a list of sub-commands.
- **Commands**: 
    - `lookup`: Linked to the `cmd.Lookup` function. It includes a `--type` flag (defaulting to "A").
    - `scan`: Linked to the `cmd.ScanPort` function. It includes a `--port` flag (defaulting to "80").
- **Execution**: It takes `os.Args`, parses them, and executes the mapped `Action` function.

### 2. `cmd/lookup.go`
Responsible for querying DNS records from a specific name server.
- **`Lookup` Handler**: Acts as the bridge between the CLI and the DNS engine. It reads the domain from the arguments and the record type from the flags.
- **`QueryDNS` Engine**: Instead of using the default OS resolver, this creates a custom `net.Resolver`. It forces the tool to use UDP to talk to `8.8.8.8:53` (Google DNS), ensuring consistent results across different environments.
- **Record Specificity**: It uses a `switch` statement to call specialized Go functions like `LookupIPAddr`, `LookupMX`, or `LookupTXT` depending on what the user asked for.

### 3. `cmd/port.go`
A simple but effective TCP port checker.
- **`ScanPort` Handler**: Parses the target domain and the target port.
- **`scanPort` Engine**: The core logic uses `net.DialTimeout`. 
    - **How it works**: It tries to establish a TCP "Three-Way Handshake". 
    - If the handshake succeeds within 10 seconds, the port is marked as **"open"**.
    - It specifically catches "timeout" errors to provide a clearer message than a generic "connection failed".

---

## 🔄 System Architecture & Flow

The tool follows a **Command Pattern** flow:

1.  **Input**: User types `go run main.go lookup google.com --type MX`.
2.  **Dispatch**: `main.go` recognizes the `lookup` command. It passes the context (arguments and flags) to `cmd.Lookup`.
3.  **Core Logic**: 
    - `cmd.Lookup` determines the user wants an `MX` record for `google.com`.
    - It calls `QueryDNS`, which opens a connection to `8.8.8.8`.
    - `QueryDNS` retrieves the mail server data.
4.  **Output**: The result is sent back up to the handler, which prints it gracefully to the terminal: `the ip for google.com is: ...`.

---

## 🚀 Future Scalability
The project is structured so that adding a new feature (like `ping` or `whois`) only requires:
1.  Creating a new file in `cmd/`.
2.  Adding a new entry in the `Commands` list in `main.go`.
