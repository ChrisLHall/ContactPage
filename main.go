package main

import (
    "fmt"
    "bytes"
    "net/http"
    "net/smtp"
)

func goBack(w http.ResponseWriter, r *http.Request, redirectUrl string) {
    http.Redirect(w, r, redirectUrl, http.StatusFound);
}

func handler(w http.ResponseWriter, r *http.Request) {
    redirectUrl := "http://chrislhall.net/site/contact.html";
    nameVal := r.PostFormValue("inputname");
    emailVal := r.PostFormValue("inputemail");
    msgVal := r.PostFormValue("inputmsg");
    if emailVal == "" || msgVal == "" {
        goBack(w, r, redirectUrl);
	return
    }
    buffer := bytes.NewBufferString("");
    fmt.Fprintf(buffer, "<div><p><strong>Message from Contact page</strong></p>" +
    		"<p>Name: %s</p><p>\nEmail: %s</p><p>\nMessage: %s</p></div>",
    		nameVal, emailVal, msgVal);

    // Connect to the remote SMTP server.
    c, err := smtp.Dial("localhost:25")
    if err != nil {
        goBack(w, r, redirectUrl);
        return
    }
    defer c.Close()
    // Set the sender and recipient.
    c.Mail("admin@chrislhall.net")
    c.Rcpt("root")
    // Send the email body.
    wc, err := c.Data()
    if err != nil {
        goBack(w, r, redirectUrl);
        return
    }
    defer wc.Close()
    header := "Subject: Contact page message\n" +
    	   "MIME-version: 1.0\n" +
	   "Content-Type: text/html; charset=\"UTF-8\"\n";
    wc.Write([]byte(header));
    if _, err = wc.Write(buffer.Bytes()); err != nil {
        goBack(w, r, redirectUrl);
        return
    }
    goBack(w, r, redirectUrl);
}

func main() {
    http.HandleFunc("/leavemsg", handler)
    http.ListenAndServe(":9001", nil)
}
