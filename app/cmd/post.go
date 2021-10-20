package cmd

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/russross/blackfriday.v2"
)

func NewPost(filename string) Post {
	post, err := ExtractMetaFromFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	mdfile, err := os.Open("./markdown/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer mdfile.Close()

	rawBytes, err := ioutil.ReadAll(mdfile)

	// Get title from first line of file
	lines := strings.Split(string(rawBytes), "\n")
	post.Title = strings.Replace(lines[0], "# ", "", -1)

	// Convert Markdown to HTML
	body := blackfriday.Run(rawBytes)
	post.Body = template.HTML(body)

	// Save file
	post.createFile()

	return post
}

type Post struct {
	Title string
	Body  template.HTML
	Date  time.Time
	Slug  string
}

func (post *Post) getURL() string {
	return fmt.Sprintf("/posts/%s", post.Slug)
}

func (post *Post) createFile() {
	// Create folder for HTML
	newPath := filepath.Join("public/posts", post.Slug)
	_ = os.MkdirAll(newPath, os.ModePerm)

	// Create HTML file
	f, err := os.Create("public/posts/" + post.Slug + "/" + "index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Generate final HTML file from template
	t, _ := template.ParseFiles("post.html")
	err = t.Execute(f, post)
	if err != nil {
		log.Fatalf("can't execute template: %v", err)
	}
	_ = f.Close()
}

func fNameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func ExtractMetaFromFile(filename string) (Post, error) {
    finfo, _ := os.Stat(filepath.Join("./markdown/", filename))
    //stat_t := finfo.Sys().(*syscall.Stat_t)
    //ctime := time.Unix(int64(stat_t.Ctim.Sec), int64(stat_t.Ctim.Nsec))
    modTime := finfo.ModTime()
    slug := fNameWithoutExtension(filename)
    return Post{Slug: slug, Date: modTime}, nil
}
