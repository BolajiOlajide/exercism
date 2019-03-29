import re


def is_pangram(sentence):
    hasAllLetters = False
    alphabets = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
    stripped_sentence = re.sub('[_ ]', '', sentence.lower())
    for letter in alphabets:
        if letter in stripped_sentence:
            hasAllLetters = True
        else:
            hasAllLetters = False
            break
    return hasAllLetters
