#!/bin/bash

repoFolder=$(mktemp -d)

bareRepoFolder=$(mktemp -d)
bareRepo="${bareRepoFolder}/alreadyInSync.git"

# Init remote repo
cd "${bareRepoFolder}"
git init --bare "${bareRepo}"

# Init test repo
cd "${repoFolder}"
git init
git remote add origin "${bareRepo}"

touch README.md
git add README.md
git commit -m "Init"

git push origin main


echo "Repo folder: ${repoFolder}"
echo "Remote folder: ${bareRepoFolder}"
