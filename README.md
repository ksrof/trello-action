# Github Actions + Trello
🤖 Easily automate the process of creating a new Trello card, use the power of Github Actions to create them when an Issue or a Pull Request is made!

## Functionalities
- [x] Create Card when creating an Issue or Pull Request
- [ ] Delete Card when closing an Issue or Pull Request

## Example
This shows an example Github Actions workflow for adding a new Trello Card to a specific list when creating an Issue or Pull Request.
```yaml
name: Create Trello Card on Issue or PR

on:
  issues:
    types: [opened]
  pull_request:
    types: [opened]

jobs:
  create-issue-card:
    runs-on: ubuntu-latest
    if: ${{ github.event_name == 'issues' }}
    steps:
      - uses: actions/checkout@v2

      - name: Setup
        uses: actions/setup-go@v2
        with:
          go-version: 1.17

      - name: Build
        run: go build -v ./...

      - name: Run
        env:
          TRELLO_KEY: ${{ secrets.TRELLO_KEY }}
          TRELLO_TOKEN: ${{ secrets.TRELLO_TOKEN }}
          TRELLO_ID_LIST: ${{ secrets.TRELLO_ID_LIST }}
          GH_TOKEN: ${{ secrets.GH_TOKEN }}
          GH_USER: ksrof
          GH_REPO: gha-trello
          ACTION: issue
        run: go run main.go

  create-pull-card:
    runs-on: ubuntu-latest
    if: ${{ github.event_name == 'pull_request' }}
    steps:
      - uses: actions/checkout@v2

      - name: Setup
        uses: actions/setup-go@v2
        with:
          go-version: 1.17

      - name: Build
        run: go build -v ./...

      - name: Run
        env:
          TRELLO_KEY: ${{ secrets.TRELLO_KEY }}
          TRELLO_TOKEN: ${{ secrets.TRELLO_TOKEN }}
          TRELLO_ID_LIST: ${{ secrets.TRELLO_ID_LIST }}
          GH_TOKEN: ${{ secrets.GH_TOKEN }}
          GH_USER: ksrof
          GH_REPO: gha-trello
          ACTION: pull
        run: go run main.go
```
**Note:** The workflow is a work in progress, a lot of improvements will be made!

## License
The MIT License (MIT) - see [`license`](https://github.com/ksrof/gha-trello/blob/main/LICENSE) for more details.