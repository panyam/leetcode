
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


def quick_select(nums, k, L, R, cmpfunc=None): 
    import random
    pivot = nums[random.randint(L, R)]
    # move data around that our array looks like:
    # items < pivot, items == pivot and items > pivot
    l,r = L,R
    if cmpfunc is None:
        cmpfunc = lambda a,b: a - b
    while l < r:
        if cmpfunc(nums[l], pivot) <= 0:
            l += 1
        else: #  nums[l] > pivot:
            nums[l],nums[r] = nums[r],nums[l]
            r -= 1
    
    # Here l == r
    # now move pivot items to end of low section
    LowStart,LowEnd = L, l - 1
    while LowStart < LowEnd:
        if cmpfunc(nums[LowStart], pivot) == 0:
            nums[LowStart],nums[LowEnd] = nums[LowEnd],nums[LowStart]
            LowEnd -= 1
        else:   # < pivot
            LowStart += 1

    print("AfterLow: ", pivot, nums)
    if k <= LowEnd:
        return quick_select(nums, k, L, LowEnd, cmpfunc)
    elif k >= r:
        return quick_select(nums, k, LowEnd + 1, R, cmpfunc)
    else:
        return nums[LowEnd + 1]
