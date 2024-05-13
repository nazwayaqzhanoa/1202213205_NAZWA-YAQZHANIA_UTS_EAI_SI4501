import pika

def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)

# Membuat koneksi dengan RabbitMQ
connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

# Mendeklarasikan exchange
channel.exchange_declare(exchange='pubsub', exchange_type='fanout')

# Mendeklarasikan antrian (queue) secara acak
result = channel.queue_declare(queue='', exclusive=True)
queue_name = result.method.queue

# Mengikat antrian ke exchange
channel.queue_bind(exchange='pubsub', queue=queue_name)

print(' [*] Waiting for messages. To exit press CTRL+C')

# Mendefinisikan callback untuk menerima pesan
channel.basic_consume(queue=queue_name, on_message_callback=callback, auto_ack=True)

# Memulai konsumsi pesan
channel.start_consuming()
