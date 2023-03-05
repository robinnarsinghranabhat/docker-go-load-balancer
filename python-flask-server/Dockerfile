FROM python:3.9-slim

WORKDIR /app

COPY . .

RUN apt-get update -y
RUN apt-get install -y iputils-ping
RUN pip install -r requirements.txt

CMD ["python", "./server.py"]