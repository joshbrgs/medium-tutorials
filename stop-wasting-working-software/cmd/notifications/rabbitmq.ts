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

export async function consumeFromQueue(channel: Channel, queueName: string): Promise<string[]> {
  const messages: string[] = [];

  await channel.consume(
    { queue: queueName },
    async (args, props, data) => {
      console.log(JSON.stringify(args));
      console.log(JSON.stringify(props));
      messages.push(new TextDecoder().decode(data));
      await channel.ack({ deliveryTag: args.deliveryTag });
    },
  );

  // Consider adding a timeout or a mechanism to stop consuming after some time
  // This is just a placeholder, adjust it based on your needs
  await new Promise((resolve) => setTimeout(resolve, 1000));

  return messages;
}
