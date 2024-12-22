# Communication Protocol for In-Memory Database

## Overview
This document describes the communication protocol used by the in-memory database. It outlines the message formats, command structures, and response handling mechanisms that clients must adhere to when interacting with the database.

## Message Format
All messages exchanged between the client and the server follow a specific format. Each message consists of a command followed by its parameters, terminated by a newline character (`\n`).

### Example Message
```
SET key "value"
```

## Commands
The following commands are supported by the in-memory database:

### SET
- **Description**: Sets a value for a specified key.
- **Syntax**: `SET <key> <value>`
- **Parameters**:
  - `key`: The key to set.
  - `value`: The value to associate with the key (can be a string, number, etc.).

### GET
- **Description**: Retrieves the value associated with a specified key.
- **Syntax**: `GET <key>`
- **Parameters**:
  - `key`: The key whose value is to be retrieved.

### LPUSH
- **Description**: Pushes one or more values onto the left end of a list.
- **Syntax**: `LPUSH <key> <value1> [<value2> ...]`
- **Parameters**:
  - `key`: The list key.
  - `value1`, `value2`, ...: Values to push onto the list.

### RPUSH
- **Description**: Pushes one or more values onto the right end of a list.
- **Syntax**: `RPUSH <key> <value1> [<value2> ...]`
- **Parameters**:
  - `key`: The list key.
  - `value1`, `value2`, ...: Values to push onto the list.

### SUBSCRIBE
- **Description**: Subscribes to a specified channel for receiving messages.
- **Syntax**: `SUBSCRIBE <channel>`
- **Parameters**:
  - `channel`: The channel to subscribe to.

### PUBLISH
- **Description**: Publishes a message to a specified channel.
- **Syntax**: `PUBLISH <channel> <message>`
- **Parameters**:
  - `channel`: The channel to publish to.
  - `message`: The message content.

## Response Handling
Responses from the server follow a specific format based on the command executed. The response may indicate success, failure, or return data.

### Example Responses
- For a successful `SET` command:
  ```
  OK
  ```
- For a successful `GET` command:
  ```
  value
  ```
- For an error:
  ```
  ERR: Invalid command
  ```

## Conclusion
Understanding the communication protocol is essential for effectively interacting with the in-memory database. Ensure that all messages adhere to the specified formats and handle responses appropriately to maintain a smooth operation.