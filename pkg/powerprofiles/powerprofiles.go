package powerprofiles

import (
	"github.com/godbus/dbus/v5"
)

type Profile struct {
	Profile string
	Driver  string
}

type Profiles struct {
	obj dbus.BusObject
}

func NewProfiles() (*Profiles, error) {

	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	powerProfiles := conn.Object("net.hadess.PowerProfiles", "/net/hadess/PowerProfiles")

	return &Profiles{
		obj: powerProfiles,
	}, nil
}

func (p *Profiles) GetProfiles() ([]Profile, error) {
	profiles := []Profile{}
	pp, err := p.obj.GetProperty("net.hadess.PowerProfiles.Profiles")
	if err != nil {
		return nil, err
	}
	ppp := []map[string]dbus.Variant{}
	err = pp.Store(&ppp)
	if err != nil {
		return nil, err
	}

	for _, profile := range ppp {
		profiles = append(profiles, Profile{
			Profile: profile["Profile"].Value().(string),
			Driver:  profile["Driver"].Value().(string),
		})
	}

	return profiles, nil
}

func (p *Profiles) SetProfile(profile string) error {
	return p.obj.Call("org.freedesktop.DBus.Properties.Set", 0, "net.hadess.PowerProfiles", "ActiveProfile", dbus.MakeVariant(profile)).Err
	//return p.obj.SetProperty("net.hadess.PowerProfiles.ActiveProfile", profile)
}
func (p *Profiles) GetActiveProfile() (string, error) {
	pp, err := p.obj.GetProperty("net.hadess.PowerProfiles.ActiveProfile")
	if err != nil {
		return "", err
	}

	return pp.Value().(string), nil
}
