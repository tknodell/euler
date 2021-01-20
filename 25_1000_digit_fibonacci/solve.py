
n1,n2=1,1
count=2 # account for first two values of 1 & 1
nextValue=0

while len(str(nextValue))!=1000:
    nextValue = n1+n2
    n1=n2
    n2=nextValue
    count=count+1
    # print(nextValue,len(str(nextValue)),count)

print(count)
