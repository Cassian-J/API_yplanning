# yplanning API

Development of a dedicated API for managing collaborative schedules, enabling the identification of shared time slots between multiple users. The API integrates a secure authentication system based on JWT and interactive documentation generated with Swagger.

**Contributors:** Mathis and Cassian

---

## Getting Started

### Prerequisites

- Latest version of Go (starting by the 1.24 version) installed on your system
- Bruno or Postman for API testing (we recommend Bruno because you can test the links without connection)

a Swagger exist in the project if you want to test the project without Bruno or Postman

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Cassian-J/API_yplanning.git
cd yplanning
```

2. Create an `.env` file in the root directory with the following configuration:
```env
PORT=8080
JWT_SECRET=YourSecureSecretHere
REFRESH_SECRET=YourSecureRefreshSecretHere
```

💡 **Note:** You can choose any available port. We use 8080 by default.

⚠️ **Security Note:** Choose strong, unique secrets for production environments.

3. Start the server:
```bash
go run main.go
```

You should see output similar to this:
```
2026/01/09 08:31:47 Database migrated successfully
2026/01/09 08:31:47 Server running on http://localhost:8080
2026/01/09 08:31:47 Swagger UI available at http://localhost:8080/swagger/index.html
```

---

## Authentication

### First-Time Setup

Before accessing protected routes, you must register and obtain authentication tokens:

1. **Register a new user** by calling the `/auth/register` endpoint
2. You'll receive two tokens in the response:
   - **Access Token:** Used to authenticate API requests
   - **Refresh Token:** Used to generate new access tokens when they expire

### Using Protected Routes

All protected routes require a valid access token. Without proper authentication, you'll receive:
- `missing token` - if no token is provided
- `invalid token` - if the token is invalid or tampered with

When your access token expires, use the refresh token to generate a new one.

---

## Testing the API

### Option 1: Swagger UI (Recommended for Exploration)

1. Navigate to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
2. Click the **Authorize** button
3. Enter your access token in the **BearerAuth (apiKey)** value field
4. Click **Authorize** to access all protected endpoints

### Option 2: Bruno or Postman

Import the API collection and use your access token in the Authorization header:
```
Authorization: Bearer YOUR_ACCESS_TOKEN_HERE
```

---

## Project Structure

The API includes:
- JWT-based authentication system
- Database migration on startup
- Interactive Swagger documentation
- RESTful API endpoints

For detailed endpoint documentation, visit the Swagger UI once the server is running.

---

## Support

If you encounter any issues, please check that:
- Go is properly installed and up to date
- Your `.env` file is correctly configured
- The port 8080 is available on your system

---