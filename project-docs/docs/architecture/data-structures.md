# Data Structures Used in the In-Memory Database

## Overview

This document provides a detailed description of the data structures utilized within the in-memory database project. These structures are designed to facilitate efficient data storage, retrieval, and manipulation, ensuring optimal performance and scalability.

## Data Structures

### 1. Hash

- **Description**: A hash is a collection of field-value pairs, allowing for efficient retrieval of values based on unique keys.
- **Implementation**: The hash is implemented using a map data structure, which provides average O(1) time complexity for set and get operations.
- **Usage**: Suitable for storing objects with multiple attributes, such as user profiles or configuration settings.

### 2. Set

- **Description**: A set is an unordered collection of unique strings, ensuring that no duplicate values are stored.
- **Implementation**: The set is implemented using a map where the keys represent the unique elements, allowing for O(1) time complexity for add and check operations.
- **Usage**: Ideal for scenarios where uniqueness is required, such as managing user permissions or tags.

### 3. List

- **Description**: A list represents a sequence of elements, supporting operations that allow for adding and removing items from both ends (FIFO).
- **Implementation**: The list is implemented using a slice, providing dynamic resizing and efficient access to elements.
- **Usage**: Useful for maintaining ordered collections, such as task queues or message buffers.

### 4. Pub/Sub

- **Description**: The publish/subscribe model allows for asynchronous communication between different components of the system.
- **Implementation**: This structure manages channels and subscriptions, enabling messages to be sent to multiple subscribers efficiently.
- **Usage**: Suitable for real-time notifications and event-driven architectures.

## Conclusion

The data structures outlined in this document are fundamental to the functionality of the in-memory database. Their design prioritizes performance and ease of use, making the database both powerful and flexible for various applications.