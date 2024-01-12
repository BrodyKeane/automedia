package main

import (
	"fmt"
  "log"
	"net/http"
	"os"
  "html/template"
)

type User struct {
  name string
  bio []byte
}

func (u *User) save_bio() error {
  filename := u.name + "_bio.txt"
  return os.WriteFile(filename, u.bio, 0600)
}

func loadBio(name string) (*User, error) {
  filename := name + "_bio.txt"
  bio, err := os.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &User{name, bio}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Path[len("/veiw/"):]
  user, _ := loadBio(name)
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", user.name, user.bio)
}

func editBioHandler(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Path[len("/veiw/"):]
  user, err := loadBio(name)
  if err != nil {
    user = &User{name: name}
  }

  fmt.Fprintf(w, "<h1>Editing %s</h1>" +
    "<form action=\"/save/%s\" method=\"{pst\">" +
    "<textarea name=\"body\">%s<textarea><br>" +
    "<input type=\"submit\" value=\"Save\">" +
    "</form>",
    user.name, user.name, user.bio,
  )
}

func main() {
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/editbio/", editBioHandler)
  http.HandleFunc("/savebio/", saveBioHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
