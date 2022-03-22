# Github Actions + Trello
🤖 Easily automate the process of creating, moving and deleting Trello Cards, use the power of Github Actions to perform requests on the Trello API when a new Issue or Pull Request gets made, updated, labeled or closed!

## Functionalities
- [x] Create card when creating an Issue or Pull Request
- [ ] Move card when updating an Issue or Pull Request
- [ ] Delete card when closing an Issue or Pull Request

## Examples
The following examples show how you can automate the process of creating a new Trello card when a new Issue or a new Pull Request is opened!

### Create Card From Issue
```yml
name: Create Card From Issue

on:
  issues:
    types: [opened]

jobs:
  create-issue-card:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Setup
        uses: actions/setup-go@v2
        with:
          go-version: 1.17
        
      - name: Run
        env:
          TRELLO_KEY: ${{ secrets.TRELLO_KEY }}
          TRELLO_TOKEN: ${{ secrets.TRELLO_TOKEN }}
          TRELLO_ID_LIST: ${{ secrets.TRELLO_ID_LIST }}
          GH_TOKEN: ${{ secrets.GH_TOKEN }}
          GH_USER: user
          GH_REPO: repo
          ACTION: issue
        run: go run cmd/create/main.go
```

### Create Card From PR
```yml
name: Create Card From PR

on:
  pull_request:
    types: [opened]

jobs:
  create-pull-card:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Setup
        uses: actions/setup-go@v2
        with:
          go-version: 1.17

      - name: Run
        env:
          TRELLO_KEY: ${{ secrets.TRELLO_KEY }}
          TRELLO_TOKEN: ${{ secrets.TRELLO_TOKEN }}
          TRELLO_ID_LIST: ${{ secrets.TRELLO_ID_LIST }}
          GH_TOKEN: ${{ secrets.GH_TOKEN }}
          GH_USER: user
          GH_REPO: repo
          ACTION: pull
        run: go run cmd/create/main.go
```

**Note:** This project is a work in progress, a lot of improvements will be made!

## Changelog
Check out [`CHANGELOG.md`](https://github.com/ksrof/gha-trello/blob/main/CHANGELOG.md) to see details about past and recent releases!

## License
The MIT License (MIT) - see [`license`](https://github.com/ksrof/gha-trello/blob/main/LICENSE) for more details.
