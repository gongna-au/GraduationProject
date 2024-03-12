import easyocr
reader = easyocr.Reader(['ch_sim','en'])  # 指定中文简体和英文
result = reader.readtext('test10.png')
for detection in result:
    text = detection[1]
    print(text)
