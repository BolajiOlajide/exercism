import re


def abbreviate(sentence):
    pattern = r'\b[a-zA-Z]'
    non_string = r'[^A-Za-z -]'
    clean_sentence = re.sub(non_string, '', sentence)
    abbreviation = re.findall(pattern, clean_sentence)
    return ''.join(abbreviation).upper()
