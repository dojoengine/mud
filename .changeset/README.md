# Changesets

Hello and welcome! This folder has been automatically generated by `@changesets/cli`, a build tool that works
with multi-package repos, or single-package repos to help you version and publish your code. You can
find the full documentation for it [in our repository](https://github.com/changesets/changesets)

We have a quick list of common questions to get you started engaging with this project in
[our documentation](https://github.com/changesets/changesets/blob/main/docs/common-questions.md)

## How changesets work in this repo

#### Repo is in normal mode

Normal mode means you don't have a prerelease

- Merging to main will make a versioning pr compiling all the changesets moving the repo into preset mode
- Until this pr merges main is still in normal mode
- In normal mode you can keep making patches to the release branch from main still without needing to cherry-pick manually publish
- As prs keep being added this pr will keep taking on commits
- When ready to release, enter prerelease mode via merging

#### Repo is in prerelease mode

- Merging to main will make/update a versioning pr
- Merge this versioning pr to make another prerelease increasing the nonce by 1 each time

#### Time to actually publish

- Running `changeset pre exit` will exit prerelease mode
- For convenience there is a github action to do it
- Once out of prerelease mode you can now publish normally do a normal changeset flow

```
Warning: If you decide to do prereleases from the main branch of your repository, without having a branch for your last stable release without the prerelease changes, you will block other changes until you are ready to exit prerelease mode. We thoroughly recommend only running prereleases from a branch other than the main branch.
```