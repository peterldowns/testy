{
  description = "check is a library for writing meaningful and readable tests";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";

    flake-utils.url = "github:numtide/flake-utils";

    flake-compat.url = "github:edolstra/flake-compat";
    flake-compat.flake = false;
  };

  outputs = { self, ... }@inputs:
    inputs.flake-utils.lib.eachDefaultSystem (system:
      let
        # standard nix definitions
        overlays = [ ];
        pkgs = import inputs.nixpkgs {
          inherit system overlays;
        };
        lib = pkgs.lib;
        # localias specific
        version = (builtins.readFile ./VERSION);
      in
      rec {
        packages = rec { };
        apps = rec { };
        devShells = rec {
          default = pkgs.mkShell {
            packages = with pkgs;
              [
                # golang
                delve
                go-outline
                go
                golangci-lint
                gopkgs
                gopls
                gotools
                # nix
                rnix-lsp
                nixpkgs-fmt
                # xcode: this wrapper symlinks to the host system.
                # other tools
                just
              ];

            shellHook = ''
              # The path to this repository
              shell_nix="''${IN_LORRI_SHELL:-$(pwd)/shell.nix}"
              workspace_root=$(dirname "$shell_nix")
              export WORKSPACE_ROOT="$workspace_root"

              # We put the $GOPATH/$GOCACHE/$GOENV in $TOOLCHAIN_ROOT,
              # and ensure that the GOPATH's bin dir is on our PATH so tools
              # can be installed with `go install`.
              #
              # Any tools installed explicitly with `go install` will take precedence
              # over versions installed by Nix due to the ordering here.
              export TOOLCHAIN_ROOT="$workspace_root/.toolchain"
              export GOROOT=
              export GOCACHE="$TOOLCHAIN_ROOT/go/cache"
              export GOENV="$TOOLCHAIN_ROOT/go/env"
              export GOPATH="$TOOLCHAIN_ROOT/go/path"
              export GOMODCACHE="$GOPATH/pkg/mod"
              export PATH=$(go env GOPATH)/bin:$PATH
              export CGO_ENABLED=1

              # Make it easy to test while developing; add the golang and nix
              # build outputs to the path.
              export PATH="$workspace_root/bin:$workspace_root/result/bin:$PATH"
            '';

            # Need to disable fortify hardening because GCC is not built with -oO,
            # which means that if CGO_ENABLED=1 (which it is by default) then the golang
            # debugger fails.
            # see https://github.com/NixOS/nixpkgs/pull/12895/files
            hardeningDisable = [ "fortify" ];
          };
        };
      });
}
