{
  pkgs,
  inputs,
  perSystem,
  ...
}:
let
  # we have to import diskoLib like this because there are some impure default imports e.g. <nixpkgs>
  diskoLib = import "${inputs.disko}/lib" {
    inherit (pkgs) lib;
    makeTest = import "${inputs.nixpkgs}/nixos/tests/make-test-python.nix";
    eval-config = import "${inputs.nixpkgs}/nixos/lib/eval-config.nix";
  };
in
# for now we only run the tests in x86_64-linux since we don't have access to a bare-metal ARM box or a VM that supports nested
# virtualisation which makes the test take forever and ultimately fail
pkgs.lib.optionalAttrs pkgs.stdenv.isx86_64 {
  golangci-lint = perSystem.self.nixos-facter.overrideAttrs (old: {
    nativeBuildInputs = old.nativeBuildInputs ++ [ pkgs.golangci-lint ];
    buildPhase = ''
      HOME=$TMPDIR
      golangci-lint run
    '';
    installPhase = ''
      touch $out
    '';
    doInstallCheck = false;
    doCheck = false;
  });

  basic = diskoLib.testLib.makeDiskoTest {
    inherit pkgs;
    name = "basic";
    disko-config = ./disko.nix;
    extraSystemConfig = {
      environment.systemPackages = [
        perSystem.self.nixos-facter
      ];
    };

    extraTestScript = ''
      import json

      report = json.loads(machine.succeed("nixos-facter --ephemeral 2>&1"))

      with subtest("Capture system"):
          assert report['system'] == '${pkgs.stdenv.hostPlatform.system}'

      with subtest("Capture virtualisation"):
          virt = report['virtualisation']
          # kvm for systems that support it, otherwise the vm test should present itself as qemu
          # todo double-check this is the same for intel
          assert virt in ("kvm", "qemu"), f"expected virtualisation to be either kvm or qemu, got {virt}"

      with subtest("Capture swap entries"):
          assert 'swap' in report, "'swap' not found in the report"

          swap = report['swap']

          expected = [
              { 'type': 'partition', 'size': 1048572, 'used': 0, 'priority': -2 },
              { 'type': 'partition', 'size': 10236, 'used': 0, 'priority': 100 }
          ]

          assert len(swap) == len(expected), f"expected {len(expected)} swap entries, found {len(swap)}"

          for i in range(2):
              assert swap[i]['path'].startswith("/dev/disk/by-uuid/"), f"expected a stable device path: {swap[i]['path']}"

              # delete for easier comparison
              del swap[i]['path']

              assert swap[i] == expected[i], f"swap[{i}] != expected[{i}] mismatch: {swap[i]} != {expected[i]}"
    '';
  };
}
