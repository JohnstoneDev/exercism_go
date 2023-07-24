package airportrobot

import "fmt"
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName()string
	Greet(name string)string
}

func SayHello(name string, implementer Greeter) string{
	return fmt.Sprintf("I can speak %v: %v!", implementer.LanguageName(), implementer.Greet(name) )
}

type Italian struct {}

func (i_face Italian) LanguageName() string{
	return "Italian"
}

func (i_face Italian) Greet(name string) string{
	return "Ciao " + name
}

type Portuguese struct {}

func (i_face Portuguese) LanguageName() string{
	return "Portuguese"
}

func (i_face Portuguese) Greet(name string) string{
	return "Olá " + name
}


