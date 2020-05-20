# check which remote branch a local branch is tracking

### `git status -sb`

a branch that is tracking remote branch

```
~/path @ docker $ git status -sb
## docker...origin/docker
```

a branch that is not

```
~/path @ refactor $ git status -sb
## refactor
```

### `git remote show origin`

```
$ git remote show origin
* remote origin
  Fetch URL: https://github.com/user/repo.git
  Push  URL: https://github.com/user/repo.git
  HEAD branch: master
  Remote branches:=
    docker                                   tracked
    master                                   tracked
    refs/remotes/origin/deleted-branch     stale (use 'git remote prune' to remove)
  Local branches configured for 'git pull':
    docker merges with remote docker
    master merges with remote master
  Local refs configured for 'git push':
    docker pushes to docker (up to date)
    master pushes to master (up to date)
```
