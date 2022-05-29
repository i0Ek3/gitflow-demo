# gitflow-demo

Git Flow demo.

## Build

```Shell
$ go run main.go

# Or

$ go build . -v
```
## For Forking Mode

- First, fork this repo and then clone it, setting remote
    - fork original repo usera/gitflow-demo to your own
    - clone your repo userb/gitflow-demo
    - cd gitflow-demo
    - git remote add upstream usera/gitflow-demo
    - git remote set-url --push upstream no_push
- Sync local repo and upstream
    - git fetch upstream
    - git checkout main
    - git rebase upstream/main
- Create new feature branch to develop
    - git checkout -b feature/add-feature_name
    - develop
- Commit your code
    - git fetch upstream
    - git rebase upstream/main
    - git add .
    - git commit
    - also you can `git rebase -i origin/main` to merge multi commits
- Push your commit
    - git push -f origin feature/add-feature_name
- Go to userb's GitHub page to create PR
- Wait to merge by owner
