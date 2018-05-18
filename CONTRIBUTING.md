# How to contribute
labchat is MIT licensed and accepts contributions via GitHub pull requests. This doucument outlines some of the conventions on commit message formatting, contact points for developers, and other resources to help get contributions into labchat.

## Getting started

- Fork the repository on GitHub
- Read the README.md for build instructions

## Reporting bugs and creating issues

Email below account about reporting bugs and issues.

- Email : [labchat-dev](https://groups.google.com/forum/?hl=en#!forum/labchat-dev)

## Contribution flow

- Create a branch which don't affect the 'master' branch.
- Commit your changes per file.
- Make sure commit messages which comply with below commit format.
- Submit a pull request to yoonsue/labchat.

### Code style

The coding style suggested by the Golang community is used in labchat.

Please follow this style to make labchat easy to review, maintain and develop.

### Format of the commit message

Please follow this format:
```sh
git commit -m "{filename what you changed or make}: {shortcut of what you changed}"
```

This is an example for above:
```sh
git commit -m "test.go: func send URL is changed"
```


