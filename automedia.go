package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
  "regexp"
  "errors"
)

var validPath = regexp.MustCompile("^/(editbio|savebio|user)/([a-zA-Z0-9]+)$")
var templates = template.Must(template.ParseFiles(
  "templates/editbio.html",
  "templates/user.html",
))

type User struct {
	Name string
	Bio  []byte
}

func (user *User) saveBio() error {
	fileusername := "bios/" + user.Name + "_bio.txt"
  return os.WriteFile(fileusername, user.Bio, 0600)
}

func main() {
	http.HandleFunc("/user/", makeHandler(userHandler))
	http.HandleFunc("/editbio/", makeHandler(editBioHandler))
	http.HandleFunc("/savebio/", makeHandler(saveBioHandler))
  
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
      println("a")
      http.NotFound(w, r)
      return
    }
    fn(w, r, m[2])
  }
}

func userHandler(w http.ResponseWriter, r *http.Request, username string) {
	user, err := loadBio(username)
  if err != nil {
    http.Redirect(w, r, "/editbio/" + username, http.StatusFound)
    return
  }
	renderTemplate(w, "user", user)
}

func editBioHandler(w http.ResponseWriter, r *http.Request, username string) {
	user, err := loadBio(username)
	if err != nil {
		user = &User{Name: username}
	}
	renderTemplate(w, "editbio", user)
}

func saveBioHandler(w http.ResponseWriter, r *http.Request, username string) {
  bio := r.FormValue("bio")
  print(bio)
  user := &User{Name: username, Bio: []byte(bio)}
  err := user.saveBio()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  http.Redirect(w, r, "/user/" + username, http.StatusFound)
}


func getusername(w http.ResponseWriter, r *http.Request) (string, error) {
  m:= validPath.FindStringSubmatch(r.URL.Path)
  if m == nil {
    http.NotFound(w, r)
    return "", errors.New("invalid Page Title")
  }
  return m[2], nil
}

func loadBio(username string) (*User, error) {
	path := "bios/" + username + "_bio.txt"
	bio, err := os.ReadFile(path)
	if err != nil {
    println("file not found")
		return nil, err
	}
	return &User{username, bio}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, user *User) {
  path :=  tmpl + ".html"
  err := templates.ExecuteTemplate(w, path, user)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
