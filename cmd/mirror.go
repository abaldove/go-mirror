package cmd

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var commandMirrorOpts = struct {
	url string
}{}

func init() {
	mirrorCmd.Flags().StringVar(&commandMirrorOpts.url, "url", "", "url=http://go-colly.org/ url to be crawled. Also, you can inform a an url with a subdirectory e.g. http://go-colly.org/articles/ ")
	RootCmd.AddCommand(mirrorCmd)
}

var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Command used to clone a url, like wget --mirror",
	RunE:  mirror,
}

func mirror(cmd *cobra.Command, args []string) error {
	baseUrl, err := url.ParseRequestURI(commandMirrorOpts.url)
	if err != nil {
		log.Errorf("Invalid URL %s", commandMirrorOpts.url)
		return err
	}
	basePath := baseUrl.Host
	allowedDomaiPattern := baseUrl.Host
	os.Mkdir(baseUrl.Host, os.ModePerm)
	if baseUrl.Path != "/" {
		allowedDomaiPattern = fmt.Sprintf(`.*%s%s.*`, baseUrl.Host, strings.ReplaceAll(baseUrl.Path, "/", "\\/"))
	} else {
		allowedDomaiPattern = fmt.Sprintf(`.*%s\/.*`, baseUrl.Host)
	}
	log.Info(allowedDomaiPattern)
	c := colly.NewCollector(
		colly.AllowedDomains(baseUrl.Host, fmt.Sprintf("www.%s", baseUrl.Host)),
		colly.URLFilters(
			regexp.MustCompile(allowedDomaiPattern),
			regexp.MustCompile(fmt.Sprintf(`.*%s\/.*jpg`, baseUrl.Host)),
			regexp.MustCompile(fmt.Sprintf(`.*%s\/.*jpeg`, baseUrl.Host)),
			regexp.MustCompile(fmt.Sprintf(`.*%s\/.*png`, baseUrl.Host)),
			regexp.MustCompile(fmt.Sprintf(`.*%s\/.*js`, baseUrl.Host)),
		),
		colly.CacheDir(fmt.Sprintf("./%s/cache", basePath)),
	)

	c.OnHTML("html", func(e *colly.HTMLElement) {
		os.MkdirAll(fmt.Sprintf("./%s/%s", basePath, e.Request.URL.Path), os.ModePerm)
		e.Response.Save(fmt.Sprintf("./%s/%sindex.html", basePath, e.Request.URL.Path))
		e.ForEach("a[href]", func(i int, h *colly.HTMLElement) {
			c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
		})
		e.ForEach("link", func(i int, h *colly.HTMLElement) {
			if h.Attr("type") == "text/css" || h.Attr("type") == "application/javascript" {
				u, _ := url.Parse(h.Request.AbsoluteURL(h.Attr("href")))
				if !u.IsAbs() {
					c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
				}
			}
		})
		e.ForEach("img", func(i int, h *colly.HTMLElement) {
			u, _ := url.Parse(h.Attr("src"))
			if !u.IsAbs() {
				c.Visit(h.Request.AbsoluteURL(h.Attr("src")))
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting URL %s", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
			err := r.Save(fmt.Sprintf("./%s/%s", basePath, r.FileName()))
			if err != nil {
				log.Error(err)
			}
			return
		}
	})

	c.Visit(baseUrl.String())

	return nil
}
