# Go CLI Task Manager

A lightweight, concurrent-safe Command Line Interface (CLI) application built from scratch in Go (Golang) for managing daily tasks. This tool features persistent storage using structured JSON flat-files, safe data manipulation with pointers, and a dual-mode file execution system.

##  Features

* **Task Instantiation:** Quickly spin up tasks with custom descriptions and automatic `"todo"` staging statuses.
* **Dual-Mode Saving Architecture:** * **Create Fresh:** Spin up a brand new, isolated file seeded with a clean JSON array list.
    * **Smart Append:** Read, unmarshal, and merge tasks into an existing file without corrupting historical data.
* **Dynamic Auto-Incrementing IDs:** Safely inspects file length bounds to calculate the next sequential task ID, protecting against memory index out-of-range panics.
* **Human-Readable JSON Indentation:** Utilizes `json.MarshalIndent` formatting to pretty-print records vertically rather than in single-line minified buffers.
* **Robust Error Handling:** Strict pointer manipulation and memory-safety check validation lines to guarantee data persistence.

---

##  Tech Stack & Concepts

* **Language:** Go (Golang)
* **Core Mechanics:** Pointers (`&`), Composite Literals, Slice Optimization (`append`), Error Propagation, File Stream I/O (`os.ReadFile`, `os.WriteFile`).
* **Data Layout:** Native JSON Marshalling / Unmarshalling Deserialization.

---

##  Project Directory Structure

```text
├── start.go      # Primary application entry point, argument router, and file-handling logic.
├── README.md     # Project documentation.
└── *.json        # Dynamically generated user task database files.
