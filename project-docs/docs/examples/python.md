# Python Examples for In-Memory Database

This document provides examples of how to interact with the in-memory database using Python. The examples cover basic operations such as setting and getting values, as well as more advanced features like publishing and subscribing to messages.

## Connecting to the Database

To connect to the in-memory database, you can use the following code snippet:

```python
from dushy_redis_lib import DushyRedisClient

# Create a connection to the database
client = DushyRedisClient.connect()
```

## String Operations

### Setting a Value

You can set a string value in the database using the `set` method:

```python
client.set("greeting", "Hello, World!")
```

### Getting a Value

To retrieve the value associated with a key, use the `get` method:

```python
value = client.get("greeting")
print(value)  # Output: Hello, World!
```

## List Operations

### Adding Values to a List

You can add values to a list using `lpush` and `rpush`:

```python
client.lpush("my_list", ["first", "second"])
client.rpush("my_list", ["third", "fourth"])
```

### Removing Values from a List

To remove values from the list, use `lpop` and `rpop`:

```python
left_value = client.lpop("my_list")
right_value = client.rpop("my_list")
print(left_value)  # Output: first
print(right_value)  # Output: fourth
```

## Set Operations

### Adding Members to a Set

You can add members to a set using the `sadd` method:

```python
client.sadd("my_set", ["apple", "banana", "cherry"])
```

### Retrieving Members of a Set

To get all members of a set, use the `smembers` method:

```python
members = client.smembers("my_set")
print(members)  # Output: ['apple', 'banana', 'cherry']
```

## Hash Operations

### Setting a Field in a Hash

You can set a field in a hash using the `hset` method:

```python
client.hset("user:1", "name", "John Doe")
```

### Getting a Field from a Hash

To retrieve a field from a hash, use the `hget` method:

```python
name = client.hget("user:1", "name")
print(name)  # Output: John Doe
```

## Pub/Sub Operations

### Subscribing to a Channel

You can subscribe to a channel to receive messages:

```python
def message_handler(message):
    print(f"Received message: {message}")

client.subscribe("my_channel", message_handler)
```

### Publishing a Message

To publish a message to a channel, use the `publish` method:

```python
client.publish("my_channel", "Hello, subscribers!")
```

## Conclusion

These examples demonstrate the basic usage of the in-memory database with Python. You can expand upon these examples to implement more complex functionalities as needed.