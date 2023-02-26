

class SlidingMedian:
    def __init__(self):
        self.maxheap = []
        self.minheap = []

    def add(num):
        self.maxheap.push(num)
        self.minheap.push(self.maxheap.pop())

        if self.minheap.size > self.maxheap.size:
            self.maxheap.push(self.minheap.pop())

    @property
    def median(self):
        if self.maxheap.size > self.minheap.size:
            return self.maxheap.top
        else:
            return (self.minheap.top + self.maxheap.top) / 2
