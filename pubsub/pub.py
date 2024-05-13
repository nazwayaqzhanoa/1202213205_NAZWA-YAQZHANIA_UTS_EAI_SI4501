import pika
import json

# Membuat koneksi dengan RabbitMQ
connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

# Mendeklarasikan exchange
channel.exchange_declare(exchange='logs', exchange_type='fanout')

# Data JSON yang akan dikirim
data = {
    "name": "John",
    "age": 30,
    "city": "New York"
}

# Mengonversi data menjadi format JSON
message = json.dumps(data)

# Mengirim pesan ke exchange
channel.basic_publish(exchange='logs', routing_key='', body=message)

print(" [x] Sent %r" % message)

# Menutup koneksi
connection.close()
