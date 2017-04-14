package main 

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "bufio"
    "github.com/lucjohnson/challenges-lucjohnson/suggestion-challenge/trie"
    "encoding/json"
    "runtime"
    "strconv"
)

var t = trie.NewTrie()

var loadedFile = "./data/wordsEn.txt"

var trieLoaded = false

var memstats = new(runtime.MemStats)

// SuggestionResponse is used to send to the client suggestions gathered from the trie based on the prefix as well as the file used to build the trie
type SuggestionResponse struct {
    Prefix string
    Suggestions []string
    File string
}

// LoadingResponse sends back a message to let a user know that the trie is still loading
type LoadingResponse struct {
    Message string
}

// Function used by the api route, calls the tries FindEntries method and returns suggestions in JSON
func getSuggestions(w http.ResponseWriter, r *http.Request) {
    prefix := r.URL.Query().Get("prefix")
    max, e := strconv.Atoi(r.URL.Query().Get("max"))
    if e != nil {
        log.Fatal(e)
    }
    if trieLoaded {
        suggestions := t.FindEntries(prefix, max)
        resp := SuggestionResponse{Prefix: prefix, Suggestions: suggestions, File: loadedFile}
        
        j, err := json.Marshal(resp)
        if nil != err {
            log.Println(err)
            w.WriteHeader(500)
            w.Write([]byte(err.Error()))
        } else {
            w.Header().Add("Content-Type", "application/json")
            w.Write(j)
        }
    } else {
        resp := LoadingResponse{Message: "Data is still loading, please wait"}
        j, err := json.Marshal(resp)
        if nil != err {
            log.Println(err)
            w.WriteHeader(500)
            w.Write([]byte(err.Error()))
        } else {
            w.WriteHeader(409)
            w.Header().Add("Content-Type", "application/json")
            w.Write(j)
        }
    }
}

// Reads from a given file and puts each entry into a trie data structure
func loadTrie() {
    if len(os.Args) > 1 {
        loadedFile = os.Args[1]
    }
    file, err := os.Open(loadedFile)
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if t.GetSize() % 10000 == 0 {
            runtime.ReadMemStats(memstats)
        }
        if memstats.Alloc > 2500000000 {
            break
        }
        t.AddEntry(scanner.Text())
    }
    file.Close()
    trieLoaded = true
}

// Serves static files, while loading the trie in the background
func main() {
    go loadTrie()
    
    http.Handle("/", http.FileServer(http.Dir("static")))
    http.HandleFunc("/api/v1/suggestions", getSuggestions)
    fmt.Println("Server is listening on port 9000 gurlllll")
    http.ListenAndServe(":9000", nil)
}