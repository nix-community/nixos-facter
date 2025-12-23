{
  perSystem,
  pkgs,
  ...
}:
let
  inherit (pkgs) stdenv;
  commonDevTools = [
    pkgs.enumer
    pkgs.delve
    pkgs.gotools
    pkgs.golangci-lint
    pkgs.cobra-cli
    pkgs.fx # json tui
  ];
  go = pkgs.go_1_24;
in
if stdenv.hostPlatform.isLinux then
  perSystem.self.nixos-facter.overrideAttrs (old: {
    GOROOT = "${old.passthru.go}/share/go";
    nativeBuildInputs =
      old.nativeBuildInputs
      ++ commonDevTools
      ++ [
        perSystem.hwinfo.default
      ];
    shellHook = ''
      # this is only needed for hermetic builds
      unset GO_NO_VENDOR_CHECKS GOSUMDB GOPROXY GOFLAGS
    '';
  })
else
  # macOS: devshell for editing, go generate, linting
  # Use `nix build .#packages.x86_64-linux.nixos-facter` for full builds (uses remote builder)
  pkgs.mkShell {
    GOROOT = "${go}/share/go";
    # Cross-compile target for go build
    GOOS = "linux";
    GOARCH = "amd64";
    nativeBuildInputs = commonDevTools ++ [
      go
    ];
    shellHook = ''
      unset GO_NO_VENDOR_CHECKS GOSUMDB GOPROXY GOFLAGS
    '';
  }
