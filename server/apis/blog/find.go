package blog

type FindBlogRequest struct {
}

type Blog struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Writer  string   `json:"writer"`
	View    int      `json:"view"`
	Heart   int      `json:"heart"`
	Tags    []string `json:"tags"`
}

type FindBlogResponse struct {
	List  []Blog `json:"list"`
	Total int    `json:"total"`
}
