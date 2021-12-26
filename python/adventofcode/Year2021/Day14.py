from collections import defaultdict
from typing import Callable


class memoized_property:
    def __init__(self, factory: Callable):
        self._factory = factory
        self._value = None
        self._cached_instance_hash = None

    def __get__(self, instance: object, owner: type):
        instance_hash = instance.__hash__()
        if instance_hash != self._cached_instance_hash:
            self._cached_instance_hash = instance_hash
            self._value = self._factory(instance)
        return self._value


class Polymer:
    def __init__(self, polymer_template: str, rules: dict[str, str]):
        self._polymer_template = polymer_template
        self._rules = rules

        self._pair_frequencies: defaultdict[str, int] = defaultdict(int)
        for i in range(len(polymer_template) - 1):
            pair = polymer_template[i:i+2]
            assert pair in rules
            self._pair_frequencies[pair] += 1

    def step(self, steps: int):
        for _ in range(steps):
            new_pair_frequencies: defaultdict[str, int] = defaultdict(int)
            for pair, freq in self._pair_frequencies.items():
                new_pair_frequencies[f"{pair[0]}{self._rules[pair]}"] += freq
                new_pair_frequencies[f"{self._rules[pair]}{pair[1]}"] += freq
            self._pair_frequencies = new_pair_frequencies

    @memoized_property
    def letter_frequencies(self) -> dict[str, int]:
        letter_frequencies: defaultdict[str, int] = defaultdict(int)
        letter_frequencies[self._polymer_template[-1]] = 1

        for pair, freq in self._pair_frequencies.items():
            letter_frequencies[pair[0]] += freq
        return letter_frequencies

    @property
    def most_common_letter(self):
        return max(self.letter_frequencies.values())

    @property
    def least_common_letter(self):
        return min(self.letter_frequencies.values())

    def __hash__(self):
        return hash(tuple(self._pair_frequencies.items()))

    def __repr__(self):
        return self._polymer_template

    @classmethod
    def from_file(cls, filename: str) -> "Polymer":
        with open(filename, 'r') as f:
            data = f.read().splitlines()
        polymer_template = data[0]
        rules = {
            pair: insertion for pair, insertion in (rule.split(' -> ') for rule in data[2:])
        }
        return cls(polymer_template, rules)


def day14a(filename: str) -> int:
    polymer = Polymer.from_file(filename)
    polymer.step(10)
    return polymer.most_common_letter - polymer.least_common_letter


def day14b(filename: str) -> int:
    polymer = Polymer.from_file(filename)
    polymer.step(40)
    return polymer.most_common_letter - polymer.least_common_letter

