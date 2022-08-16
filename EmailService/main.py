import sys, os
import json
from amqp_consumer import AMQPConsumer
from email_sender import Email, EmailSender
from template_manager import TemplateManager


templateManager = TemplateManager()


def callback(ch, method, properties, body):
  data = json.loads(body)

  email = Email(
    subject = data['subject'],
    sender = data['from'],
    to = data['to'],
    message = templateManager.compute_tempalte(data['message']['template'], data['message']['params'])
  )

  emailSender = EmailSender()
  emailSender.send(email)


def main():
  consumer = AMQPConsumer('emails')
  consumer.on_data(callback)
  consumer.consume()


if __name__ == '__main__':
  try:
    main()
  except KeyboardInterrupt:
    print('Interrupted')
    try:
      sys.exit(0)
    except SystemExit:
      os._exit(0)
