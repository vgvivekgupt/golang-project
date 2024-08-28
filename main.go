package main

import (
    "html/template"
    "net/http"
)

// Define a template for the web page
const pageTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Web Page</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f0f4f8;
            margin: 0;
            padding: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
        }
        .container {
            text-align: center;
            background-color: #fff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #333;
        }
        p {
            color: #555;
        }
        a {
            color: #007bff;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Welcome to My First Web Page</h1>
        <p>This page is served by a Go application!</p>
        <p><a href="https://golang.org">Learn more about Go</a></p>
    </div>
</body>
</html>
`

// Handler function for the root path
func handler(w http.ResponseWriter, r *http.Request) {
    t := template.New("webpage")
    t, _ = t.Parse(pageTemplate)
    t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", handler)
    port := ":8090"  // Updated port
    println("Starting server on port", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        println("Server failed:", err.Error())
    }
}

