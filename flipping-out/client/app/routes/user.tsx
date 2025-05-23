import { useFlag } from "@openfeature/react-sdk";
import {
  Alert,
  AlertDescription,
  AlertTitle,
} from "@/components/ui/alert"
import { useState, useEffect, useCallback } from "react"
import FeatureChart from "@/components/FeatureChart";
import { useWebSocket } from "../lib/socket.ts";

type RawEvent = {
  value: boolean;
};

type ChartData = {
  timestamp: number;
  percentTrue: number;
};

export default function Page() {
  // Use the "query-style" flag evaluation hook, specifying a flag-key and a default value.
  const { value: showNewMessage } = useFlag('show-banner', false);
  const [message, setMessage] = useState<string>("");
  const [events, setEvents] = useState<RawEvent[]>([]);
  const [chartData, setChartData] = useState<ChartData[]>([]);

  const handleMessage = useCallback((event: RawEvent) => {
    setEvents((prev) => {
      const updated = [...prev, event].slice(-100); // last 100 events
      const trueCount = updated.filter((e) => e.value).length;
      const percentTrue = Math.round((trueCount / updated.length) * 100);
      setChartData([
        ...chartData,
        { timestamp: Date.now(), percentTrue },
      ]);
      return updated;
    });
  }, [chartData]);

  useWebSocket(handleMessage);

  useEffect(() => {
    fetch("http://localhost:3001/v1/welcome")
      .then((res) => res.json())
      .then((data) => setMessage(data));
  }, []);

  return (
    <div className="m-5 py-6">
      <Alert className="p-5">
        <AlertTitle>FeatureFlag Banner</AlertTitle>
        {showNewMessage ? <AlertDescription>Welcome to this OpenFeature-enabled React app!</AlertDescription> :
          <AlertDescription>Welcome to this React app.</AlertDescription>}
      </Alert>

      <Alert>
        <AlertTitle>User welcome message</AlertTitle>
        <AlertDescription>{message}</AlertDescription>
      </Alert>

      <h1 className="text-xl font-semibold mb-4">Live Feature Flag Status</h1>
      <FeatureChart data={chartData} />

    </div>
  )
}
