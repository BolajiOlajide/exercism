def is_armstrong(number):
    _number = str(number)
    sum = 0
    number_length = len(_number)
    for digit in _number:
        _digit = int(digit)
        sum += pow(_digit, number_length)
    return True if sum == number else False
