"use client";

import { useMarketData } from "@/lib/useMarketData";

export function MarketList() {
  const { data, loading, error } = useMarketData();

  if (loading) return <div className="animate-pulse text-gray-400">Loading...</div>;
  if (error) return <div className="text-red-500">{error}</div>;

  return (
    <div className="space-y-4">
      {data.map(({ IndexName, Timestamp, ClosePrice, OpenPrice, Currency, Region }) => {
        const change = ClosePrice - OpenPrice;
        const percent = (change / OpenPrice) * 100;

        return (
          <div key={`${IndexName}-${Timestamp}`} className="rounded-lg p-8">
            <article>
              <header>
                <h2>{IndexName}</h2>
                <span>{Region}</span>
              </header>
              <section>
                <p> Close: <span>{ClosePrice?.toFixed(2)} {Currency}</span></p>
                <p>Change: {" "}
                  <span className={change >= 0 ? "text-green-500" : "text-red-500"}>
                    {change >= 0 ? "+" : ""}
                    {change?.toFixed(2)} ({percent?.toFixed(2)}%)
                  </span>
                </p>
              </section>
              <footer>
                <p>
                  Updated: <span>{new Date(Timestamp)?.toLocaleString()}</span>
                </p>
              </footer>
            </article>
          </div>
        );
      })}
    </div>
  );
}
