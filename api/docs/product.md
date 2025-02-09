# 商品相关 API

## 1. 获取商品列表
### 接口描述
获取所有商品的列表信息

### 请求方法
GET

### 请求路径
/product/list

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
| product_name| string | 是   | 搜索关键词|
| page   | int    | 否   | 页码     |
| size   | int    | 否   | 每页数量 |

### 响应参数
同获取商品列表接口

## 3. 获取商品详情
### 接口描述
获取指定商品的详细信息

### 请求方法
GET

### 请求路径
/product/info/{id}

### 路径参数
| 参数名| 类型 | 必选 | 说明   |
|------|------|------|--------|
| id   | int  | 是   | 商品ID |

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
        ]
    }
}

## 4. 获取分类商品
### 接口描述
获取指定分类下的所有商品

### 请求方法
GET

### 请求路径
/product/category/{id}

### 路径参数
| 参数名| 类型 | 必选 | 说明   |
|------|------|------|--------|
| id   | int  | 是   | 分类ID |

### 请求参数
| 参数名 | 类型 | 必选 | 说明     |
|--------|------|------|----------|
| page   | int  | 否   | 页码     |
| size   | int  | 否   | 每页数量 |

### 响应参数
同获取商品列表接口

## 5. 添加商品分类
### 接口描述
添加新的商品分类

### 请求方法
POST

### 请求路径
/product/add/category

### 请求参数
| 参数名      | 类型   | 必填 | 说明     | 示例     |
|------------|--------|------|----------|----------|
| name       | string | 是   | 分类名称  | "冬装"   |
| description| string | 否   | 分类描述  | "冬季服装"|

### 请求示例
{
    "name": "冬装",
    "description": "冬季服装"
}

### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": {
        "message": "添加分类成功"
    }
}

## 6. 添加商品
### 接口描述
添加新商品

### 请求方法
POST

### 请求路径
/product/add/product

### 请求参数
| 参数名      | 类型    | 必填 | 说明     | 示例     |
|------------|---------|------|----------|----------|
| name       | string  | 是   | 商品名称  | "冬季外套"|
| description| string  | 否   | 商品描述  | "保暖舒适"|
| type       | string  | 是   | 商品类型  | "服装"   |
| price      | float64 | 是   | 商品价格  | 299.99   |
| cover      | string  | 否   | 封面图URL | "http://..."|
| publish_time| string | 否   | 发布时间  | "2024-01-01"|
| link       | string  | 否   | 商品链接  | "http://..."|
| category_ids| []uint | 是   | 分类ID列表| [1,2]    |

### 请求示例
{
    "name": "冬季外套",
    "description": "保暖舒适",
    "type": "服装",
    "price": 299.99,
    "cover": "http://example.com/cover.jpg",
    "publish_time": "2024-01-01",
    "link": "http://example.com/product/1",
    "category_ids": [1, 2]
}

### 响应示例
{
    "status": 10000,
    "info": "success",
    "data": {
        "message": "添加商品成功"
    }
}