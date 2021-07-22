package main

import faktory "github.com/contribsys/faktory/client"

func main() {
	println("Pushing my first job")
	client, err := faktory.Open()
	if err != nil {
		panic(err)
	}

	job := faktory.NewJob("SomeJob", "hello world")
	err = client.Push(job)

	if err != nil {
		panic(err)
	}
	println("Job is pushed")
}
