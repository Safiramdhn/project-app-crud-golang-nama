# Choco Academy - Student Enrollment System

A Go-based CLI application for managing student enrollments in an academic setting.

## Overview

Choco Academy is a student enrollment management system that allows students to:
- Login to their accounts
- View their enrollments
- Edit class schedules
- Add new classes
- Delete enrollments
- View enrollment history

## Project Structure
project-app-crud-golang-safira/ 
├── models/ 
│ ├── classes/ 
│ │ ├── class_model.go 
│ │ └── classes.json 
│ ├── courses/ 
│ │ ├── course_model.go 
│ │ └── courses.json 
│ ├── enrollments/ 
│ │ ├── enrollment_models.go 
│ │ └── enrollments.json 
│ ├── schedules/ 
│ │ ├── schedule_model.go 
│ │ └── schedules.json 
│ └── students/ 
│ ├── student_model.go
│ └── students.json 
├── services/ 
│ ├── class_services.go 
│ ├── course_services.go
│ ├── enrollment_services.go 
│ ├── schedule_services.go 
│ └── student_services.go 
├── utils/ 
│ └── utils.go 
├── views/ 
│ ├── class_view.go 
│ ├── enrollment_view.go 
│ ├── main_view.go 
│ └── schedule_view.go
└── README.md

## Features
### Authentication
- Student login with ID and password
- Session management with timeout

### Enrollment Management
- View active enrollments
- Add new class enrollments
- Edit class schedules
- Delete enrollments
- View enrollment history

### Class Management
- View available classes
- View class schedules
- Schedule conflict detection

## Data Models

### Student
- ID
- Name
- Email
- Password

### Class
- ID
- Title
- Description
- Type
- Courses
- Instructor
- Schedules

### Course
- ID
- Title
- Description

### Schedule
- ID
- Day
- Time

### Enrollment
- ID
- Student
- Class
- Schedule
- Status

## Getting Started

### Prerequisites
- Go 1.x or higher
- Git

### Installation
1. Clone the repository
```bash git clone https://github.com/Safiramdhn/project-app-crud-golang-safira.git```
2. Navigate to the project directory
```cd project-app-crud-golang-safira```
3. Run the application
```go run main.go```

## Usage
1. Login using your student credentials
```
Student ID: S001
Password: aniketoss001
```
2. Navigate through the menu options:
  1: View Enrollments
  2: Edit enrollments schedule
  3: Add Class
  4: Delete Enrollment
  5: Enrollment History
  99: Logout

# #Data Storage
The application uses JSON files to store data:
- students.json: Student information
- courses.json: Course information
- classes.json: Class information
- schedules.json: Schedule information
- enrollments.json: Enrollment information

## Contributing
1. Fork the repository
2. Create your feature branch (git checkout -b feature/AmazingFeature)
3. Commit your changes (git commit -m 'Add some AmazingFeature')
4. Push to the branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

## Authors
- Safira - Initial work - [Safiramdhn]([url](https://github.com/Safiramdhn/))
