def rotation(n):
    rotations = set()
    for i in range(len(str(n))):
        rotations.add(int(str(n)[i:] + str(n)[:i]))
    return rotations


def isPrime(num):
    if num <= 1:
        return False
    for i in range(2, num):
        if num % i == 0:
            return False
    return True


i = 0
circleCount = 0
while i < 1000000:
    nums = rotation(i)

    isCircular = True
    for n in nums:
        if isPrime(n) == False:
            isCircular = False

    if isCircular:
        circleCount = circleCount+1
        print(i)
    i = i + 1

print("TOTAL")
print(circleCount)
