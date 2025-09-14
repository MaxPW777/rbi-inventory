# Riviera Beauty Interiors inventory app

This is an internal inventory management system for **Riviera Beauty
Interiors**.  
It helps track stock, organize items into folders, record supplier
details, and monitor shortages or usage over time.

The app provides:

- **Folders** to group items (for example, product categories or brands).
- **Items** with detailed properties: quantity, weight, volume, expiry
  dates, prices, minimum stock levels, and notes.
- **Suppliers** linked to items, so you know who to call when stock
  runs low.
- **Usages** to record stock changes (for example, when items are used or
  sold).
- **Photos** to visually document items.
- **Audit logs** to track user actions in the system.
- **Statistics** on:
  - Items running low (quantity below minimum level).
  - Items expiring soon.
  - Shortages over time.

---

## Getting started

### Ruby version

- Ruby **3.4.x** (managed with a Nix flake; see below).

### System dependencies

- [Nix](https://nixos.org/) (with flakes enabled)
- PostgreSQL 15 (auto-configured by the dev shell)
- Node.js 20 and Yarn (for JavaScript/CSS assets)
- Bundler (`gem install bundler`) if not already available

### Development environment

This project includes a **Nix flake** (`flake.nix`) that provides a full
development shell with:

- Ruby 3.4
- PostgreSQL (local instance at `localhost:5432`)
- Native dependencies for Rails gems (libxml2, libxslt, OpenSSL, etc.)

To enter the environment:

```bash
direnv allow   # if you use direnv
# or
nix develop
```

---

## Configuration

Environment variables:

- `DATABASE_URL` is set automatically by the dev shell to  
  `postgres://postgres:postgres@localhost:5432/rbi_inventory`.

Bundler is configured to install gems into a local `.gems/` directory.

---

## Database setup

Create and migrate the database:

```bash
bin/rails db:create
bin/rails db:migrate
bin/rails db:seed   # optional, add seed data here
```

Check schema consistency:

```bash
bin/rails db:schema:load
```

---

## Running the app

Start the Rails server:

```bash
bin/rails server
# → http://localhost:3000
```

Routes:

- `GET /` homepage with search field .
- `GET /items/:id` show details of one item .
- `GET /api/v1/items` list items .
- `GET /api/v1/items/:id` show one item .
- `POST /api/v1/items/:id/usages` record a usage or shortage .
- `GET /api/v1/folders`, `GET /api/v1/folders/:id` folders .
- `GET /api/v1/suppliers`, `GET /api/v1/suppliers/:id` suppliers .

---

## Usage concepts

- **Folders** organize items.
- **Items** belong to one folder and one supplier.
- **Suppliers** provide contact info so staff can quickly reorder
  items.
- **Minimum level (`min_level`)**: when stock (`quantity`) drops below
  this threshold, the item is flagged as low stock.
- **Usages**: record each stock change, whether positive or negative.
- **Audit logs**: record user actions for accountability.
- **Statistics**:
  - Items below minimum stock.
  - Items expiring soon.
  - Usage trends.

---

## Running tests

To run the test suite:

```bash
bin/rails test
```

---

## Services

- **PostgreSQL**: database (local instance started by the dev
  shell).
- **Rails Active Job**: background jobs (can use async in development).
- **Action Mailer**: email delivery (configure Simple Mail Transfer Protocol, or SMTP, in production if needed).

---

## Deployment instructions

1.  Provision a PostgreSQL database.
2.  Set environment variables:
    - `DATABASE_URL=postgres://...`
    - `RAILS_MASTER_KEY` (for encrypted credentials if used).
3.  Install system dependencies (Ruby, Node.js, Yarn, PostgreSQL client
    libraries).
4.  Run:

    ```bash
    bundle install
    yarn install --check-files
    bin/rails db:migrate
    bin/rails assets:precompile
    ```

5.  Start the app:

    ```bash
    bin/rails server -e production
    ```

---

## Monitoring

- `GET /up` — returns `200 OK` if the app booted successfully (for
  uptime checks).
- `GET /service-worker` and `GET /manifest`: Progressive Web App (PWA) endpoints.

---

## Contributing

1.  Fork and branch from `main`.
2.  Follow Rails conventions (controllers under `app/controllers`,
    models under `app/models`).
3.  Write tests for new features.
4.  Submit a pull request.

---

## License

This app is proprietary to **Riviera Beauty Interiors**.  
Unauthorized copying or distribution is prohibited.
