import torch
from PIL import Image

# 加载预训练的YOLOv5模型
model = torch.hub.load('ultralytics/yolov5', 'yolov5s', pretrained=True)

# 加载你的图像
img_path = 'test1.png'  # 更改为你的图像路径
img = Image.open(img_path)

# 进行推理
results = model(img)

# 显示结果
results.show()

# 获取检测的车牌信息
# 注意：这里假设车牌被模型标记为'class 0'
# 实际上你可能需要根据自己的数据集来调整这个类别标签
detected_plates = results.xyxy[0]  # 获取所有检测到的对象的边界框信息
for i, det in enumerate(detected_plates):
    if int(det[-1]) == 0:  # 假设车牌的类别ID为0
        print(f"车牌 {i+1}: ", det)

# 注意：上述代码是基于YOLOv5的通用目标检测。如果你要针对特定的车牌检测任务，
# 可能需要使用专门针对车牌检测训练过的模型，或者对YOLOv5模型进行微调训练。
