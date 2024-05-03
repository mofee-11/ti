# ti

ti 是一个基于文本的个人数据库

## 用例

```bash
git clone https://github.com/mofee-11/ti.git
cd ti
go run . -c ti.toml
文档总数：1000
```

为了调用简洁，编译后使用：

```bash
$ go build
$ ./ti
文档总数：1000
$ ./ti ccl
58      cclowser1m
420     ccleynebo
$ ./ti 58
---
date: 7/26/2023
path: LaciniaEget.tiff
resource: null
title: cclowser1m
---

test⁠test
```

## 文件结构

编译后与 `ti` 的整个运行生命周期读且仅读取两个文件

### 配置文件

配置文件用 toml 格式编写，它默认与 ti 可执行文件处于同一个目录，也可以用 `-c` 选项手动指定其他目录：

```bash
ti -c config.toml
```

默认配置文件内容有且仅有一行：

```toml
db = "MOCK_DATA.json"
```

通过 `db` 变量配置的数据库文件的路径

### 数据库文件

其为 json 格式的数组。其内容结构示例：

```json
[
  {
    "title": "fhuxstep0",
    "resource": null,
    "content": "₀₁₂"
  },
  {
    "a": "jmctavish4",
    "b": "(｡◕ ∀ ◕｡)"
  },
  {
    "tag": ["set","add","get","del"],
    "path": "QuisOrciNullam.gif",
    "date": "6/14/2023",
    "obj": {
      "key": "value",
      "id": 0
    }
  }
]
```

其中 ti 只关心根元素是否为数组，数组元素是否为对象。如果不是就会报错退出

而数组元素的内部具体有哪些字段、类型、嵌套。ti 统统都可以解析

而 ti 称呼这些数组元素为「文档」

## 行为解释

调用 ti 时，用户传入的参数类型会改变 ti 的行为：

- 空参数：打印数据库文件文档数量
- 1 个参数
  - 可以转换为整数：以 markdown 格式打印对应文档
  - 无法转化为整数：遍历搜索数据库，给出 20 个结果
- 多个参数：遍历搜索

当 ti 尝试以 markdown 打印文档时

- 它会寻找文档的 `content` 字段，作为 markdown 正文
- 其他字段作为 markdown 的 frontmatter 以 yaml 格式编码
- 若是这些字段为空或是不存在。ti 仅打印空行，不会报错

当 ti 尝试将参数作为搜索关键字：

- ti 会遍历每个文档的键值对，检查其值是否包含关键字
  - 包含便返回结果，不包含则继续处理下一个
- ti 搜索时会忽略键名，仅关心值
- 有多个关键字时，ti 会返回包含所有关键字的结果
