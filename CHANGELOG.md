# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),

## [Unreleased]

## [release-branch.r3] - 2022-03-25
### Added

### Changed

### Removed

## [release-branch.r2] - 2022-03-12
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

## [release-branch.r1] - 2022-02-25
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