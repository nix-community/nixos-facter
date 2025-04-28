{
  pkgs,
  perSystem,
  ...
}:
pkgs.mkShellNoCC {
  packages =
    with pkgs;
    # Pop an empty shell on systems that aren't supported by godoc
    lib.optionals (perSystem.godoc ? default) (
      [
        perSystem.godoc.default
        (pkgs.writeScriptBin "gen-reference" ''
          out="./docs/content/reference/go_doc"
          godoc -c -o $out .
        '')
        (pkgs.writeScriptBin "mkdocs" ''
          # generate reference docs first
          gen-reference
          # execute the underlying command
          ${pkgs.mkdocs}/bin/mkdocs "$@"
        '')
      ]
      ++ (with pkgs.python3Packages; [
        mike
        mkdocs-material
      ])
    );
}
