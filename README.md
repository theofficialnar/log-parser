# log-parser
CLI log parser made with Golang

## Usage

1. Build package using the go cli command.

    ```bash
      go build log-parser.go
    ```

2. Run command with the following arguments:

    - `-path` - Path to file
    - `-type` - Log type to filter for

    
    ```bash
      log-parser -type ERROR -path "./path/to/file"
    ```