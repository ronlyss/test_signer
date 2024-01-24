Overview
The Test Signer Service is a Go-based application designed to manage the process of signing and verifying tests. It offers a simple REST API to sign user-submitted answers and to verify these signatures at a later stage. This service uses a local JSON file for storing test signatures, ensuring data persistence across restarts.

Features
Sign Answers: Accepts a set of answers along with user JWT and questions, then returns a unique test signature.
Verify Signature: Allows verification of a test signature against a user's JWT, returning the validation status and related data.
Getting Started
Prerequisites
Go (Golang) [version requirement, if any]
Basic knowledge of RESTful services and JSON file handling
Installation
Clone the repository:
git clone [repository URL]
Navigate to the project directory:
cd test-signer-service
Running the Service
Execute the following command in the project root directory:

go
go run main.go
This will start the service on localhost:8080.