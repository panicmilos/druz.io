import pika, sys, os
from dotenv import load_dotenv
import json

import smtplib
from email.message import EmailMessage

from jinja2 import Environment, PackageLoader, select_autoescape

load_dotenv()

##################### AMQP #####################

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

def callback(ch, method, properties, body):
  print(" [x] Received %r" % body)

##################### EMAIL #####################

class SMTPParams:
  def __init__(self, host, port, username, password):
    self.host = host
    self.port = port
    self.username = username
    self.password = password


class DefaultSMTPParams(SMTPParams):
  def __init__(self):
    super(DefaultSMTPParams, self).__init__(
      os.environ.get("SMTP_HOST"),
      os.environ.get("SMTP_PORT"),
      os.environ.get("SMTP_USERNAME"),
      os.environ.get("SMTP_PASSWORD")
    )

class Email:
  def __init__(self, **kwargs):
    self.subject = kwargs['subject']
    self.sender = kwargs['sender']
    self.to = kwargs['to']
    self.message = kwargs['message']
      
class EmailSender:
  def __init__(self, smtpParams = DefaultSMTPParams()):
    self._server = smtplib.SMTP(smtpParams.host, smtpParams.port)
    self._server.login(smtpParams.username, smtpParams.password)

  def send(self, email):
    emailMessage = EmailMessage()
    
    emailMessage['Subject'] = email.subject
    emailMessage['From'] = email.sender
    emailMessage['To'] = email.to
    emailMessage.set_content(email.message, subtype='html')

    self._server.send_message(emailMessage)

  def __del__(self):
    self._server.quit()

##################### TEMPALTES #####################

class TemplateManager:
  def __init__(self, package_name = 'main'):
    self._env = Environment(
      loader=PackageLoader(package_name),
      autoescape=select_autoescape()
    )

  def get_template(self, template_name):
    return self._env.get_template(template_name)

  def compute_tempalte(self, template_name, params = {}):
    template = self.get_template(template_name)

    return template.render(params)

def main():
  templateManager = TemplateManager()
  print(templateManager.compute_tempalte('test_template.html', { 'name': 'Milos' }))


  # email = Email(
  # subject = 'foo',
  # sender = 'panic.milos99@gmail.com',
  # to = 'panic.milos99@gmail.com',
  # message = '<font color="red">red color text</font>')

  # emailSender = EmailSender()
  # emailSender.send(email)

  # consumer = AMQPConsumer('emails')
  # consumer.on_data(callback)
  # consumer.consume()


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
