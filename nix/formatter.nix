{
  pkgs,
  inputs,
  ...
}:
inputs.treefmt-nix.lib.mkWrapper pkgs {
  projectRootFile = ".git/config";

  programs =
    {
      alejandra.enable = true;
      deadnix.enable = true;
      gofumpt.enable = true;
      prettier.enable = true;
      statix.enable = true;
    }
    // pkgs.lib.optionalAttrs (pkgs.system != "riscv64-linux") {
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

      alejandra = {
        priority = 3;
      };

      prettier = {
        options = [
          "--tab-width"
          "4"
        ];
        includes = ["*.{css,html,js,json,jsx,md,mdx,scss,ts,yaml}"];
      };
    };
  };
}
