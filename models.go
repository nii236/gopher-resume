package main

type Resume struct {
	BasicInformation    Basic
	WorkExperience      []Work
	VolunteerExperience []Volunteer
	EducationHistory    []Education
	Awards              []Award
	Publications        []Publication
	Skills              []Skill
	Languages           []Language
	Interests           []Interest
	References          []Reference
}

type Basic struct {
	Name      string
	Label     string
	Picture   string
	Email     string
	Phone     string
	Website   string
	Summary   string
	Locations []Location
	Profiles  []Profile
}

type Location struct {
	Address     string
	PostCode    string
	City        string
	CountryCode string
	Region      string
}

type Profile struct {
	Network  string
	Username string
	URL      string
}

type Work struct {
	Company    string
	Position   string
	Website    string
	StartDate  string
	EndDate    string
	Summary    string
	Highlights []string
}

type Volunteer struct {
	Organization string
	Position     string
	Website      string
	StartDate    string
	EndDate      string
	Summary      string
	Highlights   []string
}
type Education struct {
	Institution string
	Area        string
	StudyType   string
	StartDate   string
	EndDate     string
	GPA         string
	Courses     []string
}
type Award struct {
	Title   string
	Date    string
	Awarder string
	Summary string
}
type Publication struct {
	Name        string
	Publisher   string
	ReleaseDate string
	Website     string
	Summary     string
}
type Skill struct {
	Name     string
	Level    string
	Keywords []string
}
type Language struct {
	Name  string
	Level string
}
type Interest struct {
	Name     string
	Keywords []string
}
type Reference struct {
	Name      string
	Reference string
}
