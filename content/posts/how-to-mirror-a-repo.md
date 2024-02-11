+++
title = 'How to mirror a git repository'
date = 2023-08-07
draft = false
+++

Well, It’s not actually mirror. It’s just “How to push to multiple repos but only fetch changes from one”.

Initialize a repository as always:

```bash
git init
git add README.md
git commit -m "add readme.md"
git branch -M main
git remote add origin $MAIN_REPO_URL
git push -u origin main
```

Add “only push” remotes:

```bash
git remote set-url --add --push origin $MAIN_REPO_URL
git remote set-url --add --push origin $MIRROR_REPO_URL
```