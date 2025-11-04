{
  pkgs,
  flake,
  inputs,
  ...
}:
let
  mod = inputs.treefmt-nix.lib.evalModule pkgs {
    projectRootFile = ".git/config";

    programs = {
      nixfmt.enable = true;
      deadnix.enable = true;
      gofumpt.enable = true;
      prettier.enable = true;
      statix.enable = true;
    }
    // pkgs.lib.optionalAttrs (pkgs.stdenv.system != "riscv64-linux") {
      shellcheck.enable = true;
    };

    settings = {
      global.excludes = [
        "LICENSE"
        # unsupported extensions
        "*.{gif,png,svg,tape,mts,lock,mod,sum,toml,env,envrc,gitignore}"
      ];

      formatter = {
        deadnix = {
          priority = 1;
        };

        statix = {
          priority = 2;
        };

        nixfmt = {
          priority = 3;
        };

        prettier = {
          options = [
            "--tab-width"
            "4"
          ];
          includes = [ "*.{css,html,js,json,jsx,md,mdx,scss,ts,yaml}" ];
        };
      };
    };
  };

  wrapper = mod.config.build.wrapper // {
    passthru.tests.check = mod.config.build.check flake;
  };

  unsupported = pkgs.writeShellApplication {
    name = "unsupported-platform";
    text = ''
      echo "nix fmt is not supported on ${pkgs.stdenv.hostPlatform.system}";
    '';
  };
in
# nixfmt-rfc-style is based on Haskell, which is broke on RiscV currently
if pkgs.stdenv.hostPlatform.isRiscV then unsupported else wrapper
