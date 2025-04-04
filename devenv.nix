{ pkgs, lib, config, inputs, ... }:

{
  packages = [ pkgs.git ];
  languages.go.enable = true;

  enterShell = ''
    echo "####################################"
    go version
    echo "####################################"
  '';

  # https://devenv.sh/pre-commit-hooks/
  pre-commit.hooks = {
    end-of-file-fixer.enable = true;
    shellcheck.enable = true;
    trim-trailing-whitespace.enable = true;
    typos.enable = true;
  };
}
