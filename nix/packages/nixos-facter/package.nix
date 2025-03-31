{
  lib,
  version ? "0.3.1",
  systemdMinimal,
  hwinfo,
  gcc,
  makeWrapper,
  pkg-config,
  stdenv,
  buildGo124Module,
}: let
  fs = lib.fileset;
in
  buildGo124Module (final: {
    pname = "nixos-facter";
    inherit version;

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

    vendorHash = "sha256-A7ZuY8Gc/a0Y8O6UG2WHWxptHstJOxi4n9F8TY6zqiw=";

    buildInputs = [
      systemdMinimal.dev
      hwinfo
    ];

    nativeBuildInputs = [
      gcc
      makeWrapper
      pkg-config
    ];

    ldflags = [
      "-s"
      "-w"
      "-X github.com/numtide/nixos-facter/pkg/build.Name=${final.pname}"
      "-X github.com/numtide/nixos-facter/pkg/build.Version=v${final.version}"
      "-X github.com/numtide/nixos-facter/pkg/build.System=${stdenv.hostPlatform.system}"
    ];

    postInstall = let
      binPath = lib.makeBinPath [
        systemdMinimal
      ];
    in ''
      wrapProgram "$out/bin/nixos-facter" \
          --prefix PATH : "${binPath}"
    '';

    meta = with lib; {
      description = "nixos-facter: declarative nixos-generate-config";
      homepage = "https://github.com/numtide/nixos-facter";
      license = licenses.mit;
      mainProgram = "nixos-facter";
    };
  })
