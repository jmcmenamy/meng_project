import matplotlib.pyplot as plt

data = [
    [733, 540.3],
    [1250, 727.7],
    [1676, 1085.4],
    [2692, 1439.9],
    [2468, 1406.9],
    [2955, 2239],
    [4476, 2488.4],
    [5426, 3355],
    [6435, 3994.8],
    [5257, 3216.6],
    [7836, 4882.1],
    [6272, 3494.3],
    [7272, 4023.1],
    [9670, 5691.3],
    [8874, 4706],
    [11034, 5670.4],
    [11004, 5516.3],
    [11819, 6758.6],
    [11231, 6214.6],
    [12873, 8031.6]
]

data.sort(key = lambda item: item[0])

x = [point[0] for point in data]
y = [point[1] for point in data]
y = [val/1000 for val in y]

plt.figure(figsize=(10, 6))
plt.plot(x, y, marker='o')
plt.xlabel("Number of Log Events")
plt.ylabel("Visualization Latency (seconds)")
plt.title("Correlation of Log Events With Latency to Create Visualization")
plt.grid(True)
plt.tight_layout()
plt.show()