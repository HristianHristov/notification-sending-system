# Notification Sending System

A server-side streaming gRPC-based notification sending system with support for multiple channels (email, SMS, Slack). This system allows you to send notifications via various channels and receive status updates for each channel.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
- [Dockerization](#dockerization)
- [Kubernetes Deployment](#kubernetes-deployment)
- [Extensibility](#extensibility)
- [Error Handling](#error-handling)
- [TODO](#todo)

## Features

- Server-side streaming gRPC service for sending notifications.
- Supports email, SMS, and Slack as notification channels.
- Concurrency using goroutines and sync.WaitGroup for efficient notification sending.
- Retry mechanism for handling temporary failures during notification sending.

## Getting Started

### Prerequisites

- Go (1.19 or higher)
- [Twilio account](https://www.twilio.com/) for SMS notifications
- [Slack API token](https://api.slack.com/) for Slack notifications
- SMTP server details for email notifications

### Installation

1. Clone this repository:

```sh
git clone https://github.com/yourusername/notification-sending-system.git
cd notification-sending-system
```

2. Install the required dependencies:

```sh
go mod download
```

### Configuration

Create an .env file in the root directory of your project. You can use the .env.example file as a template. Update the values for your specific configuration (SMTP, Twilio, Slack, etc.).

## Usage

1. Run the server:

```sh
go run cmd/server/main.go
```

2. To test consuming the server, a client is provided. Start the client:

```sh
go run cmd/client/main.go
```

## Dockerization

To build and run the application using Docker, follow these steps:

1. Build the Docker image:

```sh
docker build -t notification-app -f build/Dockerfile .
```

2. Run the Docker container:

```sh
 docker run -d --name notification-container -p 5001:5001 my-notification-app
```

## Kubernetes Deployment

To deploy your application to Kubernetes, you can use the following command:

```sh
kubectl apply -f notification-deployment.yaml
```

This will create a deployment with 3 replicas of your notification application, using the Docker image notification-app:latest. The deployment ensures that the specified resources limits and requests are respected, helping with effective resource management.

## Extensibility

The system is designed to be easily extensible. Adding support for new notification channels involves creating a new channel implementation that adheres to the NotificationChannel interface. Then, updating the NotificationServer to handle the new channel type and adding it to the NotificationService.

## Error Handling

The system handles errors at multiple levels, including:
-Notification channel-specific errors
-Retrying failed notifications
-Streaming errors and successes back to the client
Errors are managed using channels, goroutines, and sync.WaitGroup to ensure that failures in one channel do not affect others and do not close the client stream.

## TODO

- [ ] Error Handling Refinement
- [ ] Database Integration
- [ ] Message Queues
- [ ] Notification Templates
