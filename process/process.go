package process

import (
	"bytes"

	"text/template"

	"github.com/nii236/gopher-resume/models"
)

// ProcessCV will return a string representation of a candidate's CV
func ProcessCV(resume *models.Resume) bytes.Buffer {
	cvTemplate := `
Name: {{ .BasicInformation.Name }}
Label: {{ .BasicInformation.Label }}
Picture: {{ .BasicInformation.Picture }}
Email: {{ .BasicInformation.Email }}
Phone: {{ .BasicInformation.Phone }}
Website: {{ .BasicInformation.Website }}
Summary: {{ .BasicInformation.Summary }}
Location: {{ .BasicInformation.Location }}
Profiles: {{ .BasicInformation.Profiles }}
`
	t := template.Must(template.New("cv").Parse(cvTemplate))
	var b bytes.Buffer
	t.Execute(&b, resume)
	return b
}
