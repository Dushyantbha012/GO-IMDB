# Mini Redis Project Presentation

## **Introduction**

**Idea:**  
The idea behind this project is to create a lightweight, Redis-like in-memory data store with support for multiple data types and essential commands. The goal is to understand the underlying mechanics of In Memory Cache based Database and to provide a foundation for building high-performance, scalable applications.

**Project Genesis & Motivation**

My journey into building a In Memory Cached DataBase began during the development of a scalable web application that required a robust queuing system. After evaluating several options, I chose Redis Queue (RQ) for its:

- Simplicity and lightweight nature
- Excellent Python and JavaScript integration
- High performance characteristics
- Built-in support for job prioritization

Working with Redis opened my eyes to the elegance of in-memory data structures and their powerful applications in modern system design. Fascinated by its architecture and performance capabilities, I dove deep into Redis's internal implementation, studying:

- Memory management techniques
- Data structure implementations
- Client-server communication protocols
- Persistence strategies

This exploration inspired me to create my own Redis-like implementation to better understand these concepts hands-on.

**Purpose:**  
In Memory Cache based Databases are widely used for caching, message brokering, and real-time data processing. By implementing a simplified version, I aim to:
- Deepen our understanding of distributed systems and data stores.
- Create a customizable, lightweight alternative for specific use cases.
- Build a platform for further experimentation and learning.

---

## **Features To be Implemented in V1**

1. **Data Types Supported:**
   - **Strings:** Basic `SET` and `GET` commands.
   - **Lists:** Commands like `LPUSH`, `RPUSH`, `LPOP`, and `RPOP`.
   - **Sets:** Commands like `SADD` and `SMEMBERS`.
   - **Hashes:** Commands like `HSET` and `HGET`.

2. **Command Parsing:**
   - Efficient parsing of user commands using TCP sockets.
   - Basic error handling for incorrect commands and type mismatches.

3. **Concurrency:**
   - Safe access to shared data using mutex locks.
   - Multiple clients can connect to the server simultaneously.

4. **Extendable Design:**
   - Modular structure to easily add new data types and commands.

---

## **Technical Approach**

**1. Structure:**
- **Core Components:**
  - `store.go`: Core logic to store and manage data.
  - `server.go`: Handles TCP connections and command execution.
  - `types/`: Implements specific data types (e.g., lists, sets, hashes).

- **Flow:**
  - The server listens on a TCP port.
  - Commands are parsed and routed to corresponding functions.
  - Data is stored or retrieved from in-memory structures in `store.go`.

**2. Concurrency Management:**
- **Mutex Locks:** Ensures thread-safe operations on shared data.
- **Client Handling:** Each client is handled in a separate goroutine.

**3. Data Type Implementation:**
- **Strings:** Simple key-value pairs.
- **Lists:** Doubly linked lists to support efficient `LPUSH`, `RPUSH`, etc.
- **Sets:** Hash-based implementation for fast membership checks.
- **Hashes:** Hash maps for field-value pair storage.

---

## **Difficult Tasks**

1. **Concurrency Issues:**
   - Avoiding race conditions when multiple clients access the same key.
   - Implementing locks efficiently without introducing bottlenecks.

2. **Command Parsing:**
   - Handling malformed or incomplete commands gracefully.
   - Ensuring compatibility with a subset of the Redis command set.

3. **Memory Management:**
   - Balancing memory usage with performance.
   - Implementing efficient data structures for sorted sets and lists.

4. **Extensibility:**
   - Designing the system to allow for new data types and commands.

---

## **Future Improvements**

1. **Persistence (thinking to implement in V2):**
   - Add support for saving data to disk (similar to RDB or AOF in Redis).
   - Implement snapshotting for crash recovery.

2. **Expiration (thinking to implement in V2):**
   - Support for key expiration and eviction policies (e.g., LRU).

3. **Clustering (have to study more about it):**
   - Distribute keys across multiple instances for horizontal scaling.
   - Implement master-slave replication.

4. **Protocol Support :**
   - Implement Redis Serialization Protocol (RESP) for compatibility with Redis clients.
   - Basic commands till now support RESP

5. **Libraries for Integration (thinking to implement in V2):**
   - **JavaScript/Node.js:** Provide a library for interacting with the server in web applications.
   - **Golang:** Create a Go client for seamless server integration in backend services.
   - **Python:** Build a Python library for use in data processing and machine learning workflows.

6. **PubSub (thinking to implement in V2):**
     The Pub/Sub feature enables clients to subscribe to channels and receive messages published to those channels in real-time. This allows for efficient message broadcasting and is ideal for use cases like chat systems, real-time notifications, and live data streaming.

   **Key Features:**

   - Clients can subscribe to one or multiple channels.
   - Messages published to a channel are instantly pushed to all subscribers.
   - Supports pattern matching for subscribing to multiple channels with a single pattern.
   - Efficient handling of subscriber connections and message distribution.

7. **Advanced Data Types (have to study more about it) :**
   - Support for HyperLogLogs, streams, and geospatial indexes.

---

## **Conclusion**

This project demonstrates the core principles of building a high-performance, in-memory data store. While currently limited in scope, it serves as a foundation for learning and experimentation. With planned improvements like persistence, clustering, and client libraries, this project has the potential to evolve into a versatile and practical tool for developers.

