const { insertDoc } = require("./insert");
const { con, Queue } = require("./connect");

const consumer = async () => {
  try {
    const conn = await con();

    const ch = await conn.createChannel();

    // durable true verdiğinde bir işlem fail olursa bunu silmez tutar.
    await ch.assertQueue(Queue, { durable: true });

    await ch.consume(Queue, async (msg) => {
      const data = JSON.parse(msg.content.toString());

      await insertDoc("product", data.id, data);
      ch.ack(msg);
    });

  } catch (err) {
    console.error(err);
  }
};

consumer();
