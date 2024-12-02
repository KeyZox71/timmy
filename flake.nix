{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    pogit = {
      url = "github:y-syo/pogit";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    inputs@{ nixpkgs, self, ... }:
    let
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forEachSupportedSystem =
        f: nixpkgs.lib.genAttrs supportedSystems (system: f { pkgs = import nixpkgs { inherit system; }; });
    in
    {
      packages = forEachSupportedSystem (
        { pkgs }:
        rec {
          default = timmy;
          timmy = pkgs.buildGoModule {
			src = self;
			pname = "timmy";
			subPackages = [ "cmd/timmy" ];
			version = "0.1";
			vendorHash = "sha256-l/+TXNT7Z/CbnVCzB0B8VA7Fkj+MOOL1s3QHnZkAsUg=";
          };
        }
      );
      devShells = forEachSupportedSystem (
        { pkgs }:
        {
          default = pkgs.mkShell.override { } {
            buildInputs = with pkgs; [

            ];
            packages = with pkgs; [
              go
              gopls
              inputs.pogit.packages.${pkgs.system}.default
            ];
          };
        }
      );
    };
}
