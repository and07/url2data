package main

//Geter ...
type Geter interface {
	Get(string) (*string, error)
}
