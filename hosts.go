package AnotherSteamCommunityFix

import (
	"github.com/zyfworks/goodhosts"
	"errors"
)

func AddHosts(ip string, hostEntries... string) error {
	hosts, err := goodhosts.NewHosts()
	if err != nil {
		return err
	}

	if !hosts.IsWritable() {
		return errors.New("host file not writable, try running with elevated privileges")
	}

	if err = hosts.Add(ip, hostEntries...); err != nil {
		return err
	}

	return hosts.Flush()
}

func RemoveHosts(ip string, hostEntries... string) error {
	hosts, err := goodhosts.NewHosts()
	if err != nil {
		return err
	}

	if !hosts.IsWritable() {
		return errors.New("host file not writable, try running with elevated privileges")
	}

	if err = hosts.Remove(ip, hostEntries...); err != nil {
		return err
	}

	return hosts.Flush()
}