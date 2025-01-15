# TCP Server Usage

This document provides examples of how to interact with the GO-IMDB TCP server.

## Starting the Server

To start the server, run the following command:

```sh
go run main.go
```

The server will start listening on port `6379`.

## Connecting to the Server

You can connect to the server using any TCP client. For example, you can use `telnet`:

```sh
telnet localhost 6379
```

## Commands

### SET

Set a key to hold the string value.

```sh
SET key value
```

Example:

```sh
SET mykey "Hello, World!"
```

### GET

Get the value of a key.

```sh
GET key
```

Example:

```sh
GET mykey
```

### LPUSH

Prepend one or multiple values to a list.

```sh
LPUSH key value [value ...]
```

Example:

```sh
LPUSH mylist "first" "second"
```

### RPUSH

Append one or multiple values to a list.

```sh
RPUSH key value [value ...]
```

Example:

```sh
RPUSH mylist "third"
```

### LPOP

Remove and get the first element in a list.

```sh
LPOP key
```

Example:

```sh
LPOP mylist
```

### RPOP

Remove and get the last element in a list.

```sh
RPOP key
```

Example:

```sh
RPOP mylist
```

### SADD

Add one or more members to a set.

```sh
SADD key member [member ...]
```

Example:

```sh
SADD myset "member1" "member2"
```

### SMEMBERS

Get all the members in a set.

```sh
SMEMBERS key
```

Example:

```sh
SMEMBERS myset
```

### HSET

Set the string value of a hash field.

```sh
HSET key field value
```

Example:

```sh
HSET myhash field1 "value1"
```

### HGET

Get the value of a hash field.

```sh
HGET key field
```

Example:

```sh
HGET myhash field1
```

### SUBSCRIBE

Subscribe to a channel.

```sh
SUBSCRIBE channel
```

Example:

```sh
SUBSCRIBE mychannel
```

### PUBLISH

Publish a message to a channel.

```sh
PUBLISH channel message
```

Example:

```sh
PUBLISH mychannel "Hello, Subscribers!"
```

### PUBLISH_JSON

Publish a JSON message to a channel.

```sh
PUBLISH_JSON channel json_string
```

Example:

```sh
PUBLISH_JSON mychannel '{"key": "value"}'
```

### PUBLISH_INT

Publish an integer message to a channel.

```sh
PUBLISH_INT channel number
```

Example:

```sh
PUBLISH_INT mychannel 123
```

### PUBLISH_BIN

Publish binary data to a channel.

```sh
PUBLISH_BIN channel data
```

Example:

```sh
PUBLISH_BIN mychannel "binarydata"
```

### PUBLISH_ARRAY

Publish an array message to a channel.

```sh
PUBLISH_ARRAY channel json_array
```

Example:

```sh
PUBLISH_ARRAY mychannel '[1, 2, 3]'
```
