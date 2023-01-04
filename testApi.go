package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//declaring structures
type Thread struct {
	ThreadID         string `json:"threadID,omitempty"`
	Description      string `json:"description,omitempty"`
	Name             string `json:"name,omitempty"`
	CreationDateTime string `json:"creationDateTime,omitempty"`
}

type Post struct {
	PostID           string `json:"postID,omitempty"`
	ThreadID         string `json:"threadID,omitempty"`
	CreationDateTime string `json:"creationDateTime,omitempty"`
	Content          string `json:"content,omitempty"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

//declaring slice
var thread []Thread
var allPost []Post
var threadPost []Post

//function to get all threads
func GetAllThreads(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(thread)
	} else {
		t := Response{"405", "Method Not Allowed"}
		json.NewEncoder(w).Encode(t)
	}
}

//function to create a thread
func CreateThread(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")               //name to be received as POST request
		description := r.FormValue("description") //description to be received as POST request
		thread = append(thread, Thread{ThreadID: strconv.Itoa(rand.Intn(100)), Description: description, Name: name, CreationDateTime: time.Now().String()})

		t := Response{"200", "success"}

		json.NewEncoder(w).Encode(t)
	} else {
		t := Response{"405", "Method Not Allowed"}
		json.NewEncoder(w).Encode(t)
	}
}

//funtion to get a single thread
func GetThread(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		threadID := r.FormValue("threadID") //ThreadID to be received as Get request

		for _, item := range thread {
			if item.ThreadID == threadID {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	} else {
		t := Response{"405", "Method Not Allowed"}
		json.NewEncoder(w).Encode(t)
	}
}

//function to add post to a thread
func AddPostToThread(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		threadID := r.FormValue("threadID") //ThreadID to be received as POST request
		content := r.FormValue("content")   //content to be received as POST request
		allPost = append(allPost, Post{PostID: strconv.Itoa(rand.Intn(100)), ThreadID: threadID, CreationDateTime: time.Now().String(), Content: content})

		t := Response{"200", "success"}
		json.NewEncoder(w).Encode(t)
	} else {
		t := Response{"405", "Method Not Allowed"}
		json.NewEncoder(w).Encode(t)
	}
}

//function to get all posts of a thread
func GetAllPostOfThread(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		threadID := r.FormValue("threadID") //ThreadID to be received as POST request

		for _, item := range allPost {
			if item.ThreadID == threadID {
				threadPost = append(threadPost, item)
			}
		}

		json.NewEncoder(w).Encode(threadPost)

	} else {
		t := Response{"405", "Method Not Allowed"}
		json.NewEncoder(w).Encode(t)
	}
}

func main() {

	//filling up dummy data
	thread = append(thread, Thread{ThreadID: "1", Description: "First Thread", Name: "First", CreationDateTime: time.Now().String()})
	thread = append(thread, Thread{ThreadID: "2", Description: "Second Thread", Name: "Second", CreationDateTime: time.Now().String()})
	allPost = append(allPost, Post{PostID: "1", ThreadID: "1", CreationDateTime: time.Now().String(), Content: "content_content"})
	allPost = append(allPost, Post{PostID: "2", ThreadID: "1", CreationDateTime: time.Now().String(), Content: "content_content"})
	allPost = append(allPost, Post{PostID: "1", ThreadID: "2", CreationDateTime: time.Now().String(), Content: "content_content"})
	allPost = append(allPost, Post{PostID: "2", ThreadID: "2", CreationDateTime: time.Now().String(), Content: "content_content"})

	//url patterns and their handlers
	http.HandleFunc("/get_all_threads", GetAllThreads)
	http.HandleFunc("/create_thread", CreateThread)
	http.HandleFunc("/get_thread", GetThread)
	http.HandleFunc("/add_post_to_thread", AddPostToThread)
	http.HandleFunc("/get_allPost_of_thread", GetAllPostOfThread)

	fmt.Printf("Starting server for HTTP...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
