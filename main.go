package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/subutux/waybar-powerprofiles/pkg/powerprofiles"

	waybar "github.com/lack/gowaybarplug"
)

var set = ""
var next = false

func loop(interval time.Duration, pp *powerprofiles.Profiles) {
	wb := waybar.NewUpdater()
	for true {

		status := waybar.Status{}
		profile, err := pp.GetActiveProfile()
		log.Printf(profile)
		if err != nil {
			status.Alt = "error"
			status.Class = []string{"error"}
			status.Tooltip = fmt.Sprintf("Cannot get active profile: %v", err)
		}
		status.Text = profile
		status.Class = []string{profile}

		wb.Status <- &status

		time.Sleep(interval)

	}
}

func main() {

	flag.Parse()
	powerProfiles, err := powerprofiles.NewProfiles()
	if err != nil {
		fmt.Printf("Failed to initiate: %v\n", err)
		os.Exit(1)
	}
	if set != "" {
		err := powerProfiles.SetProfile(set)
		if err != nil {
			fmt.Printf("Failed to set profile to %s: %v\n", set, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	log.Printf(" next: %v", next)

	if next {
		profiles, err := powerProfiles.GetProfiles()
		if err != nil {
			fmt.Printf("Failed to get profiles: %v\n", err)
			os.Exit(1)
		}

		profile, err := powerProfiles.GetActiveProfile()
		if err != nil {
			fmt.Printf("Failed to get active profile: %v\n", err)
			os.Exit(1)
		}
		log.Printf(" Profiles: %v", profiles)
		for idx, p := range profiles {
			if profile == p.Profile {
				var i = idx + 1
				if i == len(profiles) {
					i = 0
				}
				err := powerProfiles.SetProfile(profiles[i].Profile)
				if err != nil {
					fmt.Printf("Failed to get active profile: %v\n", err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		}

		os.Exit(1)

	}

	loop(5*time.Second, powerProfiles)

}

func init() {
	flag.StringVar(&set, "set", "", "Set the PowerProfile")
	flag.BoolVar(&next, "next", false, "set the next PowerProfile")
}
