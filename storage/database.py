import threading
from cryptography.fernet import Fernet
import sqlite3

    
class DatabaseHelper:
    def __init__(self):
        self._lock = threading.Lock()
    
    def create_blocked_table(self):
        conn = sqlite3.connect('user_data.db')
        cursor = conn.cursor()
        cursor.execute('DROP TABLE IF EXISTS blocked_ids')
        cursor.execute('''
            CREATE TABLE blocked_ids (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                encrypted_id TEXT
            )
        ''')
        conn.commit()
        conn.close()


    def store_blocked_id(self, encrypted_id):
        with self._lock:
            conn = sqlite3.connect('user_data.db')
            cursor = conn.cursor()
            cursor.execute('INSERT INTO blocked_ids (encrypted_id) VALUES (?)', (encrypted_id,))
            conn.commit()
            conn.close()
    
    def retrieve_blocked_ids(self):
        with self._lock:
            conn = sqlite3.connect('user_data.db')
            cursor = conn.cursor()
            cursor.execute('SELECT encrypted_id FROM blocked_ids')
            encrypted_blocked_ids = cursor.fetchall()
            decrypted_blocked_ids = [cipher_suite.decrypt(encrypted_id[0]).decode() for encrypted_id in encrypted_blocked_ids]
            conn.close()
            return decrypted_blocked_ids

    

    def create_table(self):
        conn = sqlite3.connect('user_data.db')
        cursor = conn.cursor()
        cursor.execute('DROP TABLE IF EXISTS encrypted_ids')
        cursor.execute('''
            CREATE TABLE encrypted_ids (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                encrypted_id TEXT,
                topics Text,
                password TEXT
            )
        ''')
        conn.commit()
        conn.close()

    def store_encrypted_info(self, encrypted_id, topics, password):
        with self._lock:
            conn = sqlite3.connect('user_data.db')
            cursor = conn.cursor()
            cursor.execute('INSERT INTO encrypted_ids (encrypted_id, topics, password) VALUES (?, ?, ?)',
                           (encrypted_id, topics, password))

            conn.commit()
            conn.close()
    

    def retrieve_data(self):
        with self._lock:
            conn = sqlite3.connect('user_data.db')
            cursor = conn.cursor()
            cursor.execute('SELECT encrypted_id, topics, password FROM encrypted_ids')
            encrypted_data = cursor.fetchall()

            decrypted_data = []
            for encrypted_id, topics, str_encrypted_password in encrypted_data:
                decrypted_id = cipher_suite.decrypt(encrypted_id).decode()
                decrypted_topics = json.loads(topics)
                decrypted_password = cipher_suite.decrypt(str_encrypted_password).decode()

                decrypted_data.append({'user_id': decrypted_id, 'topics': decrypted_topics, 'password': decrypted_password})

            conn.close()
            return decrypted_data
