package resume

type Resume struct {
	BasicInformation    Basic         `json:"basics"`
	WorkExperience      []Work        `json:"work"`
	VolunteerExperience []Volunteer   `json:"volunteer"`
	EducationHistory    []Education   `json:"education"`
	Awards              []Award       `json:"awards"`
	Publications        []Publication `json:"publications"`
	Skills              []Skill       `json:"skills"`
	Languages           []Language    `json:"languages"`
	Interests           []Interest    `json:"interests"`
	References          []Reference   `json:"references"`
}

type Basic struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Picture  string    `json:"picture"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Website  string    `json:"website"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	Profiles []Profile `json:"profiles"`
}

type Location struct {
	Address     string `json:"address"`
	PostCode    string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type Work struct {
	Company    string   `json:"company"`
	Position   string   `json:"position"`
	Website    string   `json:"website"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

type Volunteer struct {
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	Website      string   `json:"website"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}
type Education struct {
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	GPA         string   `json:"gpa"`
	Courses     []string `json:"courses"`
}
type Award struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}
type Publication struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate"`
	Website     string `json:"website"`
	Summary     string `json:"summary"`
}
type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}
type Language struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}
type Interest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}
type Reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}
