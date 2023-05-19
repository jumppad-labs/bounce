package cmd

import (
	"encoding/json"
	"net"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/jumppad-labs/bounce/data"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:  "env",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		environment := data.Environment{}

		e := getEnvironment()
		environment.Variables = e

		u, err := getUser()
		if err != nil {
			return err
		}
		environment.User = u

		h, err := getHost()
		if err != nil {
			return err
		}
		environment.Host = h

		result, err := json.MarshalIndent(environment, "", " ")
		if err != nil {
			return err
		}
		pretty.Println(string(result))

		return nil
	},
}

func getHost() (data.Host, error) {
	result := data.Host{
		Network: map[string]string{},
	}

	hostname, err := os.Hostname()
	if err != nil {
		return result, err
	}

	result.Hostname = hostname

	ifaces, err := net.Interfaces()
	if err != nil {
		return result, err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return result, err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			result.Network[i.Name] = ip.String()
		}
	}

	return result, nil
}

func getUser() (data.User, error) {
	result := data.User{}

	u, err := user.Current()
	if err != nil {
		return result, err
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return result, err
	}

	result.UID = uid
	result.Username = u.Username
	result.Homedir = u.HomeDir

	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return result, err
	}

	result.GID = gid

	groupIDs, err := u.GroupIds()
	if err != nil {
		return result, err
	}

	for _, id := range groupIDs {
		group, err := user.LookupGroupId(id)
		if err != nil {
			return result, err
		}

		result.Groups = append(result.Groups, group.Name)
	}

	return result, nil
}

func getEnvironment() map[string]string {
	values := map[string]string{}

	for _, entry := range os.Environ() {
		key, value, _ := strings.Cut(entry, "=")
		values[key] = value
	}

	return values
}
