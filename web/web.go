package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nii236/gopher-resume/models"
	"github.com/nii236/gopher-resume/process"
)

// Serve begins the REST API that generates CVs
func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", cvHandle)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln(err)
	}

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
