const amqplib = require("amqplib");
const Queue = "insertProduct";
const con = async () => {
  let con;
  try {
    con = await amqplib.connect("amqp://localhost");

    console.log("Connected to RabbitMQ");
  } catch (err) {
    console.error(err);
  }

  return con;
};

module.exports = {
    con,
    Queue,
};