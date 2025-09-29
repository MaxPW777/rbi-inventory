{
  description = "Dev env for inventory suite";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};
    in {
      devShells.default = pkgs.mkShell {
        # Toolchain & services
        buildInputs = [
          pkgs.go_1_24 # Go toolchain
          pkgs.delve # Debugger (dlv)
          pkgs.go-tools # go vet, godoc, etc.
          pkgs.gotestsum # Pretty test runner output

          pkgs.sqlc # Type-safe DB access codegen
          pkgs.goose # DB migrations

          pkgs.just # Task runner
          pkgs.pnpm_10 # Better npm

          pkgs.docker # Optional: local containers
          pkgs.docker-compose
        ];

        shellHook = ''
          export GO111MODULE=on
          export CGO_ENABLED=1

          # Add ./bin (if you generate helpers) to PATH
          export PATH="$PWD/modules/api/bin:$PATH"
        '';
      };
    });
}
