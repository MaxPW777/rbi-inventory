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
        # toolchain
        buildInputs = [
          pkgs.elixir_1_17
          pkgs.erlang_26
          pkgs.nodejs_20
          pkgs.pnpm
          pkgs.postgresql_16
          pkgs.redis
          pkgs.just # better than make for scripts
          pkgs.openssl
          pkgs.git
        ];
        shellHook = ''
          export MIX_ENV=dev
          export PHX_SERVER=true
          export PATH=$PATH:./node_modules/.bin
        '';
      };
    });
}
