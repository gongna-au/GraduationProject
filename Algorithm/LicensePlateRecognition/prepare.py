import cv2
import numpy as np

# 读取图片
image = cv2.imread('path_to_your_image.jpg')

# 1. 灰度化
gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

# 2. 二值化
# 先使用 Otsu 的自动阈值法
_, binary = cv2.threshold(gray, 0, 255, cv2.THRESH_BINARY + cv2.THRESH_OTSU)
# 或者使用自适应阈值法进行二值化
# binary = cv2.adaptiveThreshold(gray, 255, cv2.ADAPTIVE_THRESH_GAUSSIAN_C, 
#                                cv2.THRESH_BINARY, 11, 2)

# 3. 去噪声
# 使用高斯模糊去噪声
denoised = cv2.GaussianBlur(binary, (5, 5), 0)

# 4. 锐化
# 创建一个锐化滤波器
kernel_sharpening = np.array([[-1, -1, -1],
                              [-1, 9, -1],
                              [-1, -1, -1]])
# 应用滤波器
sharpened = cv2.filter2D(denoised, -1, kernel_sharpening)

# 显示处理后的图像（调试用）
cv2.imshow('Gray', gray)
cv2.imshow('Binary', binary)
cv2.imshow('Denoised', denoised)
cv2.imshow('Sharpened', sharpened)

cv2.waitKey(0)
cv2.destroyAllWindows()
