export type MarketData = {
    IndexName: string;
    Region: string;
    Currency: string;
    Timestamp: string;
    OpenPrice: number;
    ClosePrice: number;
    High: number;
    Low: number;
    Volume: number;
};

export async function fetchLatestMarketData(): Promise<MarketData[]> {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/markets/latest`);

    if(!res.ok) {
        throw new Error('Failed to fetch market data');
    }

    return res.json();
}
