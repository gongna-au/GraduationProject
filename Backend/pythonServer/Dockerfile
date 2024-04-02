FROM python:3.8-slim
WORKDIR /app
COPY . .
RUN pip install paho-mqtt
CMD ["python", "server.py"]
