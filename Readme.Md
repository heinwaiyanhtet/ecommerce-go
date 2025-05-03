# 🧪 User CRUD API Testing (cURL)

This guide provides sample `curl` commands for testing the **Create, Read, Update, and Delete** operations on the `/users` endpoint of your Go backend.

---

## ✅ Create User (POST)

```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "email": "john@example.com"}'


curl -X GET http://localhost:8080/users/1


curl -X PUT http://localhost:8080/users/1 \
-H "Content-Type: application/json" \
-d '{"name": "Johnny D", "email": "johnnyd@example.com"}'


curl -X DELETE http://localhost:8080/users/1
