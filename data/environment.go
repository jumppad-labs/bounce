package data

type User struct {
	UID      int      `json:"uid"`
	GID      int      `json:"gid"`
	Username string   `json:"username"`
	Homedir  string   `json:"homedir"`
	Groups   []string `json:"groups"`
}

type Host struct {
	Hostname string            `json:"hostname"`
	Network  map[string]string `json:"network"`
}

type Environment struct {
	Variables map[string]string `json:"variables"`
	User      User              `json:"user"`
	Host      Host              `json:"host"`
}
