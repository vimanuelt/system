# System

The `system` command is a lightweight system information viewer designed for GhostBSD. It provides detailed and up-to-date information about your system's hostname, uptime, CPU, memory (RAM), and storage (disk). The application features a modern, dark-themed graphical interface built with the Fyne framework.

## Features

- Displays key system information:
  - **Hostname**
  - **System Uptime**
  - **CPU Details**
  - **Memory Usage** (with a progress bar)
  - **Storage Usage** (with a progress bar)
- Auto-refreshes every 5 seconds to provide live updates.
- Modern dark-themed GUI.
- Convenient Close button for easy application exit.

## Requirements

- **Operating System:** GhostBSD (or FreeBSD-based systems)
- **Go Programming Language:** Version 1.19 or later
- **Fyne Framework:** Version 2.0 or later
- System libraries:
  - `libX11`
  - `libGLU`
  - `xorg`

Install required dependencies on GhostBSD:
```sh
sudo pkg install xorg libX11 libGLU
```

## Installation

### Build from Source

1. Clone the repository:
   ```sh
   git clone https://github.com/vimanuelt/system.git
   cd system
   ```

2. Build the application using the provided Makefile:
   ```sh
   make
   ```

3. Install the application:
   ```sh
   sudo make install
   ```

The application will be installed to `/usr/local/bin/pcinfo`.

### Run the Application

Run the application from the terminal:
```sh
system
```

## Uninstallation

To remove the application:
```sh
sudo make uninstall
```

## Development

If you would like to contribute or modify the application:

1. Ensure you have Go and the Fyne framework installed.
2. Build and run the application locally:
   ```sh
   go build -o system main.go
   ./system
   ```

## License

This project is licensed under the **BSD 3-Clause License**. See the `LICENSE` file for details.

