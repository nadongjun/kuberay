import ray
import os
import time

ray.init()

@ray.remote
class Counter:
    def __init__(self):
        # Used to verify runtimeEnv
        self.name = os.getenv("counter_name")
        assert self.name == "test_counter"
        self.counter = 0

    def inc(self):
        self.counter += 1

    def get_counter(self):
        return "{} got {}".format(self.name, self.counter)

counter = Counter.remote()

for _ in range(60):
    ray.get(counter.inc.remote())
    print(ray.get(counter.get_counter.remote()))
    time.sleep(1)
