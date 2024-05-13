import pika

connection_parameters = pika.ConnectionParameters('localhost')

connection = pika.BlockingConnection(connection_parameters)

channel = connection.channel()

# Deklarasikan pertukaran (exchange)
channel.exchange_declare(exchange='pubsub', exchange_type='fanout')

message = "Hello, I want to broadcast this message!"

# Kirim pesan ke pertukaran dengan routing key kosong
channel.basic_publish(exchange='pubsub', routing_key='', body=message)

print(f"Sent message: {message}")

connection.close()
