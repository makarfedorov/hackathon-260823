from cryptography.fernet import Fernet

def generate_key():
    secret_key = Fernet.generate_key()
    with open('secret_key.txt', 'wb') as key_file:
        key_file.write(secret_key)


if __name__ == '__main__':
    generate_key()
