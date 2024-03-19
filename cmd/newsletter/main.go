package main

import "TailsOfOld/TailsOfOld/newsletter"

func main() {
	err := newsletter.SendNewsletter()
	if err != nil {
		panic(err)
	}
}
