{
  description = "NixOS Facter";

  # Add all your dependencies here
  inputs = {
    blueprint = {
      url = "github:numtide/blueprint";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.systems.follows = "systems";
    };
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "flake-utils";
    };
    systems.url = "github:nix-systems/default";
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
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

    godoc.url = "github:numtide/godoc";
    godoc.inputs.nixpkgs.follows = "nixpkgs";
    godoc.inputs.systems.follows = "systems";
    godoc.inputs.blueprint.follows = "blueprint";
    godoc.inputs.treefmt-nix.follows = "treefmt-nix";
    godoc.inputs.flake-utils.follows = "flake-utils";
    godoc.inputs.gomod2nix.follows = "gomod2nix";
  };

  # Keep the magic invocations to minimum.
  outputs = inputs:
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
