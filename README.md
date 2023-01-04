

## 运维面板



## 模块

- [ ] 首页
    - [x] 系统状态

- [ ] 软件管理
    - [ ] 软件列表
    - [ ] 软件操作

- [ ] 任务管理
    - [ ] 定时任务
    - [ ] 任务列表
    - [ ] 新增任务
    - [ ] 删除任务
    - [ ] 修改任务
    - [ ] 任务详情

- [ ] 其他
    - [ ] 登录
    

```SQL
insert into config_basic (`id`, `key`, `value`) values (1, "admin", "123456");
insert into task_basic (`id`, `spec`, `shell_path`, `name`) values (1, "ls", "./usr", "list");
```

```
curl -X POST -d '{"cron": "* * * * *", "exec": "touch /tmp/hello.txt"}' http://localhost:8080/v1/jobs
生成定时任务，ID 为 1，每分钟跑 1 次

curl http://localhost:8080/v1/jobs
[{"id":1,"next":"2023-01-04T22:59:00+08:00"} 下一次执行时间



```