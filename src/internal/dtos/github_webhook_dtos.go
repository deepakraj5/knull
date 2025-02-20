package dtos

type GithubWebhookRequestHeaders struct {
	ContentType  string `json:"Content-Type"`
	UserAgent    string `json:"User-Agent"`
	XGitHubEvent string `json:"X-GitHub-Event"`
}

type GithubWebhookRequestBody struct {
	Ref        string `json:"ref"`
	Repository struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		URL      string `json:"url"`
		CloneUrl string `json:"clone_url"`
	} `json:"repository"`
}
