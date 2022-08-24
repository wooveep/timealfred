# timealfred

>可以显示多种时间格式、时间格式转时间戳、时间戳转普通时间格式的 alfred 5 workflow 工作流
<img width="601" alt="image" src="https://user-images.githubusercontent.com/37527190/186317101-c009aaff-2801-41ca-b4b1-7cac7284d285.png">

## 条件
[Alfred 5](https://www.alfredapp.com) with the paid [Powerpack](https://www.alfredapp.com/powerpack/) upgrade

## 安装
双击 “times.alfredworkflow” 即可安装至alfred工作流中

## 使用
输入 `ts` 默认显示当前时间

输入 `ts 时间戳` 支持毫秒时间戳 显示转换后的时间

输入 `ts 时间`  支持的时间格式为 “YYYY-MM-DD hh:mm:ss” 、“YYYYMMDDhhmmss” 、“RFC3339格式” 显示转换后其他时间格式和时间戳

如果未能识别解析`ts`后的时间格式默认显示当前时间


