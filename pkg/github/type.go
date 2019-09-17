package github
import "time"
const (
        IssuesURL = "https://api.github.com/repos/"
        MethodGet     = "GET"
        MethodHead    = "HEAD"
        MethodPost    = "POST"
        MethodPut     = "PUT"
        MethodPatch   = "PATCH" // RFC 5789
        MethodDelete  = "DELETE"
        MethodConnect = "CONNECT"
        MethodOptions = "OPTIONS"
        MethodTrace   = "TRACE"
//        PersonalAccessToken = "16dbcfdffd69e12cd5c9d0236b51fb92b3172a36"
//        Userid        = "zhangchl007"
//        Repo          = "test"
)

type Issues struct {
    Number int
    HTMLURL string `json:"html_url"`
    Title string
    State string
    Locked bool
    Assignees *[]Assignees
    CreateAt time.Time `json:"created_at"`
    Body string
}

type Assignees struct {
    Login string
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
	Userid   string   `yaml:"userid"`
	Token    string   `yaml:"token"`
	Repo     string   `yaml:"repo"`
}
