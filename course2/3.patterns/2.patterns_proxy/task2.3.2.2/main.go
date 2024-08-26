package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error) // opt := github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
}

type Item struct {
	Title       string
	Description string
	Link        string
}

type GithubAdapter struct {
	RepoList RepoLister
	GistList GistLister
}

func NewGithubAdapter(githubClient *github.Client) *GithubAdapter {
	g := &GithubAdapter{
		RepoList: githubClient.Repositories,
		GistList: githubClient.Gists,
	}
	return g
}

func (g *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := g.GistList.List(ctx, username, &github.GistListOptions{})
	if err != nil {
		return nil, err
	}

	res := make([]Item, 0, len(gists))

	for _, gist := range gists {
		item := Item{
			Title:       fmt.Sprintf("Gist %s", gist.GetID()),
			Description: gist.GetDescription(),
			Link:        gist.GetHTMLURL(),
		}
		res = append(res, item)
	}
	return res, nil
}

func (g *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := g.RepoList.List(ctx, username, &github.RepositoryListOptions{})

	if err != nil {
		return nil, err
	}

	res := make([]Item, 0, len(repos))

	for _, repo := range repos {
		item := Item{
			Title:       repo.GetName(),
			Description: repo.GetDescription(),
			Link:        repo.GetHTMLURL(),
		}
		res = append(res, item)
	}
	return res, nil
}

type GithubProxy struct {
	github Githuber
	cache  map[string][]Item
}

func NewGithubProxy(githubClient *github.Client) *GithubProxy {
	return &GithubProxy{
		github: NewGithubAdapter(githubClient),
		cache:  make(map[string][]Item),
	}
}

func (g *GithubProxy) GetGists(ctx context.Context, username string) ([]Item, error) {
	if gists, ok := g.cache[fmt.Sprintf("%s's gists", username)]; ok {
		return gists, nil
	}

	gists, err := g.github.GetGists(ctx, username)
	if err != nil {
		return nil, err
	}

	g.cache[fmt.Sprintf("%s's gists", username)] = gists

	return gists, nil
}

func (g *GithubProxy) GetRepos(ctx context.Context, username string) ([]Item, error) {
	if repos, ok := g.cache[fmt.Sprintf("%s's repos", username)]; ok {
		return repos, nil
	}

	repos, err := g.github.GetRepos(ctx, username)
	if err != nil {
		return nil, err
	}

	g.cache[fmt.Sprintf("%s's repos", username)] = repos

	return repos, nil
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	g := NewGithubProxy(client)
	fmt.Println(g.GetGists(context.Background(), "Artenso"))
	starrt := time.Now()
	fmt.Println(g.GetRepos(context.Background(), "Artenso"))
	fmt.Println(time.Since(starrt)) // 267.198518ms
	starrt = time.Now()
	fmt.Println(g.GetRepos(context.Background(), "Artenso"))
	fmt.Println(time.Since(starrt)) // 63.162µs
}
