import datetime

year = 1901
sundayCount = 0

while year <= 2000:
    month = 1
    while month <= 12:
        targetDate = datetime.datetime(year, month, 1)

        # 6 is Sunday
        if targetDate.weekday() == 6:
            sundayCount = sundayCount + 1
        month = month + 1
    year = year + 1

print(sundayCount)
