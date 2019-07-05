from random import choice, random, seed


class Robot(object):
    alphabets = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']

    def __init__(self):
        self.name = Robot._generate_code()

    @classmethod
    def _generate_code(cls):
        seed()
        prefix_one = choice(Robot.alphabets).upper()
        prefix_two = choice(Robot.alphabets).upper()
        code = round(random() * 1000)

        return f'{prefix_one}{prefix_two}{code}'

    def reset(self):
        self.name = Robot._generate_code()
