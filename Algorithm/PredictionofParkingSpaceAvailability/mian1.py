# 导入必要的库
import numpy as np
import pandas as pd
from statsmodels.tsa.arima.model import ARIMA
import matplotlib.pyplot as plt

# 设置随机种子以确保结果的可重复性
np.random.seed(42)

# 自动生成365个数据点作为示例数据，代表过去一年内每天的最高停车位占用率
data = np.random.randint(65, 95, size=365)

# 创建日期范围
dates = pd.date_range(start='2023-01-01', end='2023-12-31')
df = pd.Series(data, index=dates)

# 使用ARIMA模型进行拟合，这里使用(1,1,1)作为示例参数
model = ARIMA(df, order=(1,1,1))
model_fit = model.fit()

# 进行未来7天的预测
forecast = model_fit.forecast(steps=7)

# 打印预测结果
print("未来7天的预测停车位最高占用率：")
print(forecast)

# 绘制数据和预测结果的图表
plt.figure(figsize=(10,6))
plt.plot(df, label='Actual Data')
plt.plot(forecast, label='Forecast', color='red')
plt.xlabel('Date')
plt.ylabel('Parking Occupancy Rate (%)')
plt.title('Parking Occupancy Rate Forecast')
plt.legend()
plt.show()
