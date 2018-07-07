# -*- coding: utf-8 -*-
import pandas
from gensim.models import word2vec
import json
from collections import defaultdict
from app import model


class Controller():
    def __init__(self):
        None

    def set_db(self, db: model.Db):
        self.db = db

    def calc_all(self):
        print('-> summary')
        self.summary_calc()
        print('-> word')
        self.word_calc()

    def summary_get(self) -> str:
        return json.dumps({
            'total': self.summary_total,
            'sums': self.summary_sums.to_dict(),
            'corr': self.summary_corr.to_dict(),
            'social_sums': self.summary_social_sums.to_dict(),
            'social_corr': self.summary_social_corr.to_dict(),
        })

    def summary_calc(self):
        # data frame
        sql = self.db.sql_table_all('user_stats')
        df_stats = pandas.read_sql(sql, self.db.conn)

        sql = self.db.sql_table_all('users')
        df_users = pandas.read_sql(sql, self.db.conn)

        df = pandas.merge(
            df_users[[
                'id', 'qiita_id', 'name', 'description', 'mail', 'link'
                ]],
            df_stats[[
                'user_id', 'items', 'contributions', 'followers', 'followees'
                ]],
            left_on='id', right_on='user_id'
        )

        # data frame bool
        df_b_users = df[
            ['name', 'description', 'mail', 'link']
        ] != ''
        df_b_stas = df[
            ['items', 'contributions', 'followers', 'followees']
        ] > 0
        df_b = df_b_users.join(df_b_stas)

        # social stats
        sql = self.db.sql_table_all('user_social_links')
        df_social = pandas.read_sql(sql, self.db.conn)

        social_dummies = pandas.get_dummies(df_social['service_id'])
        social_dummies.columns = ['github', 'twitter', 'facebook', 'linkedin', 'google']
        df_social = df_social[['user_id']].join(social_dummies)
        df_social_stats = df_social.groupby('user_id').max()

        # unsafe...
        self.summary_total = len(df)
        self.summary_sums = df_b.sum()
        self.summary_corr = df_b.corr()
        self.summary_social_sums = df_social_stats.sum()
        self.summary_social_corr = df_social_stats.corr()

    def word_get(self) -> str:
        return json.dumps({
            'uniques': self.word_uniques,
        })

    def word_most_similar_get(self, word: str) -> str:
        wc = model.WordCalculator(self.word_model, self.word_uniques)
        sims = wc.most_similar(word)

        return json.dumps({
            'most_similar': sims,
        })

    def word_similarity_get(self, query: str, taregt: str) -> str:
        wc = model.WordCalculator(self.word_model, self.word_uniques)
        sim = wc.similarity(query, taregt)

        return json.dumps({
            'similarity': sim,
        })
    
    def word_similarity_many_get(self, queries, targets) -> str:
        wc = model.WordCalculator(self.word_model, self.word_uniques)
        queries = wc.preprocess_many(queries)

        sim = wc.similarity_many(queries, targets)

        return json.dumps({
            'similarity': sim,
            'queries': queries,
        })

    def word_similarity_users_get(self, queries, max_num) -> str:
        # calculator
        wc = model.WordCalculator(self.word_model, self.word_uniques)
        queries = wc.preprocess_many(queries)

        # output
        similars = [{'user_id': -1, 'score': -2} for i in range(max_num)]

        # calc all
        for user_id, words in self.word_lang_map.items():
            # ith user
            score = wc.similarity_many(queries, words)

            # replace
            if score > similars[max_num-1]['score']:
                similars.pop(max_num-1)
                similars.append({'user_id': user_id, 'score': score})
                similars = sorted(similars, key=lambda s: s['score'], reverse=True)

        return json.dumps({
            'similars': similars,
            'queries': queries
        })
    
    def word_calc(self):
        print('word_calc: df')
        # data frame
        sql = self.db.sql_table_all('users')
        df_users = pandas.read_sql(sql, self.db.conn)
        sql = self.db.sql_table_all('user_stats')
        df_stats = pandas.read_sql(sql, self.db.conn)
        sql = self.db.sql_table_all('user_language_stats')
        df_langs = pandas.read_sql(sql, self.db.conn)

        def tag(series, name):
            return name+':on' if series[name] else name+':off'
        
        # data reconstruct
        print('word_calc: rc')
        lang_map = defaultdict(list)
        for i, series in df_users.iterrows():
            user_id = int(series['id'])
            lang_map[user_id].append(str(series['qiita_id']))
            lang_map[user_id].append(tag(series, 'name'))
            lang_map[user_id].append(tag(series, 'description'))
            lang_map[user_id].append(tag(series, 'mail'))
            lang_map[user_id].append(tag(series, 'link'))
            lang_map[user_id].append(tag(series, 'organization'))
            lang_map[user_id].append(tag(series, 'place'))
            lang_map[user_id].append(tag(series, 'qiita_organization'))
            lang_map[user_id].append(tag(series, 'ban'))
        for i, series in df_stats.iterrows():
            user_id = int(series['id'])
            lang_map[user_id].append(tag(series, 'items'))
            lang_map[user_id].append(tag(series, 'contributions'))
            lang_map[user_id].append(tag(series, 'followers'))
            lang_map[user_id].append(tag(series, 'followees'))
        for i, series in df_langs.iterrows():
            user_id = int(series['user_id'])
            lang_map[user_id].append(str(series['name']))
        
        # make corpus
        print('word_calc: mc')
        word_list = []
        for k, v in lang_map.items():
            word_list.extend(v)
        word_set = set(word_list)

        # word2vec
        print('word_calc: wv')
        model = word2vec.Word2Vec([word_list], size=200, min_count=1, window=5, iter=200)

        # unsafe...
        self.word_lang_map = lang_map
        self.word_uniques = word_set
        self.word_model = model
