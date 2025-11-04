#!/usr/bin/env nix-shell
#!nix-shell -i bash -p nix -p coreutils -p gnused -p gawk
# shellcheck shell=bash

set -exuo pipefail

cd "$(git rev-parse --show-toplevel)"

go mod tidy

# shellcheck disable=SC2016
failedbuild=$(nix build --log-format bar-with-logs --impure --expr '(builtins.getFlake (toString ./.)).packages.${builtins.currentSystem}.nixos-facter.overrideAttrs (_:{ vendorHash = ""; })' 2>&1 || true)
echo "$failedbuild"
checksum=$(echo "$failedbuild" | awk '/got:.*sha256/ { print $2 }')

if [ -z "$checksum" ]; then
  echo "Error: Could not extract checksum from build output"
  exit 1
fi

echo "$checksum" > nix/packages/nixos-facter/goVendorHash.txt
