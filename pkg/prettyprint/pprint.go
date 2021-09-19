package pprint

import (
	"fmt"

	"github.com/gregves/newsapi-golang-client/pkg/article"
	"github.com/pterm/pterm"
)

func PrettyPrint(articles []article.Article) {
	headerText := fmt.Sprintf("Found %d articles", len(articles))

	pterm.DefaultHeader.
		WithFullWidth().
		WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).
		WithTextStyle(pterm.NewStyle(pterm.FgBlack)).
		Println(headerText)

	for _, article := range articles {
		fmt.Println()
		pterm.FgCyan.Println(article.Title)
		fmt.Println()
		pterm.DefaultBasicText.Println(article.Description)
		pterm.FgLightBlue.Println(article.Url)
		fmt.Println()
		pterm.FgMagenta.Println(article.PublishedAt)
		fmt.Println()
		//time.Sleep(500 * time.Millisecond)
	}
}
