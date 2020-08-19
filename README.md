<img src="./logo/logo.png" width="200">

# yoxi

> All resource files come from https://noiz.io

小夕是一款白噪音软件, 让你更专注地工作、更专心学习. 关于白噪音是什么, 请右转[知乎](https://zhuanlan.zhihu.com/p/34815642)

![image.png](https://i.loli.net/2020/08/19/C2d6agQUxpBh78t.png)

## build

(目前暂时只支持`osx`)

依赖

- go1.14
- nodejs(yarn)
- unix环境

```bash
git clone https://github.com/d1y/yoxi
git clone https://github.com/d1y/yoxi_web
git clone https://github.com/d1y/yoxi_data

# 请先编译出`web`
cd yoxi_web
yarn build

cd ..
go get
cd build
chmod u+x bootstarp.sh
./bootstarp.sh
```