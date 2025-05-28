# Admission Portal Backend

---

## 🚀 Project Overview
A Go-based backend system for managing student admissions to courses. This system allows:
- Student registration and management
- Course creation and management (admin only)
- Admission/enrollment of students to courses
- Admin roles and authentication (JWT)
- Status updates and more

---

## ✨ Features
- Student registration and authentication
- Course management (CRUD)
- Admission/enrollment system
- Admin roles
- JWT authentication
- Status updates
- RESTful API endpoints

---

## 🛠️ Tech Stack
- Go 1.21
- Gin Web Framework
- MongoDB (Atlas)
- Docker & Docker Compose

---

## 📁 Project Structure
```
admission-portal-backend/
├── cmd/
│   └── main.go
├── internal/
│   ├── controllers/
│   ├── models/
│   ├── routes/
│   ├── middlewares/
│   └── config/
├── .env
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## ⚡ How to Use This Project

### 1. Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Git](https://git-scm.com/) (optional, for cloning the repo)

### 2. Setup & Run (with Docker)
1. **Clone or Download the Project**
   ```bash
   git clone <your-repo-url>
   cd admission-portal-backend
   ```
2. **Check/Edit the `.env` File**
   - Ensure `.env` exists in the root folder.
   - Set your MongoDB connection string and secrets as needed.
3. **Start the Backend**
   ```bash
   docker compose up --build
   ```
   - Wait for: `Connected to MongoDB Atlas!`
   - The backend runs at: [http://localhost:8080](http://localhost:8080)
4. **Stop the Backend**
   ```bash
   docker compose down
   ```

---

### 3. Using the API with Postman

#### Import the Collection and Environment
- [Postman Collection][collection-link]
- [Postman Environment][environment-link]

[collection-link]: https://anushkagupta-5713319.postman.co/workspace/Anushka-Gupta's-Workspace~05e96161-58f3-40d3-8c26-635c074cb518/collection/44856730-9df2a736-63e5-49ec-b900-d3953ec17329?action=share&creator=44856730&active-environment=44856730-54475ace-eaf0-4e8c-bf14-3cfcd2bda4cc
[environment-link]: https://anushkagupta-5713319.postman.co/workspace/Anushka-Gupta's-Workspace~05e96161-58f3-40d3-8c26-635c074cb518/environment/44856730-54475ace-eaf0-4e8c-bf14-3cfcd2bda4cc?action=share&creator=44856730&active-environment=44856730-54475ace-eaf0-4e8c-bf14-3cfcd2bda4cc

#### Steps
1. Open [Postman](https://www.postman.com/downloads/) (Desktop app recommended).
2. Import the collection and environment using the links above.
3. Select the environment in the top right of Postman.
4. Use the pre-configured requests to test the API (start with Create Admin or Create Student, then Login, etc.).
5. Make sure the Desktop Agent is running if you use the web version.

#### If You Cannot Fork or Import: How to Manually Recreate the Collection and Environment

If you are unable to fork or import the Postman collection and environment directly, follow these steps to manually set up everything exactly as shown:

1. **Create a New Collection**
   - In Postman, go to the Collections sidebar.
   - Click **New** > **Collection**.
   - Name it (e.g., "Admission portal") and click **Create**.

2. **Add Requests to the Collection**
   - Click your new collection.
   - Click **Add Request** for each API endpoint (see the API list in this README).
   - For each request:
     - Set the method (GET, POST, etc.) and URL (e.g., `{{base_url}}/api/students/signup`).
     - For POST/PUT, go to the **Body** tab, select **raw** and **JSON**, and paste the example body from the README.
     - Add required headers (e.g., `Content-Type: application/json`, `Authorization: Bearer {{student_token}}`).
     - Name and save the request.
   - Repeat for all endpoints you want to add.

3. **Create a New Environment**
   - Go to the Environments sidebar.
   - Click **New Environment**.
   - Add variables:
     - `base_url` → `http://localhost:8080`
     - `student_token` → (leave blank for now)
     - `admin_token` → (leave blank for now)
   - Save and name your environment (e.g., "Admission Portal Env").

4. **Use Variables in Requests**
   - In request URLs, use `{{base_url}}/api/...`.
   - In headers, use `Bearer {{student_token}}` or `Bearer {{admin_token}}` as needed.

5. **Select the Environment**
   - In the top right of Postman, select your new environment.

6. **Test the API**
   - Use the requests to interact with your backend.
   - After logging in, copy the returned token into the appropriate environment variable.

> This process ensures you have a working Postman setup identical to the shared one, even if you cannot fork or import directly.

> You can also view my shared Postman collection using the provided public link. If you want your setup to match exactly, simply copy and paste the request URLs, bodies, and headers from my collection into your own requests.

---

### 4. Environment Variables
- Edit the `.env` file in the project root to change `MONGODB_URI`, `PORT`, `JWT_SECRET`, etc.
- Restart Docker after making changes.

---

### 5. Adding & Viewing Data
- Use "Create Course" and "Apply for Admission" endpoints in Postman to add data.
- Use "Get" endpoints to view students, courses, and admissions.

---

### 6. Troubleshooting
- **ECONNREFUSED**: Make sure Docker is running and the backend is up.
- **Invalid credentials**: Double-check your email and password.
- **Unauthorized**: Ensure you are sending the correct JWT token in the Authorization header.

---

## 🛠️ Troubleshooting
If you encounter issues while running or using this project, here are some common problems and solutions:

1. **Cannot Connect to MongoDB / Database Errors**
   - **Cause:** The `.env` file is missing or the `MONGODB_URI` is incorrect.
   - **Solution:**
     - Make sure you have a `.env` file in your project root.
     - Update the `MONGODB_URI` with your own MongoDB Atlas connection string.
     - Restart Docker after making changes.

2. **ECONNREFUSED or Cannot Connect to Backend**
   - **Cause:** The backend server is not running, or Docker is not started.
   - **Solution:**
     - Ensure Docker Desktop is running.
     - Run `docker compose up --build` in your project directory.
     - Wait for the message: `Connected to MongoDB Atlas!`
     - The backend should be available at [http://localhost:8080](http://localhost:8080).

3. **Invalid Credentials / Unauthorized**
   - **Cause:** Wrong email or password, or missing JWT token in requests.
   - **Solution:**
     - Double-check your login credentials.
     - Make sure you copy the JWT token from the login response and set it in your Postman environment.
     - Add the token to the `Authorization` header as `Bearer <token>`.

4. **Cannot Import or Fork Postman Collection**
   - **Cause:** Postman permissions or public link limitations.
   - **Solution:**
     - If you cannot import or fork, manually recreate the collection and environment in Postman by copying request details from the README or the shared collection link.
     - See the "How to Manually Recreate the Collection and Environment" section above.

5. **.env File Not Included**
   - **Cause:** For security, `.env` is excluded from the repository.
   - **Solution:**
     - Create your own `.env` file in the project root.
     - Use the example in the README to set up your environment variables.
     - **Note:** You can view the sample `.env` file provided in the ZIP folder.
     - This file contains example environment variable names and formats required to run the project.

6. **Other Issues**
   - **Check the logs:**
     - Use `docker compose logs` to view backend logs for errors.
   - **Check your internet connection** (for MongoDB Atlas).
   - **Consult the README** for setup and usage instructions.

If you encounter an issue not listed here, please open an issue on GitHub or contact the maintainer.

---

## 📚 API Documentation

### Authentication
All protected endpoints require a JWT token in the `Authorization` header:
```
Authorization: Bearer <JWT_TOKEN>
```
Admin-only endpoints require the JWT token of a user with `"role": "admin"`.

---

### Student Endpoints

#### Register (Signup)
**POST** `/api/students/signup`
```json
{
  "name": "Test User",
  "email": "test.user@example.com",
  "password": "password123",
  "phone": "1234567890",
  "dateOfBirth": "2000-01-01",
  "gender": "male",
  "address": {
    "street": "123 Main St",
    "city": "New Delhi",
    "state": "DL",
    "zipCode": "110001",
    "country": "India"
  }
}

![image](https://github.com/user-attachments/assets/48ca2006-1392-4ce5-b08e-0b9df7fad5e9)

```

#### Student Login
**POST** `/api/students/login`
```json
{
  "email": "test.user@example.com",
  "password": "password123"
}

```

#### Get Profile
**GET** `/api/students/me`
- **Headers:** `Authorization: Bearer <STUDENT_JWT_TOKEN>`

![image](https://github.com/user-attachments/assets/da9535fa-8b66-4532-8c1c-b5fd161dbd01)

#### Update Profile
**PUT** `/api/students/me`
- **Headers:** `Authorization: Bearer <STUDENT_JWT_TOKEN>`
```json
{
  "name": "Test User Updated",
  "phone": "9876543210",
  "address": {
    "street": "456 New St",
    "city": "Bangalore",
    "state": "KA",
    "zipCode": "560001",
    "country": "India"
  }
}

![image](https://github.com/user-attachments/assets/6aa87a8c-85ef-4263-b926-6180bde39071)

```

---

### Admin Endpoints

#### Create Admin
**POST** `/api/students/create-admin`
- **Headers:** `X-Admin-Secret: <your-admin-secret>`
```json
{
  "name": "Test Admin",
  "email": "test.admin@example.com",
  "password": "adminpassword",
  "phone": "9876543210",
  "dateOfBirth": "1990-01-01",
  "gender": "female",
  "address": {
    "street": "456 Admin St",
    "city": "Mumbai",
    "state": "MH",
    "zipCode": "400001",
    "country": "India"
  }
}

![image](https://github.com/user-attachments/assets/a698978c-3f65-469f-94c0-b9b27f63cdfa)

```

#### Admin Login
**POST** `/api/students/login`
```json
{
  "email": "test.admin@example.com",
  "password": "adminpassword"
}
```

#### List Admins
**GET** `/api/students/admins`
- **Headers:** `Authorization: Bearer <ADMIN_JWT_TOKEN>`

![image](https://github.com/user-attachments/assets/e0217d96-9786-4176-a77b-284f9818cdf5)



![image](https://github.com/user-attachments/assets/07665231-a188-4826-bf3d-dfdb1d556629)

---

### Course Endpoints

#### Create Course (Admin Only)
**POST** `/api/courses`
- **Headers:** `Authorization: Bearer <ADMIN_JWT_TOKEN>`
```json
{
  "name": "Computer Science",
  "description": "Bachelor of Computer Science",
  "duration": "4 years",
  "seats": 50,
  "eligibilityCriteria": {
    "minimumPercentage": 75,
    "requiredSubjects": ["Mathematics", "Physics"],
    "entranceExam": true
  },
  "fees": {
    "tuitionFee": 50000,
    "admissionFee": 5000,
    "otherFees": 2000
  }
}

![image](https://github.com/user-attachments/assets/0679e654-2b59-4bd1-b806-ba519512b1d6)

```

#### Get All Courses
**GET** `/api/courses`
- **Headers:** `Authorization: Bearer <STUDENT_JWT_TOKEN>`

![image](https://github.com/user-attachments/assets/7e0e6477-df3d-4820-a11b-df8538f15062)


#### Get Course by ID
**GET** `/api/courses/:id`
- **Headers:** `Authorization: Bearer <STUDENT_JWT_TOKEN>`

![image](https://github.com/user-attachments/assets/abc4fdec-0c8a-4769-8b8a-7570e8270fca)


#### Update Course (Admin Only)
**PUT** `/api/courses/:id`
- **Headers:** `Authorization: Bearer <ADMIN_JWT_TOKEN>`
```json
{
  "name": "Computer Science Updated",
  "description": "Updated Bachelor of Computer Science",
  "seats": 60,
  "fees": {
    "tuitionFee": 55000,
    "admissionFee": 5500,
    "otherFees": 2500
  }
}
```

#### Delete Course (Admin Only)
**DELETE** `/api/courses/:id`
- **Headers:** `Authorization: Bearer <ADMIN_JWT_TOKEN>`



![image](https://github.com/user-attachments/assets/97b89a76-758e-438f-aea2-25e2c47006b4)

---

### Admission Endpoints

#### Apply for Admission
**POST** `/api/admissions`
- **Headers:** `Authorization: Bearer <STUDENT_JWT_TOKEN>`
```json
{
  "courseId": "64e8b2f4c2a4e2b1a1c2d3e4",
  "personalDetails": {
    "firstName": "Test",
    "lastName": "User",
    "email": "test.user@example.com",
    "phone": "1234567890",
    "dateOfBirth": "2000-01-01",
    "gender": "male",
    "nationality": "Indian",
    "address": {
      "street": "123 Main St",
      "city": "New Delhi",
      "state": "DL",
      "zipCode": "110001",
      "country": "India"
    }
  },
  "academicDetails": {
    "highestQualification": "Bachelor",
    "institution": "Test University",
    "yearOfCompletion": 2022,
    "percentage": 85.5,
    "documents": ["transcript.pdf", "degree.pdf"]
  },
  "documents": {
    "photo": "photo.jpg",
    "idProof": "id.pdf",
    "addressProof": "address.pdf",
    "qualificationCertificates": ["cert1.pdf", "cert2.pdf"]
  }
}

![image](https://github.com/user-attachments/assets/c94576a4-862d-4308-b13c-cdf13095c52f)

```

#### Get All Admissions
**GET** `/api/admissions`
- **Headers:** `Authorization: Bearer <STUDENT_JWT_TOKEN>`

#### Get Admission by ID
**GET** `/api/admissions/:id`
- **Headers:** `Authorization: Bearer <STUDENT_JWT_TOKEN>`

![image](https://github.com/user-attachments/assets/4c10adae-8956-43e0-8b74-d865962fc8b2)


#### Update Admission Status (Admin Only)
**PUT** `/api/admissions/:id`
- **Headers:** `Authorization: Bearer <ADMIN_JWT_TOKEN>`
```json
{
  "status": "approved",
  "comments": "Application approved after review"
}


![image](https://github.com/user-attachments/assets/dadf5d8e-4439-4109-9119-d924633d8faf)
![image](https://github.com/user-attachments/assets/d3b513d0-0238-4083-a6cf-0d944ce95217)


![image](https://github.com/user-attachments/assets/7ff22180-ee1b-4856-8ec9-33a96f8e78d9)



```

---

## ❗ Error Responses
- **Unauthorized (missing/invalid token):**
  ```json
  { "error": "Authorization header is required" }
  ```
- **Forbidden (not admin):**
  ```json
  { "error": "Only admin users can perform this action. If you believe this is a mistake, please contact support." }
  ```
- **Validation errors:**
  ```json
  { "error": "Field validation for 'Email' failed on the 'required' tag" }
  ```
  or
  ```json
  { "error": "Field validation for 'Password' failed on the 'min' tag" }
  ```
  or
  ```json
  { "error": "Field validation for 'Status' failed on the 'oneof' tag" }
  ```
- **Resource not found:**
  ```json
  { "error": "Course not found" }
  ```
  or
  ```json
  { "error": "Admission not found" }
  ```
- **Bad request (malformed JSON, invalid ObjectID, etc.):**
  ```json
  { "error": "Invalid course ID" }
  ```
  or
  ```json
  { "error": "Invalid user ID" }
  ```
- **Internal server error:**
  ```json
  { "error": "Error while creating student" }
  ```
  or
  ```json
  { "error": "Error while applying for admission" }
  ```

---

## 🧪 Testing & Development

### Local Development
1. Install dependencies:
   ```bash
   go mod download
   ```
2. Run locally:
   ```bash
   go run cmd/main.go
   ```

### Testing
```bash
go test ./...
```

---


## ⚠️ MongoDB Credentials Notice
> Security & Best Practice:
> For demonstration purposes, this project is configured with a temporary MongoDB connection string.
>
> We strongly recommend that you use your own MongoDB Atlas credentials for any personal, development, or production deployment.
>
> - Update the MONGODB_URI value in your .env file with your own connection string.
> - This ensures your data remains private, secure, and fully under your control.
>
> Sharing temporary credentials is a common practice for open-source demos, but always use your own database for real applications.

---
