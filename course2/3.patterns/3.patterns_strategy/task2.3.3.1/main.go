package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

type Item struct {
	Title       string
	Description string
	Link        string
}
type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}
type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

type GithubGist struct {
	client *github.Client
}

func NewGithubGist(githubClient *github.Client) GithubLister {
	return &GithubGist{
		client: githubClient,
	}
}

func (g *GithubGist) GetItems(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := g.client.Gists.List(ctx, username, &github.GistListOptions{})
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

type GithubRepo struct {
	client *github.Client
}

func NewGithubRepo(githubClient *github.Client) GithubLister {
	return &GithubRepo{
		client: githubClient,
	}
}

func (g *GithubRepo) GetItems(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := g.client.Repositories.List(ctx, username, &github.RepositoryListOptions{})

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

type GeneralGithub struct {
	client *github.Client
}

func NewGeneralGithub(githubClient *github.Client) GeneralGithubLister {
	return &GeneralGithub{
		client: githubClient,
	}
}

func (g *GeneralGithub) GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error) {
	return strategy.GetItems(ctx, username)
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	gist := NewGithubGist(client)
	repo := NewGithubRepo(client)
	gg := NewGeneralGithub(client)
	data, err := gg.GetItems(context.Background(), "Artenso", gist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
	data, err = gg.GetItems(context.Background(), "Artenso", repo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
