package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nii236/gopher-resume/aaa"
	"github.com/nii236/gopher-resume/models"
	"github.com/nii236/gopher-resume/process"
)

// Serve begins the REST API that generates CVs
func Serve() {
	token := os.Getenv("TOKEN")
	port := os.Getenv("PORT")
	ip := os.Getenv("IP")
	if token == "" {
		token = aaa.Generate(2, "-")
	}
	if ip == "" {
		ip = "0.0.0.0"
	}
	if port == "" {
		port = "8080"
	}

	fmt.Println("Welcome to Gopher Resume!")
	fmt.Println("Your admin token is:", token)
	fmt.Println("Hosting on:", ip+":"+port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", me)

	if err := http.ListenAndServe(ip+":"+port, mux); err != nil {
		log.Fatalln(err)
	}

}

func me(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(models.Me)
	if err != nil {
		fmt.Println("Could not marshal struct into JSON")
		w.WriteHeader(400)
		return
	}
	fmt.Fprint(w, string(resp))
}

func cvHandle(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	resume := &models.Resume{BasicInformation: models.Basic{
		Name:    vals.Get("name"),
		Label:   vals.Get("label"),
		Picture: vals.Get("picture"),
		Email:   vals.Get("email"),
		Phone:   vals.Get("phone"),
		Website: vals.Get("website"),
		Summary: vals.Get("summary"),
	},
	}

	result := process.ProcessCV(resume)
	fmt.Fprint(w, result.String())
}
