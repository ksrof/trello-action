# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),

## [release.r3](https://github.com/ksrof/trello-action/releases/tag/release.r3) - 2022-04-13
### Added
- URL parser
- Regexp validator for Github Personal Access Token
- Get specific issue or pull_request by user, repo and id
- Create Trello Card based on an event type
- Create Trello Card with Title and URL from issue or pull_request
- Vendor dependencies
- Delete Trello Card based on an id
- Get all cards from a specific Trello Board
- Get all labels from a specific issue or pull_request
- Get all lists from a specific Trello Board
- Update Trello Card based on an event type
- Add action.yml
### Changed
- Renamed files, go module, variables and fixed typos
- Renamed go module
- Add examples to the README.md
### Removed
- Removed .github/workflows folder

## [release-branch.r2](https://github.com/ksrof/trello-action/releases/tag/release.r2) - 2022-03-12
### Added
- Create card from issue or pr
- Created different yml files for different actions
- Get issues and prs from specific user and repo
- Validate, parse and set auth header
- Load env variables based on the environment
- Set query parameters
- Build a new http request with a given method, url and payload
- Added data structure for environment variables
- Added data structure for query parameters
- Added CHANGELOG.md

### Changed
- Updated .gitignore
- Updated comments for godocs
- Updated go.mod and go.sum
- Update README.md
- Update CHANGELOG.md

### Removed
- Removed gha-trello.yaml
- Removed main.go
- Removed utils folder

## [release-branch.r1](https://github.com/ksrof/trello-action/releases/tag/pre-release.r1) - 2022-02-25
### Added
- Added .yaml and .env to .gitignore
- Added github.com/spf13/viper
- Get user from api.github.com
- Get issues and pull requests from a specific repo
- Added structs to map the issue and pr data
- Added a function to parse the github personal access token
- Added a SetEnv method to get env variables
- Added method to create trello cards from issues or pulls
- Added new sections to README.md
- Created gha-trello github action
- Added gha-trello workflow
- Added example and license

### Changed
- Renamed github environment variables
- Ignore env.yaml but no other .yaml files
- Set value to prod
- github.event.name conditional was wrong
- Action name was wrong
- Action type was wrong

### Removed
- Removed todo list at top
- Deleted workflow, it wast just a test