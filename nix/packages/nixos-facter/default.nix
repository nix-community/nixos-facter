{
  flake,
  # We need the following pragma to ensure deadnix doesn't remove inputs.
  # This package is being called with newScope/callPackage, which means it is only being passed args it defines.
  # We do not use inputs directly in this file, but need it for passing to the tests.
  # deadnix: skip
  inputs,
  perSystem,
  system,
  pkgs,
  pname,
  ...
} @ args: let
  inherit (pkgs) go lib;
  fs = lib.fileset;
in
  perSystem.gomod2nix.buildGoApplication rec {
    inherit pname;
    # there's no good way of tying in the version to a git tag or branch
    # so for simplicity's sake we set the version as the commit revision hash
    # we remove the `-dirty` suffix to avoid a lot of unnecessary rebuilds in local dev
    version = lib.removeSuffix "-dirty" (flake.shortRev or flake.dirtyShortRev);

    # ensure we are using the same version of go to build with
    inherit go;

    src = fs.toSource {
      root = ../../..;
      fileset = fs.unions [
        ../../../cmd
        ../../../go.mod
        ../../../go.sum
        ../../../main.go
        ../../../pkg
      ];
    };

    modules = ./gomod2nix.toml;

    buildInputs = [
      pkgs.systemdMinimal.dev
      perSystem.hwinfo.default
    ];

    nativeBuildInputs = with pkgs; [
      gcc
      makeWrapper
      pkg-config
    ];

    ldflags = [
      "-s"
      "-w"
      "-X github.com/numtide/nixos-facter/pkg/build.Name=${pname}"
      "-X github.com/numtide/nixos-facter/pkg/build.Version=v${version}"
      "-X github.com/numtide/nixos-facter/pkg/build.System=${pkgs.stdenv.hostPlatform.system}"
    ];

    postInstall = let
      binPath = lib.makeBinPath (with pkgs; [
        systemdMinimal
      ]);
    in ''
      wrapProgram "$out/bin/nixos-facter" \
          --prefix PATH : "${binPath}"
    '';

    passthru.tests = (import ./tests) args;

    meta = with lib; {
      description = "nixos-facter: declarative nixos-generate-config";
      homepage = "https://github.com/numtide/nixos-facter";
      license = licenses.mit;
      mainProgram = "nixos-facter";
    };
  }
