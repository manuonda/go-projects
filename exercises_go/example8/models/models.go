package models

type Instructor struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type WorkShop struct {
	Title       string       `yaml:"title"`
	Instructors []Instructor `yaml:"instructors"`
}
