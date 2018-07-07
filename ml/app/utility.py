# -*- coding: utf-8 -*-
import MeCab


class Doc2Query:
    def __init__(self):
        None

    def to_query(self, doc):
        qs = []

        # special word
        phrases = doc.split('"')
        np = len(phrases)
        if (np % 2 == 1) & (np >= 3):
            for i in range(np):
                if i % 2 == 1:
                    qs.append(phrases[i])

        # mecab
        mecab = MeCab.Tagger("-Ochasen")
        words = mecab.parse(doc)
        for row in words.split('\n'):
            word = row.split('\t')[0]
            if word == 'EOS':
                break
            else:
                pos = row.split('\t')[3].split('-')[0]
                if pos == '名詞':
                    qs.append(word)

        return qs
