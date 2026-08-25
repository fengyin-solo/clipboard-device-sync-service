# Clipboard 云端剪贴板同步服务

纯 Go 标准库实现的云端剪贴板后端服务，支持多设备条目同步、分组、标签与同步日志。

## 运行说明

```bash
cd origin
/Users/fengyin/.local/go/bin/go run ./cmd/server
```

默认监听 `:8080`，可通过环境变量调整：

- `PORT` — 服务端口（默认 8080）
- `ADDR` — 完整监听地址（覆盖 PORT）
- `MAX_PAGE_SIZE` — 最大分页大小（默认 100）
- `LOG_LEVEL` — 日志级别：debug / info / warn / error（默认 info）

## API 一览

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/entries | 创建剪贴板条目 |
| GET | /api/entries | 条目列表（支持分页、group_id、content_type、status、keyword、pinned 筛选） |
| GET | /api/entries/{id} | 获取单个条目 |
| PUT | /api/entries/{id} | 更新条目（含状态机校验） |
| DELETE | /api/entries/{id} | 删除条目 |
| POST | /api/entries/batch | 批量创建条目 |
| POST | /api/entries/batch-delete | 批量删除条目 |
| POST | /api/entries/batch-status | 批量更新条目状态 |
| POST | /api/entries/batch-pin | 批量置顶/取消置顶 |
| POST | /api/groups | 创建分组 |
| GET | /api/groups | 分组列表（支持 keyword 筛选） |
| GET | /api/groups/{id} | 获取分组 |
| PUT | /api/groups/{id} | 更新分组 |
| DELETE | /api/groups/{id} | 删除分组 |
| POST | /api/tags | 创建标签 |
| GET | /api/tags | 标签列表（支持 keyword 筛选） |
| GET | /api/tags/{id} | 获取标签 |
| PUT | /api/tags/{id} | 更新标签 |
| DELETE | /api/tags/{id} | 删除标签 |
| POST | /api/entry-tags | 关联条目与标签（外键校验） |
| GET | /api/entry-tags | 关联列表（支持 entry_id、tag_id 筛选） |
| GET | /api/entry-tags/{id} | 获取关联 |
| DELETE | /api/entry-tags/{id} | 删除关联 |
| POST | /api/devices | 注册设备 |
| GET | /api/devices | 设备列表（支持 platform、keyword 筛选） |
| GET | /api/devices/{id} | 获取设备 |
| PUT | /api/devices/{id} | 更新设备 |
| DELETE | /api/devices/{id} | 删除设备 |
| POST | /api/sync-logs | 创建同步记录（外键校验） |
| GET | /api/sync-logs | 同步记录列表（支持 device_id、direction、status 筛选） |
| GET | /api/sync-logs/{id} | 获取同步记录 |
| PUT | /api/sync-logs/{id} | 更新同步记录 |
| DELETE | /api/sync-logs/{id} | 删除同步记录 |
| GET | /api/stats/overview | 聚合统计：条目总数、按类型/状态分组计数、分组数、Top 标签 |
