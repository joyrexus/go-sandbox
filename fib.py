
def fib(n):
    if n < 2: return 1 
    return fib(n-1) + fib(n-2)

if __name__ == '__main__':
    for i in range(10):
        print i, fib(i)

