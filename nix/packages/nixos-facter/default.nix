{
  inputs,
  perSystem,
  pkgs,
  ...
}:
pkgs.callPackage ./package.nix {
  hwinfo = perSystem.hwinfo.default;
}
// {
  passthru.tests = import ./tests {
    inherit pkgs inputs perSystem;
  };
}
