# Contents of /project-docs/project-docs/docs/examples/javascript.md

# JavaScript Examples for In-Memory Database

This document provides examples of how to interact with the in-memory database using JavaScript. Below are various operations you can perform, along with code snippets for each operation.

## Connecting to the Database

To connect to the in-memory database, you can use the following code:

```javascript
import { createConnection } from './dushy-redis-lib.js';

const redis = createConnection('127.0.0.1', 6379);
```

## String Operations

### Setting a String Value

```javascript
await redis.set('greeting', 'Hello, World!');
```

### Getting a String Value

```javascript
const greeting = await redis.get('greeting');
console.log('GET greeting:', greeting);
```

## List Operations

### Adding Values to a List

```javascript
await redis.lpush('mylist', ['first', 'second']);
await redis.rpush('mylist', ['third', 'fourth']);
```

### Removing Values from a List

```javascript
const leftPop = await redis.lpop('mylist');
const rightPop = await redis.rpop('mylist');
console.log('LPOP:', leftPop);
console.log('RPOP:', rightPop);
```

## Set Operations

### Adding Members to a Set

```javascript
await redis.sadd('myset', ['apple', 'banana', 'orange']);
```

### Retrieving Members of a Set

```javascript
const members = await redis.smembers('myset');
console.log('SMEMBERS:', members);
```

## Hash Operations

### Setting a Field in a Hash

```javascript
await redis.hset('user:1', 'name', 'John Doe');
```

### Getting a Field from a Hash

```javascript
const name = await redis.hget('user:1', 'name');
console.log('HGET name:', name);
```

## Pub/Sub Operations

### Subscribing to a Channel

```javascript
await redis.subscribe('my_channel', (message) => {
    console.log('Received message:', message);
});
```

### Publishing a Message to a Channel

```javascript
await redis.publishString('my_channel', 'Hello, Subscribers!');
```

## Conclusion

These examples demonstrate basic operations you can perform with the in-memory database using JavaScript. For more advanced usage and features, refer to the API documentation.