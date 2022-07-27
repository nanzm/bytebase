# Bytebase API 风格指南

Bytebase 使用 REST API，这个文档描述了相应的 API 风格指南。

The guiding principal for our style guide is **consistency**.

# 方法

## 更喜欢 PATCH 而不是 PUT

大多数时候，我们只想对资源进行部分更新，我们应该相应地使用 PATCH。另一方面，PUT 意味着用请求字段覆盖整个资源，并且更有可能意外地重置现有字段。

# 资源 URL 命名

## 使用资源 ID 来寻址特定资源

Bytebase 使用自动增量 ID 作为所有资源的主键。要寻址特定资源，我们使用 GET `/issue/42` ，如果我们想支持其他寻址机制，例如使用资源名称，我们应该使用查询参数，例如`.issue/42?name=foo`

## 使用小写，不要使用分隔符来分割单词

使用`/foo/barbaz`代替`/foo/barBaz`或`/foo/bar-baz`

*理由*：规则很简单，因此提高了一致性。例如，假设我们有一个名为`datasource`的资源， `data source`和`datasource`都是可接受的术语，根据我们的规则，它始终是`datasource` 。有时这确实会损害可读性，但大多数时候，我们在 URL 的路径组件中最多只能有 2 个单词，而我们的大脑非常擅长识别单个单词。

*注意*：使用 camelCase 或 dash-case 更为常见。然而，我们并不孤单， [Kubernetes](https://kubernetes.io/docs/reference/)也采用了这种约定。

## 集合资源也使用单数形式

使用`GET /issue`而不是`GET /issues`来获取问题列表。

*基本原理*：复数形式有多种变化，非英语母语人士很难记住所有规则。在实践中，对集合资源使用单数形式不会与单数资源混淆，因为它们使用不同的资源路径，例如`/issue`与`/issue/:id` 。

*注意*：我们知道这与常见的约定不同。但是，我们并不孤单，请参阅[这个 Kubernetes 讨论](https://github.com/kubernetes/kubernetes/issues/18622)。

## 使用单独的`/{{resource}}/batch`进行批处理操作

如果资源支持批处理操作，则在该资源下使用单独的`/batch`端点。

# 留言

## 属性名称约定

我们按照[Google JSON Style Guide](https://google.github.io/styleguide/jsoncstyleguide.xml)使用 json 消息在后端和前端之间进行通信。属性名称必须是驼峰式、ascii 字符串。不同语言的变量名应该遵循自己的语言风格，例如 Go 和 Vue。但是，我们必须为 Go API 结构中的每个字段使用 json 注释，以强制在线上使用相同的样式，并通过重构防止任何破坏性更改，因为 Go 会自动根据字段名称设置 json 属性名称。

我们可以将以下示例视为一个有趣的案例。 helloID 遵循 Go 风格，而有线消息使用 helloId 以符合 Vue 约定。

1. 转到结构字段： `helloID string `json:"helloId"`` 。
2. Json 属性名称： `helloId` 。
3. Vue 模板名称： `helloId` =&gt; `hello-id` 。

# 杂项

1. 时间戳应尽可能以秒为单位的 Unix 时间戳（UTC 时区）。名称应采用`xxTs`格式，例如`createdTs` 。需要精度的时间戳应该是纳秒，例如性能分析。名称应采用`xxNs`格式。

# 参考

1. [谷歌的 AIP](https://google.aip.dev/)
2. [Kubernetes API 参考](https://kubernetes.io/docs/reference/)
