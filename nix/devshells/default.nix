{
  perSystem,
  pkgs,
  ...
}:
perSystem.self.nixos-facter.overrideAttrs (old: {
  GOROOT = "${old.passthru.go}/share/go";
  nativeBuildInputs = old.nativeBuildInputs ++ [
    pkgs.enumer
    pkgs.delve
    pkgs.gotools
    pkgs.golangci-lint
    pkgs.cobra-cli
    pkgs.fx # json tui
    perSystem.hwinfo.default
  ];
  shellHook = ''
    # this is only needed for hermetic builds
    unset GO_NO_VENDOR_CHECKS GOSUMDB GOPROXY GOFLAGS
  '';
})
