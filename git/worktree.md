```.sh
# download the master branch code at ../aneatrepo-master
~/aneatrepo @ dev $ git worktree add ../aneatrepo-master master

# IDE makes code reading easier
~/aneatrepo @ dev $ code ../aneatrepo-master

# you can also switch to other commits (branches), if you want
~/aneatrepo-master @ master $ git checkout staging
~/aneatrepo-master @ staging $

# ... unless it's already checked out somewhere else
~/aneatrepo-master @ master $ git checkout dev
fatal: 'dev' is already checked out at '/Users/xxx/aneatrepo'

# `worktree remove` removes the entire directory
~/aneatrepo @ dev $ git worktree remove master
~/aneatrepo @ dev $ ll ../aneatrepo-master
ls: ../aneatrepo-master: No such file or directory
```
