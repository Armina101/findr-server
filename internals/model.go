package internals

import "time"

// Students represents the details of a student relevant and neccessary for IT placement.
type Students struct {
	DateOfBirth     time.Time `bson:"date_of_birth"`
	GraduationDate  time.Time `bson:"graduation_date"`
	CreatedAt       time.Time `bson:"created_at" db:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at" db:"updated_at"`
	FirstName       string    `bson:"first_name"`
	LastName        string    `bson:"last_name"`
	Email           string    `bson:"email"`
	Password        string    `bson:"password"`
	PhoneNumber     string    `bson:"phone_number"`
	Address         string    `bson:"address"`
	Major           string    `bson:"major"`
	University      string    `bson:"university"`
	PlacementStatus string    `bson:"placement_status"`
	Skills          []string  `bson:"skills"`
	Certifications  []string  `bson:"certifications"`
	Projects        []Project `bson:"projects"`
	Experience      []Job     `bson:"experience"`
	GPA             float32   `bson:"gpa"`
	ID              string    `bson:"_id,omitempty"`
}

// Project represents details of a project the student has completed while in school or relevant Non academics.
type Project struct {
	Name         string   `bson:"name"`
	Description  string   `bson:"description"`
	Technologies []string `bson:"technologies"`
	Link         string   `bson:"link"`
}

// Job represents a previous job or internship experience while in school or Not
type Job struct {
	StartDate        time.Time `bson:"start_date"`
	EndDate          time.Time `bson:"end_date,omitempty"`
	Company          string    `bson:"company"`
	Role             string    `bson:"role"`
	Responsibilities string    `bson:"responsibilities"`
}

type Company struct {
	ID                      string       `bson:"_id,omitempty"`
	Name                    string       `bson:"name"`                     // Full name of the company.
	Industry                string       `bson:"industry"`                 // Sector or industry.
	Website                 string       `bson:"website"`                  // Official website URL.
	Email                   string       `bson:"email"`                    // Contact email address.
	Phone                   string       `bson:"phone"`                    // Contact phone number.
	Street                  string       `bson:"street"`                   // Street address.
	City                    string       `bson:"city"`                     // City.
	State                   string       `bson:"state"`                    // State or province.
	Country                 string       `bson:"country"`                  // Country.
	PostalCode              string       `bson:"postal_code"`              // Postal or ZIP code.
	Founded                 time.Time    `bson:"founded"`                  // Year founded.
	NumberOfEmployees       int          `bson:"number_of_employees"`      // Total number of employees.
	Revenue                 float64      `bson:"revenue"`                  // Annual revenue.
	Description             string       `bson:"description"`              // Brief description.
	JobOpenings             []Job        `bson:"job_openings"`             // Current job openings.
	InternshipOpportunities []Internship `bson:"internship_opportunities"` // Available internships.
	RecruitmentContact      string       `bson:"recruitment_contact"`      // Contact person for recruitment.
	PreferredSkills         []string     `bson:"preferred_skills"`         // Skills valued in candidates.
	HiringProcess           string       `bson:"hiring_process"`           // Description of the hiring process.
	MissionStatement        string       `bson:"mission_statement"`        // Mission statement.
	Values                  []string     `bson:"values"`                   // Core values.
	Perks                   []string     `bson:"perks"`                    // Perks or benefits.
	WorkEnvironment         string       `bson:"work_environment"`         // Work environment and culture.
	TechStack               []string     `bson:"tech_stack"`               // Technologies used.
	Products                []string     `bson:"products"`                 // Key products or services.
	Projects                []Project    `bson:"projects"`                 // Notable projects or initiatives.
	LinkedIn                string       `bson:"linkedin"`                 // LinkedIn profile URL.
	Twitter                 string       `bson:"twitter"`                  // Twitter profile URL.
	Facebook                string       `bson:"facebook"`                
	OtherLinks              []string     `bson:"other_links"`              
}

// Internship represents an internship opportunity at the company.
type Internship struct {
	Title          string    `bson:"title"`           // Internship title.
	Description    string    `bson:"description"`     // Internship description.
	Location       string    `bson:"location"`        // Internship location.
	PostedDate     time.Time `bson:"posted_date"`     // Date the internship was posted.
	ApplicationURL string    `bson:"application_url"` // URL to apply for the internship.
}

type Mail struct {
	Source      string
	Destination string
	Message     string
	Subject     string
	Template    string
	Name        string
}

type LoginStudent struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// // Project represents a notable project the company is involved in.
// type Project struct {
//     Name         string   `bson:"name"`                  // Project name.
//     Description  string   `bson:"description"`           // Brief description of the project.
//     Technologies []string `bson:"technologies"`          // Technologies used in the project.
//     Link         string   `bson:"link"`                  // Link to the project.
// }
