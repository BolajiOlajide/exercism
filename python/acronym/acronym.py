def abbreviate(words):
    abbr = ''
    splitted_word = words.split(' ')
    for word in splitted_word:
        abbr += ''.join([i[0] for i in word.split('-') if len(i) > 0])
    return abbr.upper()
