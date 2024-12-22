# Commands Supported by the In-Memory Database

This document lists all the commands supported by the in-memory database, along with their syntax, parameters, and examples of usage.

## Command Overview

| Command      | Description                          |
|--------------|--------------------------------------|
| SET          | Sets a key-value pair                |
| GET          | Retrieves the value of a key         |
| LPUSH        | Pushes values to the left of a list  |
| RPUSH        | Pushes values to the right of a list |
| LPOP         | Pops a value from the left of a list  |
| RPOP         | Pops a value from the right of a list |
| SADD         | Adds members to a set                |
| SMEMBERS     | Gets all members of a set            |
| HSET         | Sets a field in a hash               |
| HGET         | Gets a field from a hash             |
| SUBSCRIBE    | Subscribes to a channel              |
| PUBLISH      | Publishes a message to a channel     |
| PUBLISH_JSON | Publishes JSON data to a channel     |
| PUBLISH_INT  | Publishes an integer to a channel     |
| PUBLISH_BIN  | Publishes binary data to a channel    |
| PUBLISH_ARRAY| Publishes an array to a channel      |


## Connecting from Terminal

- Use Netcat (nc) or Telnet to connect to the server:
  ```
  nc 127.0.0.1 6379
  ```
  or
  ```
  telnet 127.0.0.1 6379
  ```


## Command Syntax and Examples

### SET

**Syntax:** `SET <key> <value>`

**Example:**
```
SET mykey "Hello, World!"
```

### GET

**Syntax:** `GET <key>`

**Example:**
```
GET mykey
```

### LPUSH

**Syntax:** `LPUSH <key> <value1> [<value2> ...]`

**Example:**
```
LPUSH mylist "first" "second"
```

### RPUSH

**Syntax:** `RPUSH <key> <value1> [<value2> ...]`

**Example:**
```
RPUSH mylist "third" "fourth"
```

### LPOP

**Syntax:** `LPOP <key>`

**Example:**
```
LPOP mylist
```

### RPOP

**Syntax:** `RPOP <key>`

**Example:**
```
RPOP mylist
```

### SADD

**Syntax:** `SADD <key> <member1> [<member2> ...]`

**Example:**
```
SADD myset "apple" "banana"
```

### SMEMBERS

**Syntax:** `SMEMBERS <key>`

**Example:**
```
SMEMBERS myset
```

### HSET

**Syntax:** `HSET <key> <field> <value>`

**Example:**
```
HSET user:1 name "John Doe"
```

### HGET

**Syntax:** `HGET <key> <field>`

**Example:**
```
HGET user:1 name
```

### SUBSCRIBE

**Syntax:** `SUBSCRIBE <channel>`

**Example:**
```
SUBSCRIBE mychannel
```

### PUBLISH

**Syntax:** `PUBLISH <channel> <message>`

**Example:**
```
PUBLISH mychannel "Hello, Subscribers!"
```

### PUBLISH_JSON

**Syntax:** `PUBLISH_JSON <channel> <json_string>`

**Example:**
```
PUBLISH_JSON mychannel {"name": "John", "age": 30}
```

### PUBLISH_INT

**Syntax:** `PUBLISH_INT <channel> <number>`

**Example:**
```
PUBLISH_INT mychannel 42
```

### PUBLISH_BIN

**Syntax:** `PUBLISH_BIN <channel> <data>`

**Example:**
```
PUBLISH_BIN mychannel "Binary Data"
```

### PUBLISH_ARRAY

**Syntax:** `PUBLISH_ARRAY <channel> <json_array>`

**Example:**
```
PUBLISH_ARRAY mychannel [1, "two", {"three": 3}]
```

