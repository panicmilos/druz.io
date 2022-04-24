import pika, os
from dotenv import load_dotenv


load_dotenv()


class AMQPConnectionParams:
  def __init__(self, host, port, username, password):
    self.host = host
    self.port = port
    self.username = username
    self.password = password


class DefaultAMQPConnectionParams(AMQPConnectionParams):
  def __init__(self):
    super(DefaultAMQPConnectionParams, self).__init__(
      os.environ.get("AMQP_HOST"),
      os.environ.get("AMQP_PORT"),
      os.environ.get("AMQP_USERNAME"),
      os.environ.get("AMQP_PASSWORD")
    )


class AMQPConsumer:
  def __init__(self, queue, connectionParams = DefaultAMQPConnectionParams()):
    self._connection = pika.BlockingConnection(pika.ConnectionParameters(
      host = connectionParams.host,
      port = connectionParams.port,
      credentials=pika.PlainCredentials(connectionParams.username, connectionParams.password)
    ))
    self._channel = self._connection.channel()
    self._channel.queue_declare(queue=queue)
    self._callback = None
    self._channel.basic_consume(queue=queue, on_message_callback=self.callback, auto_ack=True)

  def callback(self, ch, method, properties, body):
    self._callback(ch, method, properties, body)

  def on_data(self, callback):
    self._callback = callback

  def consume(self):
    self._channel.start_consuming()
