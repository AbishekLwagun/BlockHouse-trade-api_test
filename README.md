# CI/CD Pipeline with Docker & EC2 Deployment

## Overview
This project demonstrates how I, **Abishek Lwagun**, set up a CI/CD pipeline using GitHub Actions to deploy a backend service to an AWS EC2 instance. The backend service exposes REST APIs to manage trade orders. I containerized the application using Docker, deployed it to EC2, and automated the process using a GitHub Actions workflow.

### Technologies Used:
- **Backend**: Golang (Gin) or Python (FastAPI)
- **CI/CD**: GitHub Actions
- **Containerization**: Docker
- **Deployment**: AWS EC2 (Ubuntu)
- **Database**: SQLite (for simplicity)

```
http://3.142.147.59:8080
```

POST Request to Create an Order:
```
curl -X POST http://3.142.147.59:8080/orders \
-H "Content-Type: application/json" \
-d '{"symbol": "AAPL", "price": 150, "quantity": 10, "order_type": "buy"}'

```
GET Request to Fetch Orders:
```
curl http://3.142.147.59:8080/orders
The POST request creates a new order, and the GET request fetches all the orders.
```

## GitHub Secrets Configuration
For the GitHub Actions workflow to SSH into EC2, I added the following secrets in my GitHub repository:
```
EC2_SSH_PRIVATE_KEY: My EC2 private SSH key.
EC2_PUBLIC_IP: The public IP of my EC2 instance.
DOCKER_USERNAME: My Docker Hub username.
DOCKER_PASSWORD: My Docker Hub password or token.
```
## Conclusion
I successfully created a Dockerized backend service, set up a CI/CD pipeline with GitHub Actions, and deployed the application to AWS EC2. The pipeline automatically builds, tests, and deploys the app when changes are made to the main branch. I also tested the POST and GET requests to ensure the orders are being processed correctly.
