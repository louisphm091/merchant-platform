# Merchant Platform (Golang)

A Go-based **Merchant Payment Platform** that simulates real-world fintech payment infrastructure.

This platform allows merchants to onboard, create payment orders, process payments via a mock checkout flow, receive webhook notifications, and manage refunds. It is designed to demonstrate key backend engineering concepts used in payment systems.

This project is also designed to integrate with **Gỗ Routex**, a bus-ticket booking system that uses this platform to process ticket payments.

---

# Architecture Overview

Merchant Platform acts as a **payment infrastructure layer** between merchants and customers.

Example integration flow:
Customer → Go Routex → Merchant Platform → Payment Processing
↓
Webhook
↓
Go Routex

### Flow

1. Merchant registers on the platform
2. Admin approves merchant account
3. Merchant generates API keys
4. Merchant creates a payment order
5. Customer completes payment via mock checkout
6. Platform updates payment status
7. Platform sends webhook notification to merchant system
8. Merchant system updates business logic (e.g., issue ticket)

---

# Features

### Merchant Management
- Merchant onboarding
- Merchant approval workflow
- Merchant profile management

### API Key Management
- Generate API keys
- Secure merchant API authentication
- HMAC signature verification

### Payment Order Management
- Create payment orders
- Payment lifecycle management
- Payment status tracking

### Mock Checkout
- Simulated checkout flow
- Payment success / failure simulation
- Realistic payment status transitions

### Webhook System
- Webhook notification delivery
- Signature verification
- Retry mechanism for failed deliveries

### Refund Management
- Full refund
- Partial refund
- Refund status tracking

### System Reliability
- Idempotency key support
- Audit logging
- Background worker for webhook delivery

---

# Tech Stack

| Component | Technology |
|----------|-----------|
Backend | Golang
Web Framework | Gin
Database | PostgreSQL
Cache / Queue | Redis
ORM | GORM
Authentication | JWT / API Key
Containerization | Docker
API Documentation | Swagger (planned)
---


