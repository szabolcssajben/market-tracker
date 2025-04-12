# Market Tracker

[![License: CC BY-NC 4.0](https://img.shields.io/badge/License-CC%20BY--NC%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc/4.0/)


| Layer  | Tech  | Purpose  |
|---|---|---|
|  Frontend |  Vercel + Next.js | One-page app (charts, UI, timers, theme, etc.)  |
| Backend  | Go (hosted on Railway)  |  API that fetches/caches stock data from external sources |
| Database | Supabase (PostgreSQL) | Stores historical stock data for caching + querying |

## Architecture Flow

1. Go API (Railway):

- Fetches stock data from third-party APIs
- Writes historical data to Supabase
- Serves normalized data to your frontend

2. Next.js Frontend (Vercel):

- Calls Go API for latest data
- Displays stock charts, market timers, sorting/filtering
- UI features: theme switcher, currency selector, etc.

3. Supabase:

- Stores index prices per market per day
- Store API response snapshots, conversion rates, metadata

## Testing (WIP)

### Backend (Go)

- Written using Go’s built-in `testing` package and `testify` for assertions and mocks.
- Coverage includes:
  - API data fetching and parsing
  - Supabase insert/query logic
  - Data normalization and struct validation
  - Utility functions (e.g. timezones, conversions)
- Tests are run with `go test ./...` via Railway CI and locally.

### Frontend (Next.js)

- End-to-end tests using **Playwright** (free and open source).
- Unit/component tests using **Vitest**.
- Coverage includes:
  - Chart rendering
  - Theme toggle
  - Market open/close timer logic
  - API integration layer
- Tests are run with `npm run test` (Vitest) and `npx playwright test` (E2E).

