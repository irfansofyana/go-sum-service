# Simple Sum Service using Go

## Prerequisite

1. Go
2. Docker (optional)

## How to Run

- Without Docker

     ```
    $ go run app/main.go
    ```

- With Docker
    - Build docker image
        ```
        $ docker build -t <desired_tag> .
        ```
    - Run the image
        ```
        $ docker run --rm -it -p <host_port>:8080 <desired_target>
        ```
