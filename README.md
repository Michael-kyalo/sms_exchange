# sms_exchange

The SMS Microservice is a lightweight and scalable service designed to facilitate sending SMS messages to customers using various SMS service providers. It abstracts the complexity of interacting with different SMS APIs and provides a simple and unified interface for sending SMS messages.

## Features

- **Asynchronous Processing**: Messages are processed asynchronously, ensuring responsiveness and fault tolerance.
- **Reliability**: Built-in error handling mechanisms and retries ensure reliable delivery of SMS messages.
- **Extensibility**: Easily extendable to support multiple SMS service providers and additional features.
- **Test Coverage**: Comprehensive unit tests and integration tests cover various scenarios to ensure robustness.
- **Modular Design**: Modular architecture with clear separation of concerns for easy maintenance and scalability.

## Installation

To install the SMS Microservice, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/Michael-Kyalo/sms_exchange.git

2. Change directory to the project root:

   ```bash
   cd sms_exchange

3. Initialize the Go module:

   ```bash
   go mod init github.com/Michael-Kyalo/sms_exchange

4. Install dependencies:

   ```bash
   go mod tidy

5. Build the project:

   ```bash
  go build -o sms_exchange ./cmd/sms/main.go
 

## Usage

To use the SMS Microservice, you need to integrate it into your existing system. Here's a basic example of how to send an SMS message using the microservice:



