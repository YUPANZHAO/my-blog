package util

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type ArticleInfo struct {
	Path      string
	Name      string
	Category  string
	Date      string
	Cover_url string
	Title     string
	Introduce []string
	Content   string
}

var (
	lastHash    plumbing.Hash
	articleList []ArticleInfo
	articleMap  map[string]*ArticleInfo
)

func StartPullThread() {
	// parameters
	url := "https://github.com/YUPANZHAO/article_repo.git"
	dir := "./article_repo"
	// clone repository
	clone(url, dir)
	// loop every one minute
	for {
		// pull repository
		pullArticles(dir)
		// sleep one minute
		time.Sleep(time.Minute)
	}
}

func GetArticles() *[]ArticleInfo {
	return &articleList
}

func GetArticle(name string) (ArticleInfo, error) {
	if articleMap[name] == nil {
		return ArticleInfo{}, errors.New("NO FOUND")
	}
	return *articleMap[name], nil
}

func clone(url string, dir string) {
	// clone repository
	git.PlainClone(dir, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
}

func pullArticles(path string) {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, _ := git.PlainOpen(path)
	// Get the working directory for the repository
	w, _ := r.Worktree()
	// pull repository
	w.Pull(&git.PullOptions{RemoteName: "origin"})
	// get head commit hash value
	ref, _ := r.Head()
	commit, _ := r.CommitObject(ref.Hash())
	// check if the hash value has changed
	if commit.Hash != lastHash {
		updateList(path)
		lastHash = commit.Hash
	}
}

func updateList(path string) {
	// loop mardown file
	var articles []ArticleInfo
	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == ".md" {
			articleInfo, err := parseArticle(path)
			if err == nil {
				articles = append(articles, articleInfo)
			}
		}
		return nil
	})
	// override old article list
	articleList = articles
	// sort
	sort.Slice(articleList, func(i, j int) bool {
		return articleList[i].Date > articleList[i].Date
	})
	// map
	articleMap = make(map[string]*ArticleInfo)
	for idx, item := range articleList {
		articleMap[item.Name] = &articleList[idx]
	}
}

func parseArticle(path string) (ArticleInfo, error) {
	// split path by character '/'
	words := strings.Split(path, "/")
	// when path no conform 'repo_name/category/filename.md'
	if len(words) != 3 {
		return ArticleInfo{}, errors.New("WRONG FILE")
	}
	// read file
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return ArticleInfo{}, errors.New("WRONG FILE")
	}
	// cast to string
	content := string(buf)
	// parse cover_url
	reg, _ := regexp.Compile(`!\[cover\](.*)`)
	cover_url := reg.FindString(content)
	cover_url = strings.TrimPrefix(cover_url, "![cover](")
	cover_url = strings.TrimSuffix(cover_url, ")")
	// parse title
	reg, _ = regexp.Compile(`# .*\n`)
	title := reg.FindString(content)
	title = strings.TrimPrefix(title, "# ")
	title = strings.TrimSuffix(title, "\n")
	// parse date
	reg, _ = regexp.Compile(`\*\*\d{4}-\d{2}-\d{2}\*\*`)
	date := reg.FindString(content)
	date = strings.Trim(date, "*")
	// parse introduce
	reg, _ = regexp.Compile(`## 简介((.|\n)+?)(\n{3})`)
	introduce := reg.FindString(content)
	introduce = strings.TrimPrefix(introduce, "## 简介\n\n")
	introduce = strings.TrimSuffix(introduce, "\n\n\n")
	// parse content
	reg, _ = regexp.Compile(`## 简介(.|\n)*`)
	content = reg.FindString(content)
	// return article info
	return ArticleInfo{
		Path:      path,
		Name:      words[2],
		Category:  words[1],
		Date:      date,
		Cover_url: cover_url,
		Title:     title,
		Introduce: strings.Split(introduce, "\n\n"),
		Content:   content,
	}, nil
}
