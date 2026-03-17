# TaskManager

A brutalist, essential, and extremely fast task manager. 
No useless charts, no distracting notifications, and no heavy JavaScript frameworks. Just your tasks and a button to complete them.

Live Demo: [Try it on Vercel](https://task-manager-eight-brown-94.vercel.app/)

## Tech Stack
Built with pure Server-Side Rendering (SSR) for maximum performance:
* Backend: Go (Golang) + chi router
* Frontend: Templ (Typed HTML compiled in Go) + Tailwind CSS 4
* Database: PostgreSQL (with pgx driver)
* Deployment: Vercel

## Core Features
* Zero Distractions: High-contrast UI. You either complete the task, or you close the app.
* Real Authentication: Login and Registration with hashed passwords (bcrypt) and secure cookies for production environments.
* Kanban Dashboard: 3-column layout (Pending, Doing, Completed). Start and finish your tasks with a seamless flow.
* Focus Mode: Integrated Pomodoro Timer (25/5 min) isolated from the rest of the application.
* Live Statistics: The landing page fetches real-time data from the database to display the exact number of registered users.

## Local Setup
Want to run the project on your machine? Follow these steps:

1. Clone the repository:
   ```bash
   git clone [https://github.com/yeddaTech/TaskManager.git](https://github.com/yeddaTech/TaskManager.git)
   cd TaskManager
   ```

2. Prerequisites: Make sure you have Go, [Templ](https://templ.guide/), and a local PostgreSQL database running.

3. Generate the Templ HTML files:
   ```bash
   templ generate
   ```

4. Start the server:
   ```bash
   go run api/index.go
   ```
   *(Tip: use `modd` if you want live-reload while writing code).*

Open `http://localhost:3000` in your browser and you are ready to go.

---
Developed by [Younesse Eddassouli (@yeddaTech)](https://github.com/yeddaTech)
