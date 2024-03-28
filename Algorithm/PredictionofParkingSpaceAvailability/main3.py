import numpy as np
import pandas as pd
from sklearn.ensemble import RandomForestRegressor
from sklearn.model_selection import train_test_split
import matplotlib.pyplot as plt

# 生成模拟数据
np.random.seed(42)
dates = pd.date_range(start='2023-01-01', end='2023-12-31', freq='D')
data = np.random.normal(100, 20, size=len(dates)) + np.sin(np.linspace(0, 3.14*2, len(dates))) * 20
df = pd.DataFrame(data, index=dates, columns=['Demand'])

# 添加时间特征
df['Month'] = df.index.month
df['DayOfWeek'] = df.index.dayofweek

# 将数据分为训练集和测试集
X = df[['Month', 'DayOfWeek']]
y = df['Demand']
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# 创建随机森林模型并训练
rf = RandomForestRegressor(n_estimators=100, random_state=42)
rf.fit(X_train, y_train)

# 进行预测
y_pred = rf.predict(X_test)

# 评估模型
plt.figure(figsize=(10, 6))
plt.plot(y_test.index, y_test, label='Actual Demand')
plt.plot(y_test.index, y_pred, label='Predicted Demand', color='red')
plt.xlabel('Date')
plt.ylabel('Parking Spot Demand')
plt.title('Parking Spot Demand Prediction')
plt.legend()
plt.show()
