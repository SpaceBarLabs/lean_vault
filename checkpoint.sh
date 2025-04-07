#! /bin/bash

# use args to get the commit message
# default to "checkpoint" if no message is provided
commit_message=${1:-"checkpoint"}

git add .
git commit -m "$commit_message"
git push origin main

