# Pulse Networking Tool 🌐

A lightweight, modular command-line utility for performing networking tasks like DNS lookups and port scanning.

## Features

- **DNS Lookups**: Query Google's public DNS (8.8.8.8) for various record types.
    - **A Records**: Retrieve IPv4 addresses (Default).
    - **CNAME Records**: Find canonical names.
    - **NS Records**: Identify authoritative name servers.
    - **MX Records**: Look up mail exchange records.
    - **TXT Records**: Read text-based metadata.
- **Port Scanning**: Check TCP port availability on any domain or IP address.

---

## Installation & Setup

Ensure you have [Go](https://go.dev/dl/) installed on your system.

```bash
# Clone the repository (if applicable)
# Navigate to the project directory
cd "path/to/ns"
```

---

## Usage

The tool uses a sub-command structure: `lookup` (alias `l`) and `scan` (alias `s`).

### 1. DNS Lookup
Perform DNS queries using the `lookup` command.

**Basic Lookup (Default: A Record)**
```bash
go run main.go lookup google.com
```

**Custom Record Types**
Use the `--type` flag to specify what you are looking for.
```bash
# Check Mail Servers (MX)
go run main.go lookup google.com --type MX

# Check Name Servers (NS) - Using Alias 'l'
go run main.go l google.com --type NS

# Check TXT Records
go run main.go l google.com --type TXT
```

### 2. Port Scanning
Check if a specific TCP port is reachable on a target host.

**Default Scan (Port 80)**
```bash
go run main.go scan google.com
```

**Custom Port**
Use the `--port` flag to define a specific target port.
```bash
# Check HTTPS (Port 443)
go run main.go scan google.com --port 443

# Check SSH (Port 22) - Using Alias 's'
go run main.go s 192.168.1.1 --port 22
```

---

## Documentation
For a deeper dive into the internal logic and file structure, check out the [documentation.md](./documentation.md) file.
