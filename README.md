# FISCO-BCOS-Demo

## 安装

### 前端安装

```sh
cd vue-manage-system
npm install
npm run serve
```

### 后端安装

```sh
go get -v github.com/Chris-Ju/FISCO-BCOS-Demo
go run .
```

### 链端安装

[参考🔗](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/installation.html)

---

## 使用说明

确保 8080 端口 以及 8000 端口空闲，服务启动后，访问 http://localhost:8080

---

## 文档

[实验报告](./16340098_鞠擘.pdf)  
[接口说明](./api/go-client/README.md)

---

## 加分析说明

- 使用了 vue 框架，实现前后端分离， vue-manage-system 模板去进行前端设计。
  - [代码链接](./vue-manage-system)
- 尝试使用 go-sdk 与 python-sdk 进行链端处理。
  - [代码链接](./fisco-bcos-go)
- 增加了管理员 api，可以对所有公司以及账单进行操作。
  - [代码链接](./controller/adminHandle.gp)


---

## 视频演示

[![mv](./image/example.png)](./image/mv.mov)