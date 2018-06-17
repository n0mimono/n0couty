# -*- coding: utf-8 -*-
import mysql.connector
from app import config
import numpy


class Db():
    def __init__(self):
        None
    
    def connect(self):
        self.conn = mysql.connector.connect(**config.db)

    def sql_table_all(self, table_name: str) -> str:
        return "select * from {table_name}".format(table_name=table_name)


class WordCalculator():
    def __init__(self, model, uniques):
        self.model = model
        self.uniques = uniques
    
    def most_similar(self, word):
        if not (word in self.uniques):
            return []
        
        words = self.model.most_similar(positive=word)
        return [{'word': words[i][0], 'score':words[i][1]} for i in range(10)]

    def similarity(self, query, target):
        if not (query in self.uniques):
            return -1
        if not (target in self.uniques):
            return -1
        
        return self.model.similarity(query, target)

    def preprocess_many(self, queries):
        return list(filter(lambda q: q in self.uniques, queries))
    
    def similarity_many(self, queries, targets):
        n = len(queries)

        score = 0
        for i in range(n):
            q = queries[i]
            sims = numpy.array([self.similarity(q, t) for t in targets])
            idx = numpy.argmax(sims)

            score = score + sims[idx]
        
        return score / n
