# -*- coding: utf-8 -*-
import os

http = {
    'port': os.environ['ML_PORT'],
}

db = {
    'host': os.environ['DB_HOST'],
    'port': os.environ['DB_PORT'],
    'user': os.environ['DB_USER'],
    'password': os.environ['DB_PASSWORD'],
    'database': os.environ['DB_NAME'],
}
