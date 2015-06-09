package main

import (
    "bytes"
    "fmt"
    "net/http"
    "net/smtp"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!<br>", r.URL.Path[1:])
    // Connect to the remote SMTP server.
    c, err := smtp.Dial("localhost:25")
    if err != nil {
        fmt.Fprintf(w, "Oh dear 1 - %s", err)
        return
    }
    defer c.Close()
    // Set the sender and recipient.
    c.Mail("admin@chrislhall.net")
    c.Rcpt("root@localhost")
    // Send the email body.
    wc, err := c.Data()
    if err != nil {
        fmt.Fprintf(w, "Oh dear 2 - %s", err)
        return
    }
    defer wc.Close()
    buf := bytes.NewBufferString("Test notification e-mail.")
    if _, err = buf.WriteTo(wc); err != nil {
        fmt.Fprintf(w, "Oh dear 3 - %s", err)
        return
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9001", nil)
}
