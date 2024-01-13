package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Redirect struct {
    Path   string `yaml:"Path"`
    Target string `yaml:"Target"`
}

type Config struct {
    Redirects []Redirect `yaml:"Redirects"`
}

var redirects map[string]string

func main() {
    yamlFilePath := flag.String("redirects", "", "Path to the redirects YAML config")
    flag.Parse()

    if *yamlFilePath == "" {
        log.Fatal("Please provide a path to a redirects YAML config")
    }

    data, err := os.ReadFile(*yamlFilePath)
    if err != nil {
        log.Fatalf("Failed to read the provided redirects YAML config: %v", err)
    }

    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("Failed to unmarshal the provided redirects YAML config: %v", err)
    }

    redirects = make(map[string]string, len(config.Redirects))

    for _, redirect := range config.Redirects {
        redirects[redirect.Path] = redirect.Target

        fmt.Printf("Registered redirect from: /redirect/%s to %s\n", redirect.Path, redirect.Target)
    }

    http.HandleFunc("/redirect/", redirect)

    fmt.Println("godirect is running on :8080")
    http.ListenAndServe(":8080", nil)
}

func redirect(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[len("/redirect/"):]
    if path == "" {
        http.Error(w, "Path is missing", http.StatusBadRequest)

        return
    }

    targetURL, found := redirects[path]
    if !found {
        http.Error(w, "Invalid redirect path", http.StatusNotFound)

        return
    }

    http.Redirect(w, r, targetURL, http.StatusTemporaryRedirect)
}

