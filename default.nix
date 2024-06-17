# { pkgs ? import <nixpkgs> {} }:
# let
# in
#   pkgs.mkShell {
#     buildInputs = [
#       pkgs.nodejs_20
#       pkgs.go
#       pkgs.air
#       pkgs.go-swag
#       pkgs.gopls
#     ];

#     # export GOPATH="/home/jbuhre/development/joepbuhre/snappic/backend/.gopath"
#     # export PATH="$PATH:/home/jbuhre/development/joepbuhre/snappic/backend/.gopath/bin"
#   shellHook = ''
#     export GOROOT="${pkgs.go}"
#   '';
# }

{ pkgs ? import <nixpkgs> { } }:

with pkgs;

let
  rstudio-custom = rstudioWrapper.override{ packages = with rPackages; [ ggplot2 dplyr rmarkdown knitr ]; };

in
mkShell {

  PROJECT_ROOT = builtins.toString ./.;

  buildInputs = [
    nodejs_20
    go
    gotools
    gopls
    go-outline
    gopkgs
    gocode-gomod
    godef
    golint
    air
    delve
    http-server
    sqlc

    # specify data analysis tools
    rstudio-custom

  ];
  shellHook = ''
    export CUR_DIR=/home/jbuhre/development/joepbuhre/api-tester
    export PATH="$PATH:$CUR_DIR/backend/.gopath/bin"
    export GOPATH="$CUR_DIR/backend/.gopath"


    cd "$CUR_DIR"
    my_pid=$$

    start-rstudio () {
      rstudio > /dev/null 2>&1 & 
    }


  '';
}

