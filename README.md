# Blockchain Simulator

A simple blockchain implementation in Go that simulates a real-world blockchain system with transaction processing, block creation, and validation.

## Overview

This project demonstrates the core concepts of blockchain technology through a working implementation that includes:

- **Blockchain Structure**: A chain of blocks with cryptographic hashing
- **Transaction Processing**: A pool-based system for batching transactions
- **Block Mining**: Automatic block creation when transaction batches are full
- **Integrity Validation**: Cryptographic verification of blockchain integrity
- **Concurrent Processing**: Multi-threaded transaction production and processing

## Features

### 🔗 Blockchain Core
- **Genesis Block**: Automatically created genesis block with index 0
- **Block Structure**: Each block contains index, transactions, previous hash, current hash, and timestamp
- **SHA256 Hashing**: Cryptographic hashing for block integrity
- **Chain Validation**: Verification of block sequence and hash integrity

### 💰 Transaction System
- **Transaction Pool**: Buffered channel-based transaction management
- **Batch Processing**: Configurable batch size for block creation
- **Timeout Mechanism**: Automatic block creation after timeout period
- **Random Transaction Generation**: Simulated user transactions with random amounts

### ⚡ Concurrent Architecture
- **Transaction Producer**: Multiple worker goroutines generating transactions
- **Transaction Consumer**: Dedicated goroutine for processing transaction batches
- **Thread-Safe Operations**: Mutex-protected blockchain operations
- **Graceful Shutdown**: Proper cleanup of goroutines and channels

## Project Structure

```
blockchain_simulator/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── README.md              # This file
└── internal/              # Core blockchain implementation
    ├── block.go           # Block structure and operations
    ├── blockchain.go      # Blockchain management
    ├── tx.go              # Transaction structure and operations
    ├── tx_pool.go         # Transaction pool and batching
    └── tx_producer.go     # Transaction generation
```

## Getting Started

### Prerequisites
- Go 1.24.0 or higher

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd blockchain_simulator
```

2. Run the simulator:
```bash
go run main.go
```

### Usage

The simulator runs automatically for 20 seconds and demonstrates:

1. **Initialization**: Creates a genesis block
2. **Transaction Generation**: 5 worker goroutines generate random transactions
3. **Block Creation**: Transactions are batched into blocks (batch size: 10)
4. **Chain Validation**: Verifies blockchain integrity
5. **Final Report**: Displays blockchain statistics

## Configuration

### Transaction Pool Settings
- **Batch Size**: 10 transactions per block (configurable in `main.go`)
- **Timeout**: 5 seconds for partial batch processing
- **Channel Buffer**: 100 transactions

### Transaction Producer Settings
- **Workers**: 5 concurrent transaction generators
- **Interval**: Random intervals between 400-800ms
- **Users**: 8 simulated users (Alice, Bob, Charlie, etc.)
- **Amount Range**: 0-100 random amounts

## Key Components

### Block Structure
```go
type Block struct {
    Index        int64
    Transactions []Transaction
    PrevHash     string
    Hash         string
    Timestamp    time.Time
}
```

### Transaction Structure
```go
type Transaction struct {
    ID        int64
    From      string
    To        string
    Amount    float64
    hash      string
    Timestamp time.Time
}
```

### Blockchain Operations
- `NewBlockchain()`: Creates new blockchain with genesis block
- `AddBlock(transactions)`: Adds new block with transaction batch
- `IsValid()`: Validates entire blockchain integrity

## Example Output

```
Blockchain initialized with genesis block
Transaction pool and producer started...
Press Ctrl+C to stop or wait 20 seconds for demo

Generated transaction: Alice -> Bob: 45.67
Added transaction to batch. Batch size: 1/10
Generated transaction: Charlie -> Diana: 23.45
Added transaction to batch. Batch size: 2/10
...
Created new block with 10 transactions. Blockchain length: 2
...

Final blockchain state:
Number of blocks: 5
Blockchain is valid: true
Block 0: 0 transactions, Hash: 0000000000000000...
Block 1: 10 transactions, Hash: a1b2c3d4e5f67890...
Block 2: 10 transactions, Hash: f1e2d3c4b5a67890...
...
```

## Technical Details

### Hashing Algorithm
- **Block Hash**: SHA256 of (index + prevHash + transactionsHash + timestamp)
- **Transaction Hash**: SHA256 of (id + from + to + amount + timestamp)

### Concurrency Model
- **Producer Pattern**: Multiple goroutines generating transactions
- **Consumer Pattern**: Single goroutine processing transaction batches
- **Channel Communication**: Buffered channels for transaction transfer
- **Mutex Protection**: Thread-safe blockchain operations

### Memory Management
- **Atomic Counters**: Thread-safe ID generation
- **Channel Buffering**: Prevents blocking on transaction submission
- **Garbage Collection**: Automatic cleanup of completed transactions

## Learning Objectives

This project demonstrates:
- **Blockchain Fundamentals**: Block structure, hashing, and chain validation
- **Concurrent Programming**: Goroutines, channels, and synchronization
- **Design Patterns**: Producer-consumer pattern implementation
- **Go Best Practices**: Error handling, memory management, and code organization

## Future Enhancements

Potential improvements for this simulator:
- **Proof of Work**: Mining difficulty and nonce calculation
- **Digital Signatures**: Transaction signing and verification
- **Network Simulation**: Multi-node blockchain network
- **Persistence**: Database storage for blockchain data
- **API Interface**: REST API for transaction submission
- **Web Interface**: Real-time blockchain visualization

## License

This project is for educational purposes. Feel free to use and modify as needed.

## Contributing

This is just a study case. Please feel free to submit pull requests or open issues for improvements and bug fixes.
