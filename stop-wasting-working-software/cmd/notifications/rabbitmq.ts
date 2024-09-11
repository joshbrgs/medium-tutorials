import { connect, Channel, ConsumeMessage } from "https://deno.land/x/amqp@v0.24.0/mod.ts";

export async function connectToRabbitMQ() {
  const connection = await connect({
    hostname: "rabbitmq",
    port: 5672,
    username: "guest",
    password: "guest",
  });

  const channel = await connection.openChannel();

  console.log("Connected to RabbitMQ");

  return { connection, channel };
}

export async function consumeFromQueue(channel: any, queue: string) {
  const messages: string[] = [];

  // Ensure queue exists
  await channel.declareQueue(queue, { durable: true });

  // Consume messages from the queue
  await channel.consume(queue, (msg: ConsumeMessage | null) => {
    if (msg) {
      const messageContent = new TextDecoder().decode(msg.body);
      console.log("Received:", messageContent);
      messages.push(messageContent);

      // Acknowledge the message after processing
      channel.ack(msg);
    }
  });

  return messages;
}
