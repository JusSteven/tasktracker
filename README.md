# Go CLI Task Manager

A lightweight, concurrent-safe Command Line Interface (CLI) application built from scratch in Go (Golang) for managing daily tasks. This tool features persistent storage using structured JSON flat-files, safe data manipulation with pointers, and a dual-mode file execution system.

## Features

* **Task Instantiation:** Quickly spin up tasks with custom descriptions and automatic `"todo"` staging statuses.
* **Dual-Mode Saving Architecture:** * **Create Fresh:** Spin up a brand new, isolated file seeded with a clean JSON array list.
  * **Smart Append:** Read, unmarshal, and merge tasks into an existing file without corrupting historical data.
* **Dynamic Auto-Incrementing IDs:** Safely inspects file length bounds to calculate the next sequential task ID, protecting against memory index out-of-range panics.
* **Human-Readable JSON Indentation:** Utilizes `json.MarshalIndent` formatting to pretty-print records vertically rather than in single-line minified buffers.
* **Robust Error Handling:** Strict pointer manipulation and memory-safety check validation lines to guarantee data persistence.

---

## Tech Stack & Concepts

* **Language:** Go (Golang)
* **Core Mechanics:** Pointers (`&`), Composite Literals, Slice Optimization (`append`), Error Propagation, File Stream I/O (`os.ReadFile`, `os.WriteFile`).
* **Data Layout:** Native JSON Marshalling / Unmarshalling Deserialization.

---

## Project Directory Structure

```text
├── start.go      # Primary application entry point, argument router, and file-handling logic.
├── README.md     # Project documentation.
└── *.json        # Dynamically generated user task database files.
```

## Installation and Environment Setup

### 1. Prerequisites

Ensure you have the Go runtime installed in your local environment (e.g. Linux or Windows)

```bash
go version
```

### 2. Clone the repository

```bash
git clone git@github.com:JusSteven/tasktracker.git
cd tasktracker
```

## Command Usage and Operations Guide

### Adding a Brand New Task

To add a new record to your tracking file, invoke the application using the `add` subcommand followed by your desired task string.

```bash
go run start.go add "Complete Go storage data pipeline"
```

#### Interactive Terminal Flow

**Adding task: Complete Go storage data pipeline**

**Do you want to add to an existing file or create a new file?
Press 1 to add to an existing file or 2 to create a new task file: 1**

**Enter the existing file name: tasks.json
Task saved to the file successfully!**

### Updating The Status of a Task

```bash
go run start.go update
```

**Enter the file name containing the task:****
tasks.json

**Enter the Task ID you want to update:****
1

**Enter the new status (e.g., in-progress, done):****
done

**Task 1 successfully updated to status: 'done'!**

### Listing all tasks

To see a clean read-out of all tasks currently stored inside a specific file, invoke the tool using the `list` subcommand

```bash
go run start.go list
```
