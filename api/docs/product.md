# 商品相关 API

## 1. 获取商品列表
### 接口描述
获取所有商品的列表信息

### 请求方法
GET

### 请求路径
/product/list

### 请求头
| 参数名        | 说明         | 示例                                    |
|--------------|--------------|----------------------------------------|
| Authorization| Bearer Token | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI... |

### 响应参数
| 参数名 | 类型   | 说明     | 示例     |
|--------|--------|----------|----------|
| status | int    | 状态码   | 10000    |
| info   | string | 状态信息 | "success"|
| data   | object | 返回数据 | 见示例   |

### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": {
        "products": [
            {
                "product_id": 1,
                "name": "商品名称",
                "description": "商品描述",
                "type": "商品类型",
                "comment_num": 0,
                "price": 99.99,
                "is_addCart": false,
                "cover": "商品封面图片URL",
                "publish_time": "2024-01-01",
                "link": "商品链接",
                "categories": [
                    {
                        "category_id": 1,
                        "name": "分类名称",
                        "description": "分类描述"
                    }
                ]
            }
        ]
    }
}

## 2. 搜索商品
### 接口描述
根据商品名称搜索商品

### 请求方法
GET

### 请求路径
/product/search

### 请求参数
| 参数名 | 类型   | 必选 | 说明     |
|--------|--------|------|----------|
| keyword| string | 是   | 搜索关键词|
| page   | int    | 否   | 页码     |
| size   | int    | 否   | 每页数量 |

### 请求头
| 参数名        | 说明         | 示例                                    |
|--------------|--------------|----------------------------------------|
| Authorization| Bearer Token | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI... |

### 响应参数
同获取商品列表接口

## 3. 获取商品详情
### 接口描述
获取指定商品的详细信息

### 请求方法
GET

### 请求路径
/product/{product_id}

### 路径参数
| 参数名     | 类型 | 必选 | 说明   |
|-----------|------|------|--------|
| product_id| int  | 是   | 商品ID |

### 请求头
| 参数名        | 说明         | 示例                                    |
|--------------|--------------|----------------------------------------|
| Authorization| Bearer Token | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI... |

### 响应参数
| 参数名 | 类型   | 说明     | 示例     |
|--------|--------|----------|----------|
| status | int    | 状态码   | 10000    |
| info   | string | 状态信息 | "success"|
| data   | object | 返回数据 | 见示例   |

### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": {
        "product_id": 1,
        "name": "商品名称",
        "description": "商品详细描述",
        "type": "商品类型",
        "comment_num": 10,
        "price": 99.99,
        "is_addCart": false,
        "cover": "商品封面图片URL",
        "images": ["图片URL1", "图片URL2"],
        "publish_time": "2024-01-01",
        "link": "商品链接",
        "categories": [
            {
                "category_id": 1,
                "name": "分类名称",
                "description": "分类描述"
            }
        ],
        "comments": [
            {
                "comment_id": 1,
                "user_id": 1,
                "username": "用户名",
                "avatar": "用户头像URL",
                "content": "评论内容",
                "create_time": "2024-01-01",
                "reply_count": 0
            }
        ]
    }
}

## 4. 获取分类商品
### 接口描述
获取指定分类下的所有商品

### 请求方法
GET

### 请求路径
/product/category/{category_id}

### 路径参数
| 参数名      | 类型 | 必选 | 说明   |
|------------|------|------|--------|
| category_id| int  | 是   | 分类ID |

### 请求参数
| 参数名 | 类型 | 必选 | 说明     |
|--------|------|------|----------|
| page   | int  | 否   | 页码     |
| size   | int  | 否   | 每页数量 |

### 请求头
| 参数名        | 说明         | 示例                                    |
|--------------|--------------|----------------------------------------|
| Authorization| Bearer Token | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI... |

### 响应参数
同获取商品列表接口