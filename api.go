package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "math/rand"
)


//declaring structures
type Thread struct{
	Thread_id		string	`json:"thread_id,omitempty"`
	Description		string	`json:"description,omitempty"`
	Name			string	`json:"name,omitempty"`
	Creation_DateTime	string	`json:"creation_DateTime,omitempty"`
}

type Post struct{
	Post_id			string	`json:"post_id,omitempty"`
	Thread_id 		string	`json:"thread_id,omitempty"`
	Creation_DateTime	string	`json:"creation_DateTime,omitempty"`
	Content			string	`json:"content,omitempty"`
}

type Response struct{
	Status			string `json:"status"`
	Message			string `json:"message"`
}

//declaring slice
var thread []Thread
var all_post []Post
var thread_post []Post

//function to get all threads 
func GetAllThreads(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(thread)
}

//function to create a thread
func CreateThread(w http.ResponseWriter, r *http.Request){
        
        name := r.FormValue("name") //name to be received as POST request
        description := r.FormValue("description") //description to be received as POST request
        thread = append(thread, Thread{Thread_id: strconv.Itoa(rand.Intn(100)), Description: description, Name: name, Creation_DateTime: "9-04-2019 14:46:35"})
        
        t := Response{"200", "success"}
        json.NewEncoder(w).Encode(t)
}

//funtion to get a single thread
func GetThread(w http.ResponseWriter, r *http.Request){
	
        thread_id := r.FormValue("thread_id") //thread_id to be received as POST request
	
	for _, item := range thread {
		if item.Thread_id == thread_id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//function to add post to a thread
func AddPostToThread(w http.ResponseWriter, r *http.Request){
        
        thread_id := r.FormValue("thread_id") //thread_id to be received as POST request
        content := r.FormValue("content") //content to be received as POST request
        all_post = append(all_post,Post{Post_id: strconv.Itoa(rand.Intn(100)), Thread_id: thread_id, Creation_DateTime: "9-04-2019 14:46:35", Content:content})
        
        t := Response{"200", "success"}
        json.NewEncoder(w).Encode(t)
}

//function to get all posts of a thread
func GetAllPostOfThread(w http.ResponseWriter, r *http.Request){
	
        thread_id := r.FormValue("thread_id") //thread_id to be received as POST request
	
	for _, item := range all_post {
		if item.Thread_id == thread_id {
			thread_post = append(thread_post, item)
		}
	}
	
	json.NewEncoder(w).Encode(thread_post)
}


func main() {


	//filling up dummy data
	thread = append(thread, Thread{Thread_id: "1", Description: "First Thread", Name: "First", Creation_DateTime: "9-04-2019 14:45:35"})	
	thread = append(thread, Thread{Thread_id: "2", Description: "Second Thread", Name: "Second", Creation_DateTime: "9-04-2019 14:46:35"})	
	all_post = append(all_post, Post{Post_id: "1", Thread_id:"1", Creation_DateTime: "9-04-2019 14:46:35", Content:"content_content"})
	all_post = append(all_post, Post{Post_id: "2", Thread_id:"1", Creation_DateTime: "9-04-2019 14:46:35", Content:"content_content"})
	all_post = append(all_post, Post{Post_id: "1", Thread_id:"2", Creation_DateTime: "9-04-2019 14:46:35", Content:"content_content"})
	all_post = append(all_post, Post{Post_id: "2", Thread_id:"2", Creation_DateTime: "9-04-2019 14:46:35", Content:"content_content"})
	
	//url patterns and their handlers
	http.HandleFunc("/get_all_threads", GetAllThreads)
	http.HandleFunc("/create_thread", CreateThread)
	http.HandleFunc("/get_thread", GetThread)
	http.HandleFunc("/add_post_to_thread", AddPostToThread)
	http.HandleFunc("/get_all_post_of_thread", GetAllPostOfThread)
	
	
	fmt.Printf("Starting server for HTTP...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	
}
