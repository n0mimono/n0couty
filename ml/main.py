# -*- coding: utf-8 -*-
from app import model
from app import controller
from app import utility
from app import config

from flask import Flask, request
from flask_cors import CORS
app = Flask(__name__)
CORS(app)

ctrl = controller.Controller()


@app.route("/")
def index():
    return "Hello, world!"


@app.route("/api/ml/summary", methods=['POST', 'GET'])
def summary():
    if request.method == 'POST':
        ctrl.summary_calc()
    return ctrl.summary_get()


@app.route("/api/ml/word", methods=['POST', 'GET'])
def word():
    if request.method == 'POST':
        ctrl.word_calc()
    return ctrl.word_get()


@app.route("/api/ml/word/most_similar")
def most_similar():
    word = request.args.get('word')
    return ctrl.word_most_similar_get(word)


@app.route("/api/ml/word/similarity")
def similarity():
    query = request.args.get('query')
    target = request.args.get('target')
    return ctrl.word_similarity_get(query, target)


@app.route("/api/ml/word/similarity_many")
def similarity_many():
    queries = request.args.get('queries').split(' ')
    targets = request.args.get('targets').split(' ')
    return ctrl.word_similarity_many_get(queries, targets)


@app.route("/api/ml/word/similarity_users", methods=['POST', 'GET'])
def similarity_users():
    if request.method == 'POST':
        doc = request.get_json().get('doc')
        queries = utility.Doc2Query().to_query(doc)
    else:
        queries = request.args.get('queries').split(' ')
    max_num = int(request.args.get('max'))
    return ctrl.word_similarity_users_get(queries, max_num)


if __name__ == "__main__":
    db = model.Db()
    db.connect()

    ctrl.set_db(db)
    ctrl.calc_all()
    
    app.run(**config.http)

