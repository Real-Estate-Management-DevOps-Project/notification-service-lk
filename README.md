# Notification Service

A dedicated microservice for handling system-wide notifications in the Real Estate Management platform. Supports Email, SMS, and Push notifications with asynchronous processing and history tracking.

## Features

✅ **Multi-Channel Support**

- **Email**: Transactional emails, alerts, and detailed reports.
- **SMS**: Time-sensitive alerts and verifications.
- **Push**: Mobile app notifications for real-time engagement.

✅ **Robust Architecture**

- **Asynchronous Processing**: Non-blocking sending mechanism for high throughput.
- **Status Tracking**: Track notification status (PENDING, SENT, FAILED).
- **Error Handling**: Detailed error logs for failed deliveries.
- **Retry Mechanism**: (Supported in future updates) Automatic retries for transient failures.

✅ **History & Analytics**

- **Audit Logs**: Full history of sent notifications.
- **User History**: Retrieve notification history for specific users.
- **Metadata**: Store subject, content, and type for complete visibility.

## API Endpoints

### Sending

- `POST /api/v1/notifications/send` - Queue a new notification
  - Supports `EMAIL`, `SMS`, `PUSH` types.
  - Input: `recipient`, `type`, `subject`, `content`.

### History & Retrieval

- `GET /api/v1/notifications/history` - Get global notification history (Paginated).
- `GET /api/v1/notifications/:id` - Get details of a specific notification.
- `GET /api/v1/notifications/user/:userId` - (Planned) Get notifications for a specific user.

### Health

- `GET /health` - Service health check.

## Setup

### Prerequisites

- Go 1.24+
- PostgreSQL 16+

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Real-Estate-Management-DevOps-Project/notification-service-lk.git
   cd notification-service
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Create the database:

   ```sql
   CREATE DATABASE real_estate_notifications;
   ```

4. Configure Environment:
   Create a `.env` file (or use defaults in `config/config.go`):
   ```env
   PORT=3004
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=real_estate_notifications
   DB_PORT=5432
   ```

### Running the Service

```bash
go run main.go
# OR use the Makefile
make run
```

### Docker

```bash
make docker
docker run -p 3004:3004 notification-service:local
```

## Example Requests

### Send Notification

```json
{
  "recipient": "user@example.com",
  "type": "EMAIL",
  "subject": "Welcome to Real Estate Platform",
  "content": "Thank you for signing up! We are excited to have you."
}
```

### Response

```json
{
  "message": "Notification queued for sending",
  "id": "a1b2c3d4-e5f6-7890-1234-567890abcdef"
}
```

## License

MIT
