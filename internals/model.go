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

// Job represents a previous job or internship experience while in school or Not
type Company struct {
	ID                      string       `bson:"_id,omitempty"`
	Name                    string       `bson:"name"`
	Industry                string       `bson:"industry"`
	Website                 string       `bson:"website"`
	Email                   string       `bson:"email"`
	Phone                   string       `bson:"phone"`
	Street                  string       `bson:"street"`
	City                    string       `bson:"city"`
	State                   string       `bson:"state"`
	Country                 string       `bson:"country"`
	PostalCode              string       `bson:"postal_code"`
	Founded                 time.Time    `bson:"founded"`
	NumberOfEmployees       int          `bson:"number_of_employees"`
	Revenue                 float64      `bson:"revenue"`
	Description             string       `bson:"description"`
	JobOpenings             []Job        `bson:"job_openings"`
	InternshipOpportunities []Internship `bson:"internship_opportunities"`
	RecruitmentContact      string       `bson:"recruitment_contact"`
	PreferredSkills         []string     `bson:"preferred_skills"`
	HiringProcess           string       `bson:"hiring_process"`
	MissionStatement        string       `bson:"mission_statement"`
	Values                  []string     `bson:"values"`
	Perks                   []string     `bson:"perks"`
	WorkEnvironment         string       `bson:"work_environment"`
	Technologies            []string     `bson:"technologies"`
	Products                []string     `bson:"products"`
	Projects                []Project    `bson:"projects"`
	LinkedIn                string       `bson:"linkedin"`
	Twitter                 string       `bson:"twitter"`
	Facebook                string       `bson:"facebook"`
	OtherLinks              []string     `bson:"other_links"`
	SocialResponsibility    string       `bson:"social_responsibility"`
	EnvironmentalImpact     string       `bson:"environmental_impact"`
	LegalStructure          string       `bson:"legal_structure"`
	Awards                  []string     `bson:"awards"`
	Partners                []string     `bson:"partners"`
	MarketShare             float64      `bson:"market_share"`
	FinancialHealth         string       `bson:"financial_health"`
	CustomerSatisfaction    float64      `bson:"customer_satisfaction"`
	MarketPresence          []string     `bson:"market_presence"`
	HistoricalMilestones    []string     `bson:"historical_milestones"`
	LeadershipTeam          []Person     `bson:"leadership_team"`
	FutureGoals             string       `bson:"future_goals"`
}


// Job represents a job opening.
type Job struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Location    string `bson:"location"`
}


// Project represents a notable project or initiative.
type Project struct {
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	StartDate   time.Time `bson:"start_date"`
	EndDate     time.Time `bson:"end_date"`
}


// Person represents a member of the leadership team.
type Person struct {
	FullName  string `bson:"full_name"`
	Position  string `bson:"position"`
	LinkedIn  string `bson:"linkedin"`
	Biography string `bson:"biography"`
}

// Internship represents an internship opportunity at the company.
type Internship struct {
	Title          string    `bson:"title"`
	Description    string    `bson:"description"`
	Location       string    `bson:"location"`
	PostedDate     time.Time `bson:"posted_date"`
	ApplicationURL string    `bson:"application_url"`
}

// Mail represents the structure of the mail message sent to the users email address.
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
