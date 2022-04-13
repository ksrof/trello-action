# 🤖 Trello Action
Easily automate the process of creating, moving and deleting Trello Cards, harness the power of Github Actions to perform requests on the [Trello API](https://developer.atlassian.com/cloud/trello/guides/rest-api/api-introduction/) when `issues` or `pull requests` get `opened`, `reopened`, `labeled` or `closed`!

## Creating a card
Cards can be created either when an `issue` or a `pull request` gets `opened` or `reopened`. It uses their `title` and their `htmlURL` to fill both `name` and `url` in the Trello Card.

```yml
on:
 issues:
  types: [opened, reopened]
 pull_request:
  types: [opened, reopened]

env:
 TRELLO_KEY: ${{ secrets.TRELLO_KEY }}
 TRELLO_TOKEN: ${{ secrets.TRELLO_TOKEN }}
 TRELLO_ID_BOARD: ${{ secrets.TRELLO_ID_BOARD }}
 TRELLO_ID_LIST: ${{ secrets.TRELLO_ID_LIST }}
 GH_TOKEN: ${{ secrets.GH_TOKEN }}
 GH_USER: username
 GH_REPO: repository

jobs:

 issue-card:
  if: github.event_name == 'issues'
  runs-on: ubuntu-latest
  steps:
   - uses: actions/checkout@v3
   - uses: actions/setup-go@v3
     with:
      go-version: 1.18
   - run: go run cmd/create/main.go
     env:
      GH_EVENT: ${{ github.event_name }}
      GH_ID: ${{ github.event.issue.number }}

 pull-card:
  if: github.event_name == 'pull_request'
  runs-on: ubuntu-latest
  steps:
   - uses: actions/checkout@v3
   - uses: actions/setup-go@v3
     with:
      go-version: 1.18
   - run: go run cmd/create/main.go
     env:
      GH_EVENT: ${{ github.event_name }}
      GH_ID: ${{ github.event.pull_request.number }}
```

In order for the workflow to run appropriately, you need to add the following environment variables to your actions secrets: `TRELLO_KEY`, `TRELLO_TOKEN`, `TRELLO_ID_BOARD`, `TRELLO_ID_LIST` and `GH_TOKEN`.

## Deleting a card
Cards can be deleted either when an `issue` or a `pull request` gets `closed`. No matter in what list the card is, it will get deleted.

```yml
on:
 issues:
  types: [closed]
 pull_request:
  types: [closed]

env:
 TRELLO_KEY: ${{ secrets.TRELLO_KEY }}
 TRELLO_TOKEN: ${{ secrets.TRELLO_TOKEN }}
 TRELLO_ID_BOARD: ${{ secrets.TRELLO_ID_BOARD }}
 TRELLO_ID_LIST: ${{ secrets.TRELLO_ID_LIST }}
 GH_TOKEN: ${{ secrets.GH_TOKEN }}
 GH_USER: username
 GH_REPO: repository

jobs:

 issue-card:
  if: github.event_name == 'issues'
  runs-on: ubuntu-latest
  steps:
   - uses: actions/checkout@v3
   - uses: actions/setup-go@v3
     with:
      go-version: 1.18
   - run: go run cmd/delete/main.go
     env:
      GH_EVENT: ${{ github.event_name }}
      GH_ID: ${{ github.event.issue.number }}

 pull-card:
  if: github.event_name == 'pull_request'
  runs-on: ubuntu-latest
  steps:
   - uses: actions/checkout@v3
   - uses: actions/setup-go@v3
     with:
      go-version: 1.18
   - run: go run cmd/delete/main.go
     env:
      GH_EVENT: ${{ github.event_name }}
      GH_ID: ${{ github.event.pull_request.number }}
```

## Updating a card
Cards can be updated either when an `issue` or a `pull request` gets `labeled`. To be able to move cards between lists on your Trello Board each label that you add must be equal to the name of the list where you want the card to move to. Delete the previous label/s before jumping lists again.

```yml
on:
 issues:
  types: [labeled]
 pull_request:
  types: [labeled]

env:
 TRELLO_KEY: ${{ secrets.TRELLO_KEY }}
 TRELLO_TOKEN: ${{ secrets.TRELLO_TOKEN }}
 TRELLO_ID_BOARD: ${{ secrets.TRELLO_ID_BOARD }}
 TRELLO_ID_LIST: ${{ secrets.TRELLO_ID_LIST }}
 GH_TOKEN: ${{ secrets.GH_TOKEN }}
 GH_USER: username
 GH_REPO: repository

jobs:

 issue-card:
  if: github.event_name == 'issues'
  runs-on: ubuntu-latest
  steps:
   - uses: actions/checkout@v3
   - uses: actions/setup-go@v3
     with:
      go-version: 1.18
   - run: go run cmd/update/main.go
     env:
      GH_EVENT: ${{ github.event_name }}
      GH_ID: ${{ github.event.issue.number }}

 pull-card:
  if: github.event_name == 'pull_request'
  runs-on: ubuntu-latest
  steps:
   - uses: actions/checkout@v3
   - uses: actions/setup-go@v3
     with:
      go-version: 1.18
   - run: go run cmd/update/main.go
     env:
      GH_EVENT: ${{ github.event_name }}
      GH_ID: ${{ github.event.pull_request.number }}
```

## Contact
Whether you have any feedback, suggestion or question please don't doubt and feel free to contact me, I'll be glad to answer!

- [Email](mailto:kevinsunercontacto@gmail.com)
- [Medium](https://medium.com/@ksrof)
- [Twitter](https://twitter.com/itsksrof)

## Changelog
Check out [`CHANGELOG.md`](https://github.com/ksrof/trello-action/blob/main/CHANGELOG.md) to see details about past and recent releases!

## License
The MIT License (MIT) - see [`license`](https://github.com/ksrof/trello-action/blob/main/LICENSE) for more details.

