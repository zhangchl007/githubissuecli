package github

import "time"

const (
	IssuesURL     = "https://api.github.com/repos/"
	ReposURL      = "https://api.github.com/user/repos"
	UserReposURL  = "https://api.github.com/repos/"
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

type Issues struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	Locked    bool
	Assignees *[]Assignees
	CreateAt  time.Time `json:"created_at"`
	Body      string
}

type Assignees struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issueyamlfile struct {
	Title     string   `yaml:"title"`
	Body      string   `yaml:"body"`
	Assignees []string `yaml:"assignees"`
	State     string   `yaml:"state"`
	Locked    bool     `yaml:"locked"`
	Labels    []string `yaml:"labels"`
}

type Userinfo struct {
	Userid string `yaml:"userid"`
	Token  string `yaml:"token"`
	Repo   string `yaml:"repo"`
}

type Repos struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Private     bool   `json:"private"`
	Owner       Owner
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Fork        bool
	Teamurl     string    `json:"team_url"`
	CreateAt    time.Time `json:"created_at"`
	Forks       int       `json:"forks"`
	Openissues  int       `json:"open_issues"`
	Language    string    `json:"language"`
}

type Owner struct {
	Login string `json:"login"`
}
