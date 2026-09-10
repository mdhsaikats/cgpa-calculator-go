<div align="center">

# CGPA Calculator Go

Calculate your cumulative grade point average from semester GPAs and credit hours.

[![Build](https://img.shields.io/github/actions/workflow/status/mdhsaikats/cgpa-calculator-go/go.yml?branch=main&label=build)](https://github.com/mdhsaikats/cgpa-calculator-go/actions)
[![Version](https://img.shields.io/badge/version-0.1.0-informational)](https://github.com/mdhsaikats/cgpa-calculator-go/releases)
[![License](https://img.shields.io/badge/license-not--set-lightgrey)](#license--acknowledgements)
[![Go](https://img.shields.io/badge/Go-1.26.4-00ADD8?logo=go&logoColor=white)](https://go.dev/)

</div>

> A lightweight, interactive command-line utility for students who want a quick, credit-weighted CGPA calculation.

> **Visual placeholder:** Add a terminal screenshot or architecture diagram here before publishing, for example `docs/images/terminal-preview.png`.

## Table of Contents

- [About the Project](#about-the-project)
- [Key Features](#key-features)
- [Architecture & Tech Stack](#architecture--tech-stack)
- [Prerequisites](#prerequisites)
- [Getting Started / Installation](#getting-started--installation)
- [Usage & Quickstart](#usage--quickstart)
- [Environment Variables](#environment-variables)
- [Testing & Linting](#testing--linting)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License & Acknowledgements](#license--acknowledgements)

## About the Project

CGPA Calculator Go removes the repetitive arithmetic from calculating a cumulative GPA across multiple semesters. It applies the standard weighted-average formula:

```text
CGPA = sum(SGPA x credit hours) / sum(credit hours)
```

The project exists as a small, dependency-free learning and productivity tool for students. It is useful for checking academic progress locally, validating a manual calculation, or practicing Go fundamentals such as input handling, loops, slices, and numeric calculations.

## Key Features

- **Interactive CLI** - Collects the number of semesters and prompts for each SGPA and credit-hour value.
- **Credit-weighted calculation** - Gives each semester the correct influence based on its credit hours.
- **Two-decimal output** - Prints a readable CGPA rounded to two decimal places.
- **Input validation** - Rejects missing, non-numeric, and non-positive semester counts.
- **Simple exit flow** - Type `exit` at the semester prompt to quit cleanly.

## Architecture & Tech Stack

The application is a single-process CLI with no network, persistence, or external service dependencies.

```text
User input
	|
	v
Interactive Go CLI (main.go)
	|
	v
Weighted CGPA calculation -> Formatted terminal output
```

| Layer          | Technology                                     |
| -------------- | ---------------------------------------------- |
| Frontend       | Terminal prompts and formatted stdout          |
| Backend        | Go 1.26.4 standard library (`fmt`, `strconv`)  |
| Database       | None; calculations are in-memory only          |
| DevOps / Infra | Go toolchain; optional GitHub Actions workflow |

## Prerequisites

- Go `1.26.4` or a compatible newer Go release
- Git, for cloning the repository
- A terminal that supports standard input and output

Verify the Go installation:

```bash
go version
```

## Getting Started / Installation

Clone the repository and enter the project directory:

```bash
git clone https://github.com/mdhsaikats/cgpa-calculator-go.git
cd cgpa-calculator-go
```

Download and verify module dependencies:

```bash
go mod download
go mod verify
```

This project currently uses only the Go standard library, so no third-party dependency installation or database migration is required.

## Usage & Quickstart

Run the calculator directly from the repository root:

```bash
go run .
```

Example session:

```text
Enter Total Semester (or type 'exit' to quit):
3
Enter your SGPA for s :- 1 3.50
Enter your Credit Hour for s :- 1 20
Enter your SGPA for s :- 2 3.80
Enter your Credit Hour for s :- 2 22
Enter your SGPA for s :- 3 3.70
Enter your Credit Hour for s :- 3 21
Your CGPA is: 3.67
```

To build and run a local binary instead:

```bash
go build -o bin/cgpa-calculator .
./bin/cgpa-calculator
```

There is no local development server or API endpoint. The application runs entirely in the terminal.

## Environment Variables

The current CLI does not read environment variables.

| Variable | Type | Required (Yes/No) | Description                                                       | Default |
| -------- | ---- | ----------------- | ----------------------------------------------------------------- | ------- |
| None     | N/A  | No                | Configuration is provided interactively through terminal prompts. | N/A     |

## Testing & Linting

Run the standard Go checks:

```bash
# Run all package tests
go test ./...

# Run static analysis
go vet ./...

# Format Go source files
gofmt -w main.go
```

The repository does not yet include automated test files. For an optional stricter lint pass, install [golangci-lint](https://golangci-lint.run/) and run:

```bash
golangci-lint run
```

## Roadmap

- [x] Interactive semester and credit-hour input
- [x] Credit-weighted CGPA calculation
- [x] Input validation for semester count
- [ ] Validate SGPA and credit-hour values individually
- [ ] Add automated unit tests for calculation and input handling
- [ ] Add non-interactive flags or file input for batch calculations
- [ ] Add continuous integration for formatting, tests, and vetting

## Contributing

Contributions are welcome.

1. Fork the repository and create a focused branch.
2. Make the smallest change that solves the issue.
3. Run `gofmt`, `go test ./...`, and `go vet ./...` before opening a pull request.
4. Describe the user impact and verification steps in the pull request.
5. Open an issue first for larger behavior changes so the approach can be discussed.

Please keep changes focused, document user-visible behavior, and avoid committing generated binaries or local configuration.

## License & Acknowledgements

The project does not currently include a license file. The intended release license is MIT; add a `LICENSE` file and update the badge before publishing.

Built with the [Go standard library](https://pkg.go.dev/std). Thanks to the Go community for the tooling and documentation that make small command-line utilities easy to build and share.
