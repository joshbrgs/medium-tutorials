import { connect } from "https://deno.land/x/amqp_ts/mod.ts";

export async function connectToRabbitMQ() {
  const connection = await connect({
    hostname: "localhost",
    port: 5672,
    username: "guest",
    password: "guest",
  });

  const channel = await connection.openChannel();
  await channel.declareQueue("test_queue", { durable: true });

  console.log("Connected to RabbitMQ and queue declared");

  return { connection, channel };
}

export async function consumeFromQueue(channel: any) {
  const messages: string[] = [];

  await channel.consume("test_queue", (msg: any) => {
    const messageContent = new TextDecoder().decode(msg.body);
    console.log("Received:", messageContent);
    messages.push(messageContent);
  });

  return messages;
}
