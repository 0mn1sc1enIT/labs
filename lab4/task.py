def htower(N, a, b):
    if N == 1:
        print(N, a, b)
    else:
        htower(N-1, a, 6-a-b)
        print(N, a, b)
        htower(N-1, 6-a-b, b) 
n = int(input())
htower(n, 1, 3)