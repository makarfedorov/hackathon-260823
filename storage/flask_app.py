from flask import Flask, request
from cryptography.fernet import Fernet
import sqlite3
import threading
import json
from database import DatabaseHelper

app = Flask(__name__)

with open('secret_key.txt', 'rb') as key_file:
    SECRET_KEY = key_file.read()
cipher_suite = Fernet(SECRET_KEY)




db_helper = DatabaseHelper()
#db_helper.create_table()
#db_helper.create_blocked_table()

@app.route('/store_id', methods=['POST'])
def store_id():
    data = request.json
    user_id = data.get('user_id')
    topics = json.dumps(data.get('topics'))
    password = data.get('password')
    encrypted_id = cipher_suite.encrypt(str(user_id).encode())
    encrypted_password = cipher_suite.encrypt(str(password).encode())
    
    
    db_helper.store_encrypted_info(encrypted_id, topics, encrypted_password)

    return "Consultant information encrypted and stored successfully!"


@app.route('/get_data', methods=['GET'])
def get_data():
    decrypted_data = db_helper.retrieve_data()
    return {'decrypted_data': decrypted_data}



@app.route('/block_user', methods=['POST'])
def block_user():
    data = request.json
    user_id = data.get('user_id')
    encrypted_id = cipher_suite.encrypt(str(user_id).encode())

    db_helper.store_blocked_id(encrypted_id)

    return "User ID blocked successfully!"

@app.route('/get_blocked_ids', methods=['GET'])
def get_blocked_ids():
    blocked_ids = db_helper.retrieve_blocked_ids()
    return {'blocked_ids': blocked_ids}

if __name__ == '__main__':
    app.run(debug=True)
