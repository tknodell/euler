
def isPalindrome(n):
    n = str(n)
    reverse = n[::-1]
    return n == reverse


def toBinary(n):
    return bin(n).replace("0b", "")


i = 1
sum = 0
while i < 1000000:
    if isPalindrome(i) and isPalindrome(toBinary(i)):
        sum = sum+i
    i = i+1

print(sum)
