package models

// Me is you
var Me = Resume{
	BasicInformation: Basic{
		Name:  "John Nguyen",
		Email: "jtnguyen236@gmail.com",
	},
	WorkExperience:      []Work{},
	VolunteerExperience: []Volunteer{},
	EducationHistory:    []Education{},
	Awards:              []Award{},
	Publications:        []Publication{},
	Skills:              []Skill{},
	Languages:           []Language{},
	Interests:           []Interest{},
	References:          []Reference{},
}
