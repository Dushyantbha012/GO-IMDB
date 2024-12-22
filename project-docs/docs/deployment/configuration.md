# Configuration Options for Deployment

This document outlines the configuration options available for deploying the in-memory database. Proper configuration is essential for optimal performance and functionality.

## Environment Variables

The following environment variables can be set to customize the behavior of the in-memory database:

- **DB_HOST**: Specifies the hostname of the database server. Default is `127.0.0.1`.
- **DB_PORT**: Specifies the port on which the database server listens. Default is `6379`.

## Configuration Files

The in-memory database can also be configured using a configuration file. The default configuration file is named `config.json` and should be placed in the root directory of the project. Below is an example of the configuration file structure:

```json
{
  "host": "127.0.0.1",
  "port": 6379,
}
```

### Configuration File Options

- **host**: The hostname of the database server.
- **port**: The port number for the database server.

## Applying Configuration

To apply the configuration, ensure that the environment variables are set or the configuration file is correctly formatted and placed in the appropriate directory. The in-memory database will automatically read these settings upon startup.

## Conclusion

Proper configuration is crucial for the effective deployment of the in-memory database. Adjust the environment variables and configuration file as needed to suit your deployment environment.