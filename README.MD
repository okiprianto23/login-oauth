# **OAuth 2.0 Server with Golang, PostgreSQL, and JWT**

This is a simple **OAuth 2.0 authentication server** built using **Golang**, **Gin**, **PostgreSQL**, and **JWT**. This server allows users to **log in** using their credentials and obtain an **access token** to access protected resources.

## **Features**
- ✅ Secure login with JWT
- ✅ Protected API routes with middleware
- ✅ PostgreSQL database integration
- ✅ Structured project architecture

---

## **1. Prerequisites**
Before running this project, make sure you have the following installed:
- [Golang (v1.20+)](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/downloads)

---

## **2. Installation**
### **Clone the repository**
```sh
git clone https://github.com/yourusername/oauth-server.git
cd oauth-server
```

### **Install dependencies**
```sh
go mod tidy
```

---

## **3. Database Setup**
### **Create a PostgreSQL Database**
```sql
CREATE DATABASE oauth_db;
```

### **Set up `.env` file**
Create a `.env` file in the root directory and update the database credentials:
```
DB_HOST=localhost
DB_USER=oauth
DB_PASSWORD=P@ssw0rd
DB_NAME=oauth
DB_PORT=5442
JWT_SECRET=f0972a20ea7f1f23f863e99314c70c3d0525c17000ea56cc8b6044f555262444
JWT_REFRESH_SECRET=5a5b1f1174e0d629cba2daf895b03146edd4134092e89824cf87ce0162714b12
```

---

## **4. Run the Application**
### **Docker Compose Up**
````sh
docker-compose up -d
````

### **Expected Output**
```
[+] Running 2/2
 ✔ Network login-oauth_default  Created                                                                                                                                                                                                                                                                        0.1s 
 ✔ Container postgres_oauth     Started 
```

### **Start the Server**
```sh
go run main.go
```

### **Expected Output**
```
Database connected successfully
Database migration completed successfully!
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /auth/register            --> github.com/okiprianto23/login-oauth/controllers.Register (3 handlers)
[GIN-debug] POST   /auth/token               --> github.com/okiprianto23/login-oauth/controllers.Login (3 handlers)
[GIN-debug] POST   /auth/refresh             --> github.com/okiprianto23/login-oauth/controllers.RefreshToken (3 handlers)
[GIN-debug] GET    /protected/data           --> github.com/okiprianto23/login-oauth/routes.SetupRouter.func1 (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :9899
```

---

## **5. API Endpoints**
### **Authentication**
| Method | Endpoint         | Description             |
|--------|------------------|-------------------------|
| `POST` | `/auth/register` | Register user login     |
| `POST` | `/auth/login`    | Login and get JWT token |
| `POST` | `/auth/refresh`  | Refresh JWT token       |

### **Protected Routes**
| Method | Endpoint          | Description |
|--------|------------------|-------------|
| `GET`  | `/protected/data` | Access protected resource (requires token) |

---

## **6. Usage**
### **Login and Get Token**
Use the following **cURL** request to log in:

```sh
curl -X POST "http://localhost:8080/auth/login"      -H "Content-Type: application/json"      -d '{"username": "admin", "password": "P@ssw0rd"}'
```
#### **Response**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg1NzEzNjYsInVzZXJuYW1lIjoiYWRtaW4ifQ.KIo7gYWWu9dcTE6XFdVqA9f3-LcA4Hi_TQQUaZGLKfs"
}
```

### **Access Protected Endpoint**
Use the obtained token to access a protected route:

```sh
curl -X GET "http://localhost:8080/protected/data"      -H "Authorization: Bearer your_generated_jwt_token"
```

#### **Response**
```json
{
  "message": "Protected Data"
}
```

---

## **7. Project Structure**
```
/oauth-server
│── /config          # Configuration files (DB, ENV)
│── /controllers     # API logic (Auth, User)
│── /models         # Database models
│── /routes         # API routes
│── /middlewares    # JWT middleware
│── /utils          # Helper functions (JWT, Hashing)
│── main.go         # Entry point
│── go.mod          # Dependencies
│── .env            # Environment variables
```

---

## **8. Future Enhancements**
- ✅ Implement **User Registration**
- ✅ Add **Password Hashing** for security
- ✅ Implement **Refresh Tokens**
- ✅ Add **Role-Based Authorization**

---

## **9. License**
This project is licensed under the **MIT License**.

