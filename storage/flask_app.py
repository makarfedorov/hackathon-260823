from flask import Flask, request
from cryptography.fernet import Fernet
import sqlite3
import threading


app = Flask(__name__)

with open('secret_key.txt', 'rb') as key_file:
    SECRET_KEY = key_file.read()
#SECRET_KEY = Fernet.generate_key()
cipher_suite = Fernet(SECRET_KEY)




class DatabaseHelper:
    def __init__(self):
        self._lock = threading.Lock()
    
    
    
    def create_table(self):
        conn = sqlite3.connect('user_data.db')
        cursor = conn.cursor()
        cursor.execute('DROP TABLE IF EXISTS encrypted_ids')
        cursor.execute('''
            CREATE TABLE encrypted_ids (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                encrypted_id TEXT,
                topics BLOB,
                password TEXT
            )
        ''')
        conn.commit()
        conn.close()

    def store_encrypted_info(self, encrypted_id, name, password):
        with self._lock:
            conn = sqlite3.connect('user_data.db')
            cursor = conn.cursor()
            cursor.execute('INSERT INTO encrypted_ids (encrypted_id, topics, password) VALUES (?, ?, ?)',
                           (encrypted_id, name, password))

            conn.commit()
            conn.close()

db_helper = DatabaseHelper()
db_helper.create_table()
@app.route('/store_id', methods=['POST'])
def store_id():
    data = request.json
    user_id = data.get('user_id')
    topics = data.get('topics')
    password = data.get('password')


    encrypted_id = cipher_suite.encrypt(str(user_id).encode())
    encrypted_password = cipher_suite.encrypt(str(password).encode())
    db_helper.store_encrypted_info(encrypted_id, topics, encrypted_password)

    return "Consultant information encrypted and stored successfully!"




if __name__ == '__main__':
    app.run(debug=True)
