# Dagreleaser

This is a proof of concept module that demonstrates how to release distributable packages using Dagger, similar to how [Goreleaser](https://goreleaser.com/) works.

**Note**: This is a proof of concept and is not intended for production use.

It uses the experimental interface support.

## Demo

Building a release:

```shell
cd demo
dagger call release -o output
open output
```

Publishing to GitHub:

```shell
dagger call release-and-publish --token env:GITHUB_TOKEN
```
