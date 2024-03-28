import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from statsmodels.tsa.statespace.sarimax import SARIMAX

# 创建模拟数据：模拟一年内每天的停车位需求
np.random.seed(42)
data = np.random.normal(100, 10, 365) + np.sin(np.linspace(0, 3.14*2, 365)) * 20
dates = pd.date_range(start='2023-01-01', end='2023-12-31')
df = pd.Series(data, index=dates)

# 拟合SARIMA模型：这里的参数需要根据实际数据调整
# 参数示例：(1, 1, 1) x (1, 1, 1, 12) 表示季节性周期为一年
model = SARIMAX(df, order=(1, 1, 1), seasonal_order=(1, 1, 1, 12))
model_fit = model.fit(disp=False)

# 进行预测：预测接下来30天的需求
forecast = model_fit.forecast(steps=30)

# 绘制原始数据和预测结果
plt.figure(figsize=(10, 6))
plt.plot(df, label='Actual Demand')
plt.plot(forecast, label='Forecasted Demand', color='red')
plt.xlabel('Date')
plt.ylabel('Parking Spot Demand')
plt.title('Parking Spot Demand Forecast')
plt.legend()
plt.show()
