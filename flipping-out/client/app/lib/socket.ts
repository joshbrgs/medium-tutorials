import { useEffect, useRef } from "react";

export function useWebSocket(onMessage: (data: any) => void) {
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:3001/v1/ws");

    socket.onmessage = (event) => {
      const parsed = JSON.parse(event.data);
      onMessage(parsed);
    };

    socket.onerror = (err) => {
      console.error("WebSocket error:", err);
    };

    socketRef.current = socket;

    return () => {
      socket.close();
    };
  }, [onMessage]);
}
