#!/bin/bash

git tag -d v0.0.7
git tag -d v0.3.10
git tag -d v0.3.11
git tag -d v0.5.01
git tag -d v0.5.11
git tag -d v0.5.12
git tag -d v0.5.13
git tag -d v0.5.14
git tag -d v0.5.15
git tag -d v0.5.16
git tag -d v0.5.17
git tag -d v0.5.18

git push origin :refs/tags/v0.0.7
git push origin :refs/tags/v0.3.10
git push origin :refs/tags/v0.3.11
git push origin :refs/tags/v0.5.01
git push origin :refs/tags/v0.5.11
git push origin :refs/tags/v0.5.12
git push origin :refs/tags/v0.5.13
git push origin :refs/tags/v0.5.14
git push origin :refs/tags/v0.5.15
git push origin :refs/tags/v0.5.16
git push origin :refs/tags/v0.5.17
git push origin :refs/tags/v0.5.18

