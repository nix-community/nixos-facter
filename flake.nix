{
  description = "NixOS Facter";

  # Add all your dependencies here
  inputs = {
    blueprint = {
      url = "github:numtide/blueprint";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.systems.follows = "systems";
    };
    systems.url = "github:nix-systems/default";
    nixpkgs.url = "git+https://github.com/NixOS/nixpkgs?shallow=1&ref=nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    flake-utils.inputs.systems.follows = "systems";
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    disko.url = "github:nix-community/disko";
    disko.inputs.nixpkgs.follows = "nixpkgs";

    hwinfo.url = "github:numtide/hwinfo";
    hwinfo.inputs.nixpkgs.follows = "nixpkgs";
    hwinfo.inputs.systems.follows = "systems";
    hwinfo.inputs.blueprint.follows = "blueprint";
  };

  # Keep the magic invocations to minimum.
  outputs =
    inputs:
    inputs.blueprint {
      prefix = "nix/";
      inherit inputs;
      systems = [
        "aarch64-linux"
        "riscv64-linux"
        "x86_64-linux"
      ];
    };
}
