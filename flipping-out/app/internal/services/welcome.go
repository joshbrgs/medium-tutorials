package services

import "fmt"

type WelcomeService interface {
	HelloWorld() string
	HelloWorldAgain(user string) string
}

type welcomeService struct {
}

func NewWelcomeService() WelcomeService {
	return &welcomeService{}
}

// Deprecated use HelloWorldAgain
func (s *welcomeService) HelloWorld() string {
	return "HelloWorld"
}

func (s *welcomeService) HelloWorldAgain(user string) string {
	return fmt.Sprintf("Hello there %s", user)
}
