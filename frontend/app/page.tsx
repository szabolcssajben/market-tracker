import { MarketList } from "@/components/MarketList";

export default function Home() {
  return (
    <main className="min-h-screen bg-white dark:bg-gray-900 text-black dark:text-white flex items-center justify-center flex-col">
      <h1 className="text-4xl font-bold">Market Tracker</h1>
      <MarketList />
    </main>
  );
}
