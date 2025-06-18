# IoTask — In-Memory Task Processing Service.

This service allows you to run, track, and delete long tasks, entirely in-memory - without the use of a database, third-party queues, or external infrastructure.

## Installation

1. **Clone the repository:**

```bash
git clone https://github.com/artnikel/iotask.git
cd iotask
```

2. **Build the project:**
   
```bash
go build -o iotask main.go
```

3. **Run the server:**

```bash
./iotask
```

You should see the server running at: `http://localhost:8080`


## API Endpoints

### 1. Create Ping Task
- **URL:** `/tasks/ping`
- **Method:** `POST`
- **Description:** Schedules a TCP ping task to check if a host is reachable on port 80.
- **Response:**
  ```json
  {
    "task_id": "your-generated-id"
  }
  ```

### 2. Get Task Status
- **URL:** `/tasks/{id}`
- **Method:** `GET`
- **Description:** Returns the task status, creation time, current iteration, and execution duration.
- **Response (in-progress):**
  ```json
    {
    "task_id": "task-id",
    "status": "in_progress",
    "created_at": "2025-06-10T12:00:00Z",
    "current_num": 45,
    "duration_seconds": 45.01234
    }
  ```
- **Response (completed):**
   ```json
    {
    "task_id": "task-id",
    "status": "completed",
    "created_at": "2025-06-10T12:00:00Z",
    "completed_at": "2025-06-10T12:03:00Z",
    "current_num": 200,
    "duration_seconds": 200.0,
    "total_duration_seconds": 200.0
    }
  ```

### 3. Delete Task
- **URL:** `/tasks/{id}`
- **Method:** `DELETE`
- **Description:** Deletes a task by its ID.
- **Response (completed):** `204 No Content`


## Logging

The application uses structured logging to track important events and errors. Logs are written to a file configured in the application settings. There are two main loggers:

- **Info logger**: Records informational messages about normal operations.
- **Error logger**: Records errors and warnings to help diagnose issues.

The log file `app.log` is created immediately after the server starts, located in the `logs` directory by default.

## Testing


```bash
make test
```

## Linter


```bash
make lint
```

