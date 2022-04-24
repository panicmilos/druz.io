import os
from dotenv import load_dotenv
import smtplib
from email.message import EmailMessage


load_dotenv()


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
