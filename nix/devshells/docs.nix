{
  pkgs,
  ...
}:
pkgs.mkShellNoCC {
  packages = with pkgs.python3Packages; [
    mike
    mkdocs
    mkdocs-material
  ];
}
