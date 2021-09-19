package article

import (
	"time"
)

type (
	Article struct {
		Source      map[string]string `json:"source"`
		Author      string            `json:"author"`
		Title       string            `json:"title"`
		Description string            `json:"description"`
		Url         string            `json:"url"`
		Content     string            `json:"content"`
		PublishedAt time.Time         `json:"publishedAt"`
	}
)
