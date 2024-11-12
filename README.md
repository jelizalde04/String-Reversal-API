# String Reversal API

This API takes a string and returns the reversed version of the string.

## Setup Instructions

1. Clone the repository:
    ```bash
    git clone https://github.com/jelizalde04/tring-reversal-api.git
    cd string-reversal-api
    ```

2. Build and run the Docker container locally:
    ```bash
    docker build -t string-reversal-api .
    docker run -p 8080:8080 string-reversal-api
    ```

3. Access the API at `http://localhost:8080`. You can send a GET request to this URL to reverse a string.

## Example url

http://localhost:8080/reverse?string=Hello%20World
