# URL Shortener

A simple URL shortener built with **Go (Golang)** for the backend and **HTML, CSS, and JavaScript** for the frontend.

## Features

- Shorten long URLs into short, easy-to-share links.
- Automatically redirect users when they visit a short URL.
- Custom error page for invalid short URLs.
- Simple and responsive frontend.

## Project Structure

```
├── frontend
│   ├── static
│   │   ├── styles.css
│   │   ├── script.js
│   ├── template
│   │   ├── index.html
│   │   ├── error.html
├── handler
│   ├── shortenurl.go
│   ├── redirecthandler.go
│   ├── errorpage.go
|   ├── generatekey.go
├── main.go
├── README.md
├── go.mod
```

## Installation

### Prerequisites

Ensure you have **Go** installed on your system. You can download it from [golang.org](https://go.dev/).

### Steps

1. Clone the repository:
   ```sh
   git clone https://github.com/bernotieno/url-shortener.git
   cd url-shortener
   ```
2. Run the Go server:
   ```sh
   go run main.go
   ```
3. Open your browser and go to:
   ```
   http://localhost:9000
   ```

## API Endpoints

### Shorten URL

**POST** `/api/shorten`

- **Request Body (JSON):**
  ```json
  { "url": "https://example.com" }
  ```
- **Response (JSON):**
  ```json
  { "short_url": "/abc123" }
  ```

### Redirect to Original URL

- Visiting `http://localhost:9000/abc123` redirects to `https://example.com`.

## Error Handling

- If a non-existent short URL is accessed, a custom error page is shown instead of the default HTTP error message.

## Deployment
Find the application online on: https://url-shortener-fj08.onrender.com
