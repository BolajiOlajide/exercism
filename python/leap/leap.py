def is_leap_year(year):
    try:
        _year = int(year)
        if (_year % 4 == 0) :
            if (_year % 100 == 0):
                if (_year % 400 == 0):
                    return True
                else:
                    return False
            else:
                return True
        else:
            return False
    except ValueError:
        raise Exception("Input is not a valid integer")
