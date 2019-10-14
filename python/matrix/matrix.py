class Matrix(object):
    def __init__(self, matrix_string):
        self._matrix = map(
            lambda x: x.split(" "),
            matrix_string.split("\n")
        )
        self.matrix = list(self._matrix)

    def row(self, index):
        return list(map(int, self.matrix[index - 1]))

    def column(self, index):
        cols = list(map(lambda x: x[index - 1], self.matrix))
        return list(map(int, cols))
