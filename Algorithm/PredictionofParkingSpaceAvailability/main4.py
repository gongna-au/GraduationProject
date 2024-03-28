import numpy as np
import pandas as pd
from sklearn.preprocessing import MinMaxScaler
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import LSTM, Dense
import matplotlib.pyplot as plt

# 生成模拟数据
np.random.seed(42)
dates = pd.date_range(start='2023-01-01', end='2023-12-31', freq='D')
data = np.random.normal(100, 20, size=len(dates)) + np.sin(np.linspace(0, 3.14*2, len(dates))) * 20
df = pd.DataFrame(data, index=dates, columns=['Demand'])

# 数据预处理：归一化
scaler = MinMaxScaler(feature_range=(0, 1))
scaled_data = scaler.fit_transform(df)

# 准备训练数据
def create_dataset(data, time_step=1):
    X, Y = [], []
    for i in range(len(data) - time_step - 1):
        a = data[i:(i + time_step), 0]
        X.append(a)
        Y.append(data[i + time_step, 0])
    return np.array(X), np.array(Y)

time_step = 10
X, y = create_dataset(scaled_data, time_step)
X = X.reshape(X.shape[0], X.shape[1], 1)

# 划分训练集和测试集
train_size = int(len(X) * 0.8)
test_size = len(X) - train_size
X_train, X_test = X[0:train_size], X[train_size:len(X)]
y_train, y_test = y[0:train_size], y[train_size:len(y)]

# 创建LSTM模型
model = Sequential([
    LSTM(50, activation='relu', input_shape=(time_step, 1)),
    Dense(1)
])
model.compile(optimizer='adam', loss='mean_squared_error')

# 训练模型
model.fit(X_train, y_train, epochs=100, batch_size=32, verbose=1)

# 预测
train_predict = model.predict(X_train)
test_predict = model.predict(X_test)

# 反归一化
train_predict = scaler.inverse_transform(train_predict)
test_predict = scaler.inverse_transform(test_predict)
original_ytrain = scaler.inverse_transform(y_train.reshape(-1, 1))
original_ytest = scaler.inverse_transform(y_test.reshape(-1, 1))

# 由于我们的预测是基于过去time_step天的数据来预测下一天，
# 所以我们需要从时间序列中剔除前time_step天以匹配预测数据的日期
test_dates = dates[train_size + time_step:len(dates)-1]
assert len(test_dates) == len(test_predict), "日期数组和预测结果的长度不匹配"
# 绘制结果
plt.figure(figsize=(10,6))
plt.plot(test_dates, original_ytest, label='Actual')
plt.plot(test_dates, test_predict, label='Predicted', color='red')
plt.title('Parking Spot Demand Prediction using LSTM')
plt.xlabel('Date')
plt.ylabel('Demand')
plt.xticks(rotation=45)  # Rotate date labels for better readability
plt.legend()
plt.show()

