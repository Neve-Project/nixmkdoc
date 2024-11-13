# SPDX-License-Identifier: GPL-2.0-or-later
{
  description = "nixMkDoc";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = {
    nixpkgs,
    self,
    ...
  }: let
    systems = ["x86_64-linux" "aarch64-linux"];
  in {
    packages = nixpkgs.lib.genAttrs systems (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        nixmkdoc = pkgs.stdenv.mkDerivation {
          pname = "nixmkdoc";
          version = "0.0.1";
          src = ./.;
          buildInputs = [pkgs.go];

          buildPhase = ''
            export GOCACHE=$(mktemp -d /tmp/go-build-cache-XXXXXX)
            go build -o nixmkdoc ./main.go
          '';

          installPhase = ''
            mkdir -p $out/bin
            cp nixmkdoc $out/bin/
          '';

          meta = with pkgs.lib; {
            description = "Write your nix documentation starting from lib.mkOption";
            license = licenses.mit;
            maintainers = with maintainers; [Matteo Cavestri];
          };
        };
      }
    );
    defaultPackage.x86_64-linux = self.packages.x86_64-linux.nixmkdoc;
    defaultPackage.aarch64-linux = self.packages.aarch64-linux.nixmkdoc;
  };
}
