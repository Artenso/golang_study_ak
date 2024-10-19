package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-github/v53/github"
)

type MockRepoLister struct{}

type MockGistLister struct{}

func (m *MockRepoLister) List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	res := make([]*github.Repository, 4)
	for i := range res {
		testName := fmt.Sprintf("TestRepo%d", i)
		testDescription := fmt.Sprintf("test description of repo %d", i)
		testURL := fmt.Sprintf("https://github.com/%s/%s", username, testName)
		res[i] = &github.Repository{
			Name:        &testName,
			Description: &testDescription,
			HTMLURL:     &testURL,
		}
	}
	return res, nil, nil
}

func (m *MockGistLister) List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	res := make([]*github.Gist, 4)
	for i := range res {
		testID := fmt.Sprintf("%d", i)
		testDescription := fmt.Sprintf("test description of gist %d", i)
		testURL := fmt.Sprintf("https://gist.github.com/%s/%s", username, testID)
		res[i] = &github.Gist{
			ID:          &testID,
			Description: &testDescription,
			HTMLURL:     &testURL,
		}
	}
	return res, nil, nil
}

func TestGithubAdapter_GetGists(t *testing.T) {
	type fields struct {
		RepoList RepoLister
		GistList GistLister
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Item
		wantErr bool
	}{
		{
			name: "#1",
			fields: fields{
				RepoList: &MockRepoLister{},
				GistList: &MockGistLister{},
			},
			args: args{
				ctx:      context.Background(),
				username: "testUser",
			},
			want: []Item{
				{
					Title:       "Gist 0",
					Description: "test description of gist 0",
					Link:        "https://gist.github.com/testUser/0",
				},
				{
					Title:       "Gist 1",
					Description: "test description of gist 1",
					Link:        "https://gist.github.com/testUser/1",
				},
				{
					Title:       "Gist 2",
					Description: "test description of gist 2",
					Link:        "https://gist.github.com/testUser/2",
				},
				{
					Title:       "Gist 3",
					Description: "test description of gist 3",
					Link:        "https://gist.github.com/testUser/3",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GithubAdapter{
				RepoList: tt.fields.RepoList,
				GistList: tt.fields.GistList,
			}
			got, err := g.GetGists(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GithubAdapter.GetGists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GithubAdapter.GetGists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGithubAdapter_GetRepos(t *testing.T) {
	type fields struct {
		RepoList RepoLister
		GistList GistLister
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Item
		wantErr bool
	}{
		{
			name: "#1",
			fields: fields{
				RepoList: &MockRepoLister{},
				GistList: &MockGistLister{},
			},
			args: args{
				ctx:      context.Background(),
				username: "testUser",
			},
			want: []Item{
				{
					Title:       "TestRepo0",
					Description: "test description of repo 0",
					Link:        "https://github.com/testUser/TestRepo0",
				},
				{
					Title:       "TestRepo1",
					Description: "test description of repo 1",
					Link:        "https://github.com/testUser/TestRepo1",
				},
				{
					Title:       "TestRepo2",
					Description: "test description of repo 2",
					Link:        "https://github.com/testUser/TestRepo2",
				},
				{
					Title:       "TestRepo3",
					Description: "test description of repo 3",
					Link:        "https://github.com/testUser/TestRepo3",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GithubAdapter{
				RepoList: tt.fields.RepoList,
				GistList: tt.fields.GistList,
			}
			got, err := g.GetRepos(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GithubAdapter.GetRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GithubAdapter.GetRepos() = %v,\nwant %v", got, tt.want)
			}
		})
	}
}
