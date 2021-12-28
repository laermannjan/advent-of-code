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