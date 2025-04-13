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

###  Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) before opening issues or pull requests.

### App flow
[Fetch + Normalize] → [Store in DB] → [GET endpoints] → [Frontend renders] → [Filtering / Sorting]

[Timezone Utils] → [Market Timer]
[Theme Toggle] → [Layout] → [Testing]

### Stock API Considerations

| **API Provider**         | **Cost (Free Tier)**              | **Market Coverage**               | **Historical Data Availability** | **API Limitations (Free Tier)**                          |
|---------------------------|-----------------------------------|------------------------------------|-----------------------------------|----------------------------------------------------------|
| **Financial Modeling Prep (FMP)** | Free tier available              | S&P 500, Nikkei 225, FTSE 100      | Daily data from 1990 onwards     | Rate limits apply; real-time data may be limited         |
| **Alpha Vantage**         | Free tier (5 requests/min, 500/day) | S&P 500, Nikkei 225, FTSE 100      | Daily data available             | Limited to 5 requests per minute and 500 per day         |
| **Finnhub**               | Free tier (60 requests/min)       | S&P 500, Nikkei 225, FTSE 100      | Historical data available         | Some data may require premium access                     |
| **Marketstack**           | Free tier (100 requests/month)    | S&P 500, Nikkei 225, FTSE 100      | Historical data available         | Limited to 100 requests per month                        |
| **Polygon.io**            | No free tier for indices          | S&P 500, Nikkei 225, FTSE 100      | Historical data available         | Requires paid subscription for index data                |
| **QUODD**                 | Enterprise pricing                | S&P 500, Nikkei 225, FTSE 100      | Historical data available         | Enterprise-level API; pricing upon request               |

### Financial Modeling Prep

Since they offer the best historic and market data, I decided to go with them for now.
Plan is to get the historic data, save it into the db, fetch the market data every 12 hours based on when the markets open and save the new data in to the db.
