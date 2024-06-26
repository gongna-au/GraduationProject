## 环境
```shell
python3 -m venv venv
```

```shell
source venv/bin/activate
```

```shell
pip install matplotlib
```

```shell
pip install numpy
```
## 时间序列预测算法




### 1.自回归积分滑动平均模型（ARIMA）
一个经典的时间序列预测模型，适用于预测一段时间内停车需求的变化。

```shell
pip install numpy pandas matplotlib statsmodels
```

```shell
python3 main1.py
```

结果
```text
未来7天的预测停车位最高占用率：
2024-01-01    79.707733
2024-01-02    79.906696
2024-01-03    79.912598
2024-01-04    79.912773
2024-01-05    79.912778
2024-01-06    79.912778
2024-01-07    79.912778
Freq: D, Name: predicted_mean, dtype: float64
```



### 2.季节性自回归积分滑动平均模型（SARIMA）
ARIMA模型的扩展，特别适合处理具有季节性变化的数据，如工作日与周末的停车位需求差异。

```shell
pip install numpy pandas matplotlib scikit-learn
```

```shell
python3 main2.py
```
在这个例子中，首先生成了一个模拟的停车位需求时间序列数据，然后使用SARIMAX类从statsmodels库来拟合一个SARIMA模型。模型的参数(1, 1, 1) x (1, 1, 1, 12)是示例参数，实际应用中需要根据数据进行调整。这里的季节性周期被设置为12，假设我们的数据是按月度的季节性模式，对于每日数据可能需要调整。

拟合模型后，使用模型的forecast方法来预测未来30天的停车位需求，并使用matplotlib绘制出原始数据和预测结果。

请注意，SARIMA模型参数的选择（包括季节性参数）

## 机器学习算法

机器学习算法能够处理复杂的数据特征，并从历史数据中学习模式来预测未来的可用性。

### 随机森林（Random Forest）：一个强大的分类和回归工具，能够处理大量特征并预测停车位的可用性。

```shell
pip install numpy pandas matplotlib scikit-learn
```

```shell
python3 main3.py
```
首先生成了模拟的每日停车位需求数据，并根据日期添加了月份和星期几作为特征，以捕获数据的季节性和周期性模式。然后，我们将数据集分为训练集和测试集，使用随机森林模型进行训练，并对测试集进行预测。最后，我们通过图表直观展示了实际需求和模型预测的对比。

请注意，随机森林模型的预测效果很大程度上依赖于所选特征能否合理捕捉时间序列数据的季节性和趋势。

### 梯度提升机（Gradient Boosting Machine, GBM）：一种高效的预测模型，特别适合处理非线性关系和复杂的数据结构。
### 深度学习模型：如卷积神经网络（CNN）和循环神经网络（RNN），可以从时间序列数据中捕获深层特征，适合于动态预测停车位的可用性。
### LSTM模型进行时间序列预测

这个例子采用了LSTM，一种常用于捕捉长期依赖关系的循环神经网络（RNN）结构，来预测停车位的可用性。下面是使用LSTM模型，但相似的方法也可用于GRU或Transformer模型。

首先生成了模拟的每日停车位需求数据，并进行了归一化处理。然后，定义了一个create_dataset函数来准备时间序列数据集，并使用了一个LSTM模型来进行训练和预测。最后，我们通过图表直观展示了实际需求与模型预测的对比。

请注意，预测结果的准确性受多种因素影响，包括模型结构、超参数设置、数据预处理方法等。如果预测结果不够精确，可以尝试调整这些因素来优化模型。

```shell
pip install tensorflow
```

```shell
python3 main4.py
```
**横轴（X轴）**代表时间。由于这是一个时间序列预测的例子，横轴通常用来表示时间的流逝，可以是天、月、年等时间单位。在这个特定的例子中，由于我们生成了每日的数据点，并尝试预测未来的需求，横轴代表的是连续的日子。

**纵轴（Y轴）**代表停车位的需求量。在这个场景中，需求量可能指的是某个时间点的停车位占用率，或者是特定时间段内寻求停车位的车辆数量。纵轴的数值表示需求的大小，根据预处理步骤（如归一化），这些值可能代表实际需求量的原始数值，或者是经过转换的数值范围。

## 回归分析
线性回归：可以用来预测基于时间和其他因素（如特殊活动、假期）对停车位需求的影响。
逻辑回归：虽然通常用于分类问题，但也可以应用于预测停车位的占用概率。

## 聚类分析

### K-均值聚类（K-means Clustering）：通过聚类分析不同时间段或地点的停车模式，帮助预测特定时段或区域的停车位可用性。

## 神经网络和深度学习

长短期记忆网络（LSTM）：一种特殊的RNN，非常适合处理和预测时间序列数据中的长期依赖问题。