nestguard
=========

Small utility to check the status of nested git repositories.

Currently it gives warnings for:

- Nested git repositories with uncommitted changes
- Nested editable github repositories which have a different local revision than in the pip `requirements.txt` in the root of the outer repository
- Different editable github repository revisions in `requirements.txt` and `requirements-production.txt`

No options/customization available currently.

## Getting started

Make sure you have `GOPATH` set. Then install `nestguard` by:

    go get github.com/vigoo/nestguard
    
Run `GOPATH/bin/nestguard` from the root of a git repository.
