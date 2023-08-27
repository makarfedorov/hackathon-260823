from database import DatabaseHelper
db_helper = DatabaseHelper()

def create_tables():
    db_helper.create_table()
    db_helper.create_blocked_table()

if __name__ == '__main__':
    create_tables()
