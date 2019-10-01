def slices(series, length):
    string_length = len(series)
    if (string_length < length) or (length < 1):
        raise ValueError(".+")

    result = []

    for index, char in enumerate(series):
        upper_bound = index + length
        if upper_bound > string_length:
            break
        string_range = range(index, upper_bound)
        new_string = ""
        for i in string_range:
            new_string += series[i]
        result.append(new_string)
    return result
