import matplotlib.pyplot as plt
import numpy as np

data_shiviz = [
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

data_disviz = [
    [733, 772.4],
    [1250, 749.7],
    [1676, 883.6],
    [2692, 897],
    [2468, 989],
    [2955, 1199.6],
    [4476, 1397.9],
    [5426, 1549],
    [6435, 2195.6],
    [5257, 1441.4],
    [7836, 1903.5],
    [6272, 1831.1],
    [7272, 1513.2],
    [9670, 1968.4],
    [8874, 2005.5],
    [11034, 2268.8],
    [11004, 2997.2],
    [11819, 2410.6],
    [11231, 2521.3],
    [12873, 2459.1]
]

data_disviz_reload = [
    [733, 633.4],
    [1250, 692.8],
    [1676, 709],
    [2692, 769],
    [2468, 773.4],
    [2955, 806.8],
    [4476, 803.2],
    [5426, 926.9],
    [6435, 971.7],
    [5257, 835.8],
    [7836, 1030.3],
    [6272, 1006],
    [7272, 1276.4],
    [9670, 1437.2],
    [8874, 1231.5],
    [11034, 1301.1],
    [11004, 1330.5],
    [11819, 1139.1],
    [11231, 1478.9],
    [12873, 1287.3]
]

plt.figure(figsize=(10, 6))

for dataset, label in [
    (data_shiviz, "Shiviz"),
    (data_disviz, "DisViz (first load)"),
    (data_disviz_reload, "DisViz (reload)")
]:
    # sort by x
    ds = sorted(dataset, key=lambda item: item[0])
    x = np.array([pt[0] for pt in ds])
    y = np.array([pt[1] / 1000 for pt in ds])

    # fit line
    m, b = np.polyfit(x, y, 1)
    y_fit = m * x + b
    # format slope and intercept in scientific notation
    m_sci = f"{m:.2e}"
    b_sci = f"{b:.1e}"

    # plot original data
    label = label + f", fit: y = {m_sci}x + {b_sci}"
    line, = plt.plot(x, y, marker='o', linestyle='-', label=label)
    color = line.get_color()

    # plot dashed fit in lighter shade
    plt.plot(
        x, y_fit,
        linestyle='--',
        color=color,
        alpha=0.5,
        # label=f"{label} fit: y = {m_r}x + {b_r}"
    )

plt.xlabel("Number of Log Events")
plt.ylabel("Visualization Latency (seconds)")
plt.title("Performance Comparison between ShiViz and DisViz")
plt.grid(True)
plt.legend()
plt.tight_layout()
plt.show()