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
  const { value: showNewMessage } = useFlag('show-banner', false);
  const [message, setMessage] = useState<string>("");
  const [events, setEvents] = useState<RawEvent[]>([]);
  const [chartData, setChartData] = useState<ChartData[]>([]);

  // Stable callback that won't cause websocket reconnections
  const handleMessage = useCallback((event: RawEvent) => {
    console.log("Received websocket event:", event);
    
    setEvents((prev) => {
      const updated = [...prev, event].slice(-100); // last 100 events
      const trueCount = updated.filter((e) => e.value).length;
      const percentTrue = Math.round((trueCount / updated.length) * 100);
      
      setChartData((prevChart) => [
        ...prevChart.slice(-50), // Keep last 50 chart points
        { timestamp: Date.now(), percentTrue },
      ]);
      
      return updated;
    });
  }, []); // Empty dependency array - callback never changes

  useWebSocket(handleMessage);

  useEffect(() => {
    fetch("http://localhost:3001/v1/welcome")
      .then((res) => res.json())
      .then((data) => setMessage(data))
      .catch((err) => {
        console.error("Failed to fetch welcome message:", err);
        setMessage("Failed to load welcome message");
      });
  }, []);

  return (
    <div className="m-5 py-6">
      <Alert className="p-5">
        <AlertTitle>FeatureFlag Banner</AlertTitle>
        {showNewMessage ? 
          <AlertDescription>Welcome to this OpenFeature-enabled React app!</AlertDescription> :
          <AlertDescription>Welcome to this React app.</AlertDescription>
        }
      </Alert>

      <Alert>
        <AlertTitle>User welcome message</AlertTitle>
        <AlertDescription>{message}</AlertDescription>
      </Alert>

      <div className="mt-6">
        <h1 className="text-xl font-semibold mb-4">Live Feature Flag Status</h1>
        <p className="text-sm text-gray-600 mb-4">
          Events processed: {events.length} | Latest: {events[events.length - 1]?.value ? 'True' : 'False'}
        </p>
        <FeatureChart data={chartData} />
      </div>
    </div>
  )
}
