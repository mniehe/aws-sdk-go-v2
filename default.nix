{ pkgs ? import <nixpkgs> { } }:
with pkgs;
mkShell {
  name = "rivian-cli";
  packages = [
    buildPackages.go
  ];

  buildInputs = [
    go
    gotools
    gopls
    go-outline
    gocode
    gopkgs
    gocode-gomod
    godef
    golint

    awscli2
    graphviz
  ];
}

