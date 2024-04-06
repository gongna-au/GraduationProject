import pandas as pd
import numpy as np
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import LSTM, Dense
from sklearn.preprocessing import MinMaxScaler
import matplotlib.pyplot as plt

# 加载数据
df = pd.read_csv('dataset.csv', parse_dates=['LastUpdated'], index_col='LastUpdated')

# 使用占用率作为我们要预测的特征
occupancy = df['Occupancy'].values.reshape(-1, 1)

# 数据归一化
scaler = MinMaxScaler(feature_range=(0, 1))
occupancy_scaled = scaler.fit_transform(occupancy)

# 数据集构建函数
def create_dataset(data, look_back=1):
    X, Y = [], []
    for i in range(len(data) - look_back - 1):
        a = data[i:(i + look_back), 0]
        X.append(a)
        Y.append(data[i + look_back, 0])
    return np.array(X), np.array(Y)

# 分割数据集
look_back = 1
X, y = create_dataset(occupancy_scaled, look_back)
X = np.reshape(X, (X.shape[0], 1, X.shape[1]))

# 构建LSTM模型
model = Sequential()
model.add(LSTM(4, input_shape=(1, look_back)))
model.add(Dense(1))
model.compile(loss='mean_squared_error', optimizer='adam')

# 划分训练集和测试集
train_size = int(len(X) * 0.67)
test_size = len(X) - train_size
trainX, testX = X[0:train_size], X[train_size:len(X)]
trainY, testY = y[0:train_size], y[train_size:len(y)]

# 训练模型
model.fit(trainX, trainY, epochs=100, batch_size=1, verbose=2)

# 进行预测
trainPredict = model.predict(trainX)
testPredict = model.predict(testX)

# 反归一化
trainPredict = scaler.inverse_transform(trainPredict)
trainY = scaler.inverse_transform([trainY])
testPredict = scaler.inverse_transform(testPredict)
testY = scaler.inverse_transform([testY])

# 计算误差
trainScore = np.sqrt(np.mean((trainPredict[:,0] - trainY[0])**2))
print(f'Train Score: {trainScore:.2f} RMSE')
testScore = np.sqrt(np.mean((testPredict[:,0] - testY[0])**2))
print(f'Test Score: {testScore:.2f} RMSE')

# 绘制结果
plt.figure(figsize=(10, 6)) # 设置图形大小
plt.plot(df.index, scaler.inverse_transform(occupancy_scaled), label='Actual Occupancy', color='grey', alpha=0.5) # 实际占用率
plt.plot(df.index[:len(trainPredict)+look_back], np.append(np.full(look_back, np.nan), trainPredict[:,0]), label='Train Predictions', color='blue') # 训练集预测
# 假设我们已经正确计算了测试集预测的起始索引 start_index
#start_index_for_plot = len(trainPredict) + (look_back * 2) + 1
start_index_for_plot = len(trainPredict) + (look_back * 2)
# 确保绘图时 x 和 y 的长度匹配
# 注意这里使用了与 testPredict 相同长度的时间索引切片来确保匹配

plt.plot(df.index[start_index_for_plot:start_index_for_plot + len(testPredict)], testPredict[:,0], label='Test Predictions', color='red')

#plt.plot(df.index[len(trainPredict)+(look_back*2)+1:], testPredict[:,0], label='Test Predictions', color='red') # 测试集预测
plt.legend()
plt.xlabel('Time')
plt.ylabel('Occupancy')
plt.title('Parking Space Occupancy Predictions')
plt.tight_layout() # 自动调整子图参数, 使之填充整个图像区域
plt.show(block=True)