# TaskFlow – Full-Stack Project Management App

## 1. Overview

TaskFlow is a full-stack project and task management application built using a monorepo architecture. It consists of three core services orchestrated via Docker Compose:

- PostgreSQL – relational database  
- Go Backend API – handles business logic and data access  
- React Frontend UI – user interface for interacting with the system  

### Tech Stack

**Backend**
- Go 1.22
- Gin (HTTP framework)
- GORM (ORM) + PostgreSQL driver
- JWT authentication
- bcrypt (password hashing)
- slog (structured logging)

**Frontend**
- React 19 + TypeScript
- Vite
- React Router DOM
- TanStack React Query
- React Hook Form + Zod
- Tailwind CSS + Radix UI

---

## 2. Architecture Decisions

### Backend Architecture (Clean/Layered)

- controller/ → HTTP handlers only  
- input_mapper/ → request parsing + validation  
- usecase/ → business logic  
- repository/ → DB interaction  
- output_mapper/ → response shaping  
- model/ → domain models  

**Why this structure?**
- Clear separation of concerns
- Business logic independent of frameworks
- Easier testing and maintenance

---

### Frontend Architecture (Feature-First)

Organized by features:
- auth/
- projects/
- tasks/
- users/

Shared:
- lib/ → API + utils
- components/ui/ → reusable UI

**Why?**
- Scalable and modular
- Easier onboarding
---

## 3. Running Locally

### Steps

git clone https://github.com/priyanshi0205/taskflow-priyanshi.git

cd taskflow-priyanshi

docker compose up --build

### Access

App available on : http://localhost:3000  

---

## 4. Running Migrations

Database schema is managed using versioned migration files with up and down scripts.

up file → apply schema changes

down file → rollback schema changes

Migrations run automatically on startup.

---

## 5. Test Credentials

Email:    test@example.com  
Password: password123  

---

## 6. API Reference

### Auth

Register - POST /api/auth/register  

Request:

{

  "email": "test@example.com",
  
  "password": "password123",
  
  "name": "Test User"
  
}

Response:

{

    "id":1,
    
    "email":"test@example.com
    
}

Login - POST /api/auth/login  

Request:

{

    "email":"test@example.com",
    
    "password":"password123"
    
}

Response:

{

    "token" : "jwt-token"
    
}

### Projects

Get Projects - GET /api/projects  


Create Projects - POST /api/projects  

Request:

{

  "name": "My Project",
  
  "description":"Add description"
  
}


### Tasks

Get Tasks - GET /api/projects/:id/tasks

Create Tasks - POST /api/projects/:id/tasks

Request:

{

  "title": "Task title",
  
  "status": "todo",
  
  "description": "Add Task description"
  
}

Update Tasks - PATCH /api/tasks/:id  

---

## 7. What I’d Do With More Time

- Add tests (unit + integration)
- Implement WebSockets

### Final Thoughts

Built with a focus on clean architecture, scalability, and production readiness.
