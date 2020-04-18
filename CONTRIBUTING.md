# Contributing to user_agent

## Check that your changes do not break anything

You can safely run the `ci` target locally:

```
$ make ci
```

This will run all the make targets that will also be run on the CI, which are
described below.

### Running tests

Call the `test` make target:

```bash
$ make test
```

### Checking the style

This project uses some standard tools in order to check the Go style. You can
run them by calling the `validate-go` make target:

```bash
$ make validate-go
```

This target is included inside of the `validate` one.

### Git validation

In order to ensure that the git log is as maintainable as possible, the
[git-validation](https://github.com/vbatts/git-validation) tool is used. You can
install this tool by running:

```bash
$ go get -u github.com/vbatts/git-validation
```

If you already have this tool installed, then simply perform:

```bash
$ make git-validation
```

This target is included inside of the `validate` one.

## Issue reporting

I'm using [Github](https://github.com/mssola/user_agent) in order to host the
code. Thus, in order to report issues you can do it on its [issue
tracker](https://github.com/mssola/user_agent/issues). A couple of notes on
reports:

- Check that the issue has not already been reported or fixed in `master`.
- Try to be concise and precise in your description of the problem.
- Provide a step by step guide on how to reproduce this problem.
- Provide the version you are using (the commit SHA, if possible).

## Pull requests

- Write a [good commit message](https://chris.beams.io/posts/git-commit/).
- Make sure that tests are passing on your local machine (it will also be
checked by the CI system whenever you submit the pull request).
- Update the [changelog](./CHANGELOG.md).
- Try to use the same coding conventions as used in this project.
- Open a pull request with *only* one subject and a clear title and
description. Refrain from submitting pull requests with tons of different
unrelated commits.
