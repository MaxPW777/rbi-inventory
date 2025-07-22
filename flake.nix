{
  description = "Full Rails DevShell (Nix Flake) with Ruby, PostgreSQL, Redis, native deps";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {inherit system;};
        ruby = pkgs.ruby;
        nativeDeps = [
          pkgs.openssl
          pkgs.libyaml
          pkgs.git
          pkgs.curl
          pkgs.redis
          pkgs.pkg-config
          pkgs.zlib
          pkgs.libffi
          pkgs.libxml2
          pkgs.libxslt
          pkgs.nodejs_20
          pkgs.yarn
        ];
        # Postgres service (for local dev DB)
        postgresPort = 5432;
        postgresUser = "postgres";
        postgresPassword = "postgres";
        postgresDb = "rbi_inventory";
        postgresDataDir = ".nix-pgdata";
      in {
        devShells.default = pkgs.mkShell {
          name = "rails-dev-shell";
          buildInputs = [ruby pkgs.bundler pkgs.postgresql] ++ nativeDeps;
          # Start Redis and Postgres in the background when entering shell
          shellHook = ''
            export GEM_HOME=$PWD/.gems
            export GEM_PATH=$GEM_HOME
            export PATH=$GEM_HOME/bin:$PATH

            export PGUSER=${postgresUser}
            export PGPASSWORD=${postgresPassword}
            export PGDATABASE=${postgresDb}
            export PGHOST=localhost
            export PGPORT=${toString postgresPort}
            export DATABASE_URL="postgres://${postgresUser}:${postgresPassword}@localhost:${toString postgresPort}/${postgresDb}"

            # Setup and start Postgres
            if [ ! -d "${postgresDataDir}" ]; then
              echo "Initializing Postgres database in ${postgresDataDir}..."
              initdb -D ${postgresDataDir}
              echo "host all all 0.0.0.0/0 md5" >> ${postgresDataDir}/pg_hba.conf
              echo "listen_addresses='*'" >> ${postgresDataDir}/postgresql.conf
            fi
            echo "Starting Postgres..."
            pg_ctl -D ${postgresDataDir} -o "-p ${toString postgresPort}" -w start
            createdb -h localhost -p ${toString postgresPort} -U ${postgresUser} ${postgresDb} 2>/dev/null || true
            psql -h localhost -p ${toString postgresPort} -U ${postgresUser} -c "ALTER USER ${postgresUser} WITH PASSWORD '${postgresPassword}';" 2>/dev/null || true

            # Start Redis
            echo "Starting Redis server..."
            (redis-server --daemonize yes)

            echo "Ruby: $(ruby --version)"
            echo "Bundler: $(bundler --version)"
            echo "Postgres, Redis started for dev env."
          '';
          # Clean up processes when you exit
          postShellHook = ''
            echo "Stopping Postgres and Redis..."
            pg_ctl -D ${postgresDataDir} -w stop || true
            redis-cli shutdown || true
          '';
        };
      }
    );
}
