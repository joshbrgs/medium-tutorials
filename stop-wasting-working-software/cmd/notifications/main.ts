import { serve } from "https://deno.land/std/http/server.ts";
import { connectToRabbitMQ, consumeFromQueue } from "./rabbitmq.ts";

let cachedMessages: string[] = [];

const { channel } = await connectToRabbitMQ();

const handler = async (req: Request): Promise<Response> => {
  if (req.method === "GET" && new URL(req.url).pathname === "/notify") {
    if (cachedMessages.length === 0) {
      cachedMessages = await consumeFromQueue(channel);
    }

    const response = {
      messages: cachedMessages,
    };

    return new Response(JSON.stringify(response), {
      headers: { "Content-Type": "application/json" },
    });
  }
  return new Response("Not Found", { status: 404 });
};

console.log("HTTP server running on http://localhost:8001");
await serve(handler, { port: 8001 });
