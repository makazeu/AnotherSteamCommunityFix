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
		return errors.New("hosts文件不可写，请以更高的权限运行")
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
		return errors.New("hosts文件不可写，请以更高的权限运行")
	}

	if err = hosts.Remove(ip, hostEntries...); err != nil {
		return err
	}

	return hosts.Flush()
}