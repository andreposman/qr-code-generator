# QRCode Generator

## Steps to run the project

1. Clone the repo:

    `git clone https://github.com/andreposman/qr-code-generator.git`


2. Go to the directory:

    `cd qr-code-generator`

3. Install dependencies with: 

    `go mod download`


4. Run the project:

    `go run ./cmd/main.go`

5. The server will start running in the 9000 port, you can make an POST call like this:
    ````
    {
        "data": "https://github.com/andreposman",
        "file_name": "qrcode_name"
    }
    ````
 