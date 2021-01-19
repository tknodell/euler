import csv
from string import ascii_uppercase


def letterValues():
    scores = {}
    for i, c in enumerate(ascii_uppercase):
        scores[c] = i+1
    return scores


def calcNameScore(n):
    res = 0
    for letter in n:
        res += letterScores[letter]
    return res


listOfNames = []
with open('p022_names.txt', 'r') as f:
    for row in csv.reader(f):
        listOfNames.append(row[0])

listOfNames.sort()

letterScores = letterValues()

totalSum = 0
for index, name in enumerate(listOfNames):
    score = calcNameScore(name)
    print(name, index+1, score)
    totalSum += score*(index+1)

print("---")
print(totalSum)
