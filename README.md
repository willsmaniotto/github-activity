# github-activity

A CLI tool to fetch and display the recent public activity of a GitHub user.

Built as part of the [roadmap.sh GitHub User Activity project](https://roadmap.sh/projects/github-user-activity).

## Requirements

- Go 1.23+

## Usage

```bash
go run main.go <github-username>
```

**Example:**

```bash
go run main.go torvalds
```

## Output

The tool prints recent events from the GitHub API, including:

- **PushEvent** — commits pushed to a repository
- **PullRequestEvent** — pull requests opened, closed, or merged
- **CreateEvent** — branches or tags created
- **WatchEvent** — repositories starred
- **IssuesEvent** — issues opened, closed, or edited
- **IssueCommentEvent** — comments added to issues
- Other events are shown with their raw event type

**Example output:**

```
Pushed to torvalds/linux at 2024-01-15T10:30:00Z
  - net: fix null pointer dereference
Created branch in torvalds/linux at 2024-01-14T09:00:00Z
Started to watch golang/go at 2024-01-13T08:00:00Z
```

## How It Works

The tool calls the [GitHub Events API](https://docs.github.com/en/rest/activity/events) (`GET /users/{username}/events`) and formats the response into a human-readable summary. No authentication is required; only public events are shown.
