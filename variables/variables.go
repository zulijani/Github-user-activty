package variables

var Github_url string

type Info struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		Id           int    `json:"id"`
		DisplayLogin string `json:"display_login"`
	}
	Repo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	Payload struct {
		RepoId int    `json:"repository_id"`
		Action string `json:"action"`
		Issue  struct {
			Id int `json:"id"`
		} `json:"issue"`
		PULLRequest struct {
			Id int `json:"id"`
		} `json:"pull_request"`
		Commits []struct {
			Sha     string `json:"sha"`
			Message string `json:"message"`
			URL     string `json:"url"`
		} `json:"commits"`
	}
}