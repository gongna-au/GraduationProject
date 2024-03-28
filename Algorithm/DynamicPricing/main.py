import pandas as pd
import matplotlib.pyplot as plt

data = pd.read_csv('Parking Data.csv')

df = pd.DataFrame(data)

# 定义定价规则函数
def dynamic_pricing(row):
    base_price = 5  # 基础价格
    parking_space_addition = 0.1  # 每个停车位增加的费用
    leased_addition = 2  # 如果是租赁的，额外增加的费用

    # 计算总价格
    price = base_price + (row['Total Parking Spaces'] * parking_space_addition)
    
    if row['Owned/Leased'] == 'LEASED':
        price += leased_addition

    return price

# 应用定价规则
df['Dynamic Price'] = df.apply(dynamic_pricing, axis=1)

owned = df[df['Owned/Leased'] == 'OWNED']
leased = df[df['Owned/Leased'] == 'LEASED']

# 绘制拥有状态下的停车位数量与动态价格的关系
plt.scatter(owned['Total Parking Spaces'], owned['Dynamic Price'], color='blue', label='Owned')

# 绘制租赁状态下的停车位数量与动态价格的关系
plt.scatter(leased['Total Parking Spaces'], leased['Dynamic Price'], color='red', label='Leased')

# 添加图例
plt.legend()

# 添加标题和轴标签
plt.title('Dynamic Price vs. Total Parking Spaces')
plt.xlabel('Total Parking Spaces')
plt.ylabel('Dynamic Price')

# 显示图表
plt.show()