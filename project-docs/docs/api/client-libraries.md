# Client Libraries for In-Memory Database

## Overview

This document provides details on the client libraries available for interacting with the in-memory database. These libraries simplify the process of connecting to the database and executing commands.

## Available Client Libraries

### JavaScript Client

- **Library Name**: `dushy-redis-lib.js`
- **Installation**:
  To install the JavaScript client library, use npm:

  ```
  npm install dushy-redis-lib
  ```

- **Usage Example**:
  ```javascript
  import { createConnection } from 'dushy-redis-lib';

  const redis = createConnection('127.0.0.1', 6379);

  async function example() {
      await redis.set('key', 'value');
      const value = await redis.get('key');
      console.log(value); // Outputs: value
  }

  example();
  ```

### Python Client

- **Library Name**: `dushy_redis_lib.py`
- **Installation**:
  To install the Python client library, you can use pip:

  ```
  pip install dushy-redis-lib
  ```

- **Usage Example**:
  ```python
  from dushy_redis_lib import DushyRedisClient

  client = DushyRedisClient.connect()
  client.set('key', 'value')
  value = client.get('key')
  print(value)  # Outputs: value
  ```

## Features

- **Asynchronous Operations**: Both clients support asynchronous operations, allowing for non-blocking interactions with the database.
- **Pub/Sub Support**: The clients provide built-in support for publish/subscribe messaging patterns.
- **Data Structures**: The libraries support various data structures such as strings, lists, sets, and hashes.

## Conclusion

These client libraries provide a convenient way to interact with the in-memory database, enabling developers to integrate database functionality into their applications easily. For more detailed usage and advanced features, refer to the respective library documentation.