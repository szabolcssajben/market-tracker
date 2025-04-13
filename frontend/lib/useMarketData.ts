import { useEffect, useState } from "react";
import { fetchLatestMarketData, MarketData } from "./api";


export function useMarketData() {
  const [data, setData] = useState<MarketData[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchLatestMarketData()
      .then(setData)
      .catch((err) => setError("Failed to load market data " + err.message))
      .finally(() => setLoading(false));
  }, []);

  return { data, loading, error };
}
