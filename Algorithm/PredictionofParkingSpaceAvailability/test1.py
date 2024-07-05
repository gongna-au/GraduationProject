import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from keras.models import Sequential
from keras.layers import LSTM, Dense
from sklearn.preprocessing import MinMaxScaler
from sklearn.model_selection import train_test_split

# 加载数据
data = """
BHMNCPPLS01,450,197,2016-10-22
BHMNCPPLS01,450,192,2016-10-23
BHMNCPPLS01,450,76,2016-10-24
BHMNCPPLS01,450,96,2016-10-25
BHMNCPPLS01,450,115,2016-10-26
BHMNCPPLS01,450,123,2016-10-27
BHMNCPPLS01,450,127,2016-10-28
BHMNCPPLS01,450,176,2016-10-29
BHMNCPPLS01,450,148,2016-10-30
BHMNCPPLS01,450,66,2016-10-31
BHMNCPPLS01,450,82,2016-11-01
BHMNCPPLS01,450,87,2016-11-02
BHMNCPPLS01,450,111,2016-11-03
BHMNCPPLS01,450,103,2016-11-04
BHMNCPPLS01,450,183,2016-11-05
BHMNCPPLS01,450,191,2016-11-06
BHMNCPPLS01,450,92,2016-11-07
BHMNCPPLS01,450,112,2016-11-08
BHMNCPPLS01,450,103,2016-11-09
BHMNCPPLS01,450,102,2016-11-10
BHMNCPPLS01,450,127,2016-11-11
BHMNCPPLS01,450,214,2016-11-12
BHMNCPPLS01,450,156,2016-11-13
BHMNCPPLS01,450,77,2016-11-14
BHMNCPPLS01,450,99,2016-11-15
BHMNCPPLS01,450,100,2016-11-16
BHMNCPPLS01,450,121,2016-11-17
BHMNCPPLS01,450,106,2016-11-18
BHMNCPPLS01,450,224,2016-11-19
BHMNCPPLS01,450,246,2016-11-20
BHMNCPPLS01,450,116,2016-11-21
BHMNCPPLS01,450,94,2016-11-22
BHMNCPPLS01,450,113,2016-11-23
BHMNCPPLS01,450,123,2016-11-24
BHMNCPPLS01,450,146,2016-11-25
BHMNCPPLS01,450,377,2016-11-26
BHMNCPPLS01,450,250,2016-11-27
BHMNCPPLS01,450,109,2016-11-29
BHMNCPPLS01,450,113,2016-11-30
BHMNCPPLS01,450,159,2016-12-01
BHMNCPPLS01,450,149,2016-12-02
BHMNCPPLS01,450,144,2016-12-03
BHMNCPPLS01,450,145,2016-12-04
BHMNCPPLS01,450,184,2016-12-05
BHMNCPPLS01,450,127,2016-12-06
BHMNCPPLS01,450,153,2016-12-07
BHMNCPPLS01,450,157,2016-12-08
BHMNCPPLS01,450,170,2016-12-09
BHMNCPPLS01,450,204,2016-12-10
BHMNCPPLS01,450,307,2016-12-11
BHMNCPPLS01,450,179,2016-12-12
BHMNCPPLS01,450,153,2016-12-13
BHMNCPPLS01,450,144,2016-12-14
BHMNCPPLS01,450,176,2016-12-15
BHMNCPPLS01,450,157,2016-12-16
BHMNCPPLS01,450,215,2016-12-17
BHMNCPPLS01,450,244,2016-12-18
BHMNCPPLS01,450,280,2016-12-19

"""
data = [row.split(',') for row in data.strip().split('\n')]
df = pd.DataFrame(data, columns=['Location', 'Total_Spaces', 'Occupied_Spaces', 'Date'])

# 预处理
df['Occupied_Spaces'] = df['Occupied_Spaces'].astype(int)
df['Date'] = pd.to_datetime(df['Date'])
df['Day_of_Year'] = df['Date'].dt.dayofyear

X = df[['Day_of_Year']].values
y = df['Occupied_Spaces'].values

scaler = MinMaxScaler(feature_range=(0, 1))
X_scaled = scaler.fit_transform(X)
y_scaled = scaler.fit_transform(y.reshape(-1,1))

# 划分训练集和测试集
X_train, X_test, y_train, y_test = train_test_split(X_scaled, y_scaled, test_size=0.2, random_state=42)

# LSTM需要输入形状为[样本数, 时间步, 特征数]
X_train = X_train.reshape((X_train.shape[0], 1, X_train.shape[1]))
X_test = X_test.reshape((X_test.shape[0], 1, X_test.shape[1]))

# 构建LSTM模型
model = Sequential()
model.add(LSTM(50, input_shape=(X_train.shape[1], X_train.shape[2])))
model.add(Dense(1))
model.compile(optimizer='adam', loss='mean_squared_error')

# 训练模型
model.fit(X_train, y_train, epochs=50, batch_size=72, validation_data=(X_test, y_test), verbose=2, shuffle=False)

# 预测
y_pred = model.predict(X_test)
y_pred = scaler.inverse_transform(y_pred)  # 反归一化

# 打印预测结果（这里简化处理，仅为演示）
print(y_pred)
y_test_inv = scaler.inverse_transform(y_test.reshape(-1, 1))

# 绘制实际值和预测值
plt.figure(figsize=(10, 6))
plt.plot(y_test_inv, label='Actual', color='blue', marker='o')
plt.plot(y_pred, label='Predicted', color='red', linestyle='--', marker='x')
plt.title('Parking Space Occupancy: Actual vs Predicted')
plt.xlabel('Sample')
plt.ylabel('Occupied Spaces')
plt.legend()
plt.show()