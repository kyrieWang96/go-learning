# Migrations

请在此处添加migration的sql脚本
该事例用的是PostgresSQL执行操作

## how-to

### CLI usage

1. 首先使用`migrate`[获取地址](https://github.com/golang-migrate/migrate/releases)命令行工具创建 migration 文件

``` bash
# 例如：在这里创建了一个新的migration文件
migrate create -ext sql -dir database/migrations -seq user
``` 

运行完命令之后将会在此文件夹先生成`000001_user.down.sql`和`000001_user.up.sql`文件, 分别对应了创建和回退用户表的sql脚本

2. 编辑sql脚本文件

```sql
--- 000001_user.up.sql, 在该文件中添加升级脚本
create table public.users
(
    id         serial primary key,
    name       varchar(255)             not null,
    email      varchar(255)             not null,
    password   varchar(255)             not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null
);

```

```sql
--- 000001_user.down.sql, 在该文件中添加回退脚本
drop table "public"."users";
```

3. 使用migrate工具同步数据库

``` bash
migrate -database 'postgres://username:password@ip:port/database_name?sslmode=disable' -path database/migrations up 1
```
执行完成之后的输出应该如下所示

```text
1/u user (462.58435ms)
```

同时会在数据库中创建`public.users`表和`public.shema_migrations`表，其中`public.shema_migrations`表用于记录已经执行过的migration，会在其中记录下同步的版本

4. 修改user表

运行如下命令

```bash
migrate create  -ext sql -dir database/migrations -seq user
```

运行完命令之后将会在此文件夹先生成`000002_user.down.sql`和`000002_user.up.sql`文件。同样在这两个文件中添加sql脚本，分别对应升级和回退用户表的sql脚本

```sql
--- 000002_user.up.sql, 在该文件中添加升级脚本
alter table public.users add column phone integer;

```

```sql
--- 000002_user.down.sql, 在该文件中添加回退脚本
alter table public.users drop column phone;
```

之后运行命令行工具，升级数据库

```bash
migrate -database 'postgres://username:password@ip:port/database_name?sslmode=disable' -path database/migrations up 2
```

执行完成后，将会在数据库`public.users`表的`phone`字段，同时在`public.shema_migrations`表中记录下升级的版本

5. 回滚

```bash
migrate -database 'postgres://username:password@ip:port/database_name?sslmode=disable' -path database/migrations down 1
```

这样将会回滚到上一个版本，即`000001_user.up.sql`中创建的数据库表

## TODO
1。执行失败的回滚操作暂未记录，后续调研
2。mysql执行是一样的，适用于测试环境初始化表及初始化数据，若需要在线上环境初始化一样执行就行