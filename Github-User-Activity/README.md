# GitHub User Activity CLI

A simple command-line tool to fetch and display recent GitHub user activity.

## Description

This CLI tool allows you to fetch the recent activity of any GitHub user and displays it in your terminal. It uses the GitHub API to retrieve user events and presents them in a readable format.

## Features

- Fetch recent GitHub user activity
- Display activity in a clean, readable format
- Error handling for invalid usernames and API failures
- No external dependencies for API calls (uses native Fetch API)

## Installation

1. Ensure you have Node.js installed
2. Clone this repository or download the script
3. Navigate to the project directory

## Usage

Run the script with a GitHub username as an argument:

```bash
node index.js <username>
```

Example:

```bash
node index.js kamranahmedse
```

### Output Format

The tool will display user activities in the following format:

```
[Event Type] in [Repository Name]
```

Example output:

```
Push in kamranahmedse/developer-roadmap
Issue in kamranahmedse/developer-roadmap
Watch in kamranahmedse/developer-roadmap
```

### Error Handling

The tool handles the following error cases:

- Invalid or non-existent GitHub username
- API request failures
- Empty activity lists

## API Endpoint

The tool uses the GitHub API endpoint:

```
https://api.github.com/users/<username>/events
```

## Code Structure

The script consists of two main functions:

**`fetchGitHubUser(username)`** – Fetches user activity data from GitHub API  
- Takes a GitHub username as input  
- Returns a Promise that resolves to the activity data  
- Throws errors for invalid users or API failures  

**`printGitHubRepos(events)`** – Formats and displays the activity data  
- Takes an array of event objects  
- Prints formatted output to console  
- Handles empty activity lists

