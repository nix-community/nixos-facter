#!/usr/bin/env bash

set -eu -o pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"
cd "$SCRIPT_DIR/.."

version=${1:-}
if [[ -z $version ]]; then
  echo "USAGE: $0 version" >&2
  exit 1
fi

if [[ "$(git symbolic-ref --short HEAD)" != "main" ]]; then
  echo "must be on main branch" >&2
  exit 1
fi

# ensure we are up-to-date
uncommitted_changes=$(git diff --compact-summary)
if [[ -n $uncommitted_changes ]]; then
  echo -e "There are uncommitted changes, exiting:\n${uncommitted_changes}" >&2
  exit 1
fi
git pull git@github.com:numtide/nixos-facter main
unpushed_commits=$(git log --format=oneline origin/main..main)
if [[ $unpushed_commits != "" ]]; then
  echo -e "\nThere are unpushed changes, exiting:\n$unpushed_commits" >&2
  exit 1
fi
sed -i "s/version = \".*\"/version = \"$version\"/" ./nix/packages/nixos-facter/package.nix
git add ./nix/packages/nixos-facter/package.nix
git commit -m "release: v${version}"
nix flake check -vL
git tag "v${version}"

echo "now run 'git push --tags origin main'"
