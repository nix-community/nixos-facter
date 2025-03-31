args @ {
  # We need the following pragma to ensure deadnix doesn't remove inputs.
  # This package is being called with newScope/callPackage, which means it is only being passed args it defines.
  # We do not use inputs directly in this file, but need it for passing to the tests.
  flake,
  # deadnix: skip
  inputs,
  # deadnix: skip
  system,
  perSystem,
  pkgs,
  ...
}: let
  inherit (pkgs) lib;
in
  pkgs.callPackage ./package.nix {
    hwinfo = perSystem.hwinfo.default;

    # there's no good way of tying in the version to a git tag or branch
    # so for simplicity's sake we set the version as the commit revision hash
    # we remove the `-dirty` suffix to avoid a lot of unnecessary rebuilds in local dev
    version = lib.removeSuffix "-dirty" (flake.shortRev or flake.dirtyShortRev);
  }
  // {
    passthru.tests = import ./tests args;
  }
