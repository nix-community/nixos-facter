args@{
  # We need the following pragma to ensure deadnix doesn't remove inputs.
  # This package is being called with newScope/callPackage, which means it is only being passed args it defines.
  # We do not use inputs directly in this file, but need it for passing to the tests.
  # deadnix: skip
  inputs,
  # deadnix: skip
  system,
  perSystem,
  pkgs,
  ...
}:
pkgs.callPackage ./package.nix {
  hwinfo = perSystem.hwinfo.default;
}
// {
  passthru.tests = import ./tests args;
}
