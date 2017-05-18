package main

import (
	"bytes"
	"flag"
	//"fmt"
	"github.austin.utexas.edu/kriegrj/autorisierungaufkuendigungdienst/authentifizierung"
	"github.austin.utexas.edu/kriegrj/go-github/github"
	"log"
	"os"
	"os/user"
)

func main() {

	var buf bytes.Buffer
	debug := log.New(&buf, "DEBUG: ", log.Lshortfile)
	type Flag struct {
		uPtr *string
		oPtr *string
		err  error
	}
	usr, e := user.Current()
	if e != nil {
		log.Println("Error retrieving current user: %v", e)
	}
	token := os.Getenv("GHE_TOKEN")
	if token == "" {
		log.Fatal("No OAuth token!\n\n")
	} else {
		log.Println("Successfully loaded OAuth token for ", usr.Username)
	}

	client := github.NewClient(authentifizierung.Authn(token))
	f := new(Flag)
	f.rPtr = flag.String("r", "", "repo")
	f.oPtr = flag.String("o", "", "owner")
	f.aPtr = flag.String("a", "", "action")
	flag.Parse()
	switch *f.oPtr {
	case "contributors":
					stats, _, err := client.Repositories.ListContributorsStats(f.oPtr, f.rPtr)
	case "activity":
					stats, _, err := client.Repositories.ListCommitActivity(


	log.Printf("user: %v\trepo: %v\towner: %v\taction: %v", usr, *f.rPtr, *f.oPtr, *f.aPtr)

}