import numpy as np
import matplotlib.pyplot as plt

# 设置随机种子以确保结果的可重复性
np.random.seed(42)

# 生成数据集
n = 100  # 数据点的数量
x = np.linspace(0, 10, n)  # 生成等间隔的x值
true_slope = 2.5
y_true = true_slope * x + np.random.normal(0, 1, n)  # 真实的y值，包含一些随机噪声

# 生成预测数据集
# 为简化，这里使用与真实数据相同的模型，但通常预测值会来自于某个模型的输出
y_pred = true_slope * x + np.random.normal(0, 1, n)  # 预测的y值，也包含随机噪声

# 计算MSE
mse = np.mean((y_true - y_pred) ** 2)

# 绘制真实值与预测值
plt.figure(figsize=(10, 6))
plt.scatter(x, y_true, label='True Values', alpha=0.6)  # 真实值
plt.plot(x, y_pred, color='red', label='Predictions')  # 预测值
plt.title(f'Mean Squared Error: {mse:.2f}')
plt.xlabel('X')
plt.ylabel('Y')
plt.legend()
plt.grid(True)
plt.show()

# 输出数据集
for xi, yi, yhat in zip(x, y_true, y_pred):
    print(f'X: {xi:.2f}, Y_True: {yi:.2f}, Y_Pred: {yhat:.2f}')
