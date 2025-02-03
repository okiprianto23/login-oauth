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
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=oauth_db
DB_PORT=5432
JWT_SECRET=your_secret_key
```

---

## **4. Run the Application**
### **Start the Server**
```sh
go run main.go
```

### **Expected Output**
```
Database connected successfully
[GIN-debug] Listening and serving HTTP on :8080
```

---

## **5. API Endpoints**
### **Authentication**
| Method | Endpoint       | Description |
|--------|---------------|-------------|
| `POST` | `/auth/token` | Login and get JWT token |

### **Protected Routes**
| Method | Endpoint          | Description |
|--------|------------------|-------------|
| `GET`  | `/protected/data` | Access protected resource (requires token) |

---

## **6. Usage**
### **Login and Get Token**
Use the following **cURL** request to log in:

```sh
curl -X POST "http://localhost:8080/auth/token"      -H "Content-Type: application/json"      -d '{"username": "admin", "password": "password"}'
```
#### **Response**
```json
{
  "token": "your_generated_jwt_token"
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

