## 命令链接数据库

>   mysql -u root -p

## 创建数据库

>   CREATE DATABASE 数据库名

## 删除数据库

> drop database 数据库名

## 选择数据库

> use 数据库名

## 数据库类型

### 数值类型

 TINYINT(小整数) SMALLINT(大整数) MEDIUMINT(大整数) INT(大整数) INTEGER(大整数) FLOAT(单精度浮点数) DOUBLE (双精度浮点) DECIMAL(小数值)

<table class="reference">
    <tbody>
        <tr>
            <th width="10%">类型</th>
            <th width="15%">大小</th>
            <th width="30%">范围（有符号）</th>
            <th width="30%">范围（无符号）</th>
            <th width="15%">用途</th>
        </tr>
        <tr>
            <td>
                TINYINT
            </td>
            <td>
                1 字节
            </td>
            <td>
                (-128，127)
            </td>
            <td>
                (0，255)
            </td>
            <td>
                小整数值
            </td>
        </tr>
        <tr>
            <td>
                SMALLINT
            </td>
            <td>
                2 字节
            </td>
            <td>
                (-32 768，32 767)
            </td>
            <td>
                (0，65 535)
            </td>
            <td>
                大整数值
            </td>
        </tr>
        <tr>
            <td>
                MEDIUMINT
            </td>
            <td>
                3 字节
            </td>
            <td>
                (-8 388 608，8 388 607)
            </td>
            <td>
                (0，16 777 215)
            </td>
            <td>
                大整数值
            </td>
        </tr>
        <tr>
            <td>
                INT或INTEGER
            </td>
            <td>
                4 字节
            </td>
            <td>
                (-2 147 483 648，2 147 483 647)
            </td>
            <td>
                (0，4 294 967 295)
            </td>
            <td>
                大整数值
            </td>
        </tr>
        <tr>
            <td>
                BIGINT
            </td>
            <td>
                8 字节
            </td>
            <td>
                (-9,223,372,036,854,775,808，9 223 372 036 854 775 807)
            </td>
            <td>
                (0，18 446 744 073 709 551 615)
            </td>
            <td>
                极大整数值
            </td>
        </tr>
        <tr>
            <td>
                FLOAT
            </td>
            <td>
                4 字节
            </td>
            <td>
                (-3.402 823 466 E+38，-1.175 494 351 E-38)，0，(1.175 494 351 E-38，3.402 823 466 351 E+38)
            </td>
            <td>
                0，(1.175 494 351 E-38，3.402 823 466 E+38)
            </td>
            <td>
                单精度<br>浮点数值
            </td>
        </tr>
        <tr>
            <td>
                DOUBLE
            </td>
            <td>
                8 字节
            </td>
            <td>
                (-1.797 693 134 862 315 7 E+308，-2.225 073 858 507 201 4 E-308)，0，(2.225 073 858 507 201 4 E-308，1.797
                693 134 862 315 7 E+308)
            </td>
            <td>
                0，(2.225 073 858 507 201 4 E-308，1.797 693 134 862 315 7 E+308)
            </td>
            <td>
                双精度<br>浮点数值
            </td>
        </tr>
        <tr>
            <td>
                DECIMAL
            </td>
            <td>
                对DECIMAL(M,D) ，如果M&gt;D，为M+2否则为D+2
            </td>
            <td>
                依赖于M和D的值
            </td>
            <td>
                依赖于M和D的值
            </td>
            <td>
                小数值
            </td>
        </tr>
    </tbody>
</table>

### 日期和时间类型

<table class="reference">
    <tbody>
        <tr>
            <th width="10%">
                类型
            </th>
            <th width="10%">
                大小<br>(字节)
            </th>
            <th width="40%">
                范围
            </th>
            <th width="20%">
                格式
            </th>
            <th>
                用途
            </th>
        </tr>
        <tr>
            <td width="10%">
                DATE
            </td>
            <td width="10%">
                3
            </td>
            <td>
                1000-01-01/9999-12-31
            </td>
            <td>
                YYYY-MM-DD
            </td>
            <td>
                日期值
            </td>
        </tr>
        <tr>
            <td width="10%">
                TIME
            </td>
            <td width="10%">
                3
            </td>
            <td>
                '-838:59:59'/'838:59:59'
            </td>
            <td>
                HH:MM:SS
            </td>
            <td>
                时间值或持续时间
            </td>
        </tr>
        <tr>
            <td width="10%">
                YEAR
            </td>
            <td width="10%">
                1
            </td>
            <td>
                1901/2155
            </td>
            <td>
                YYYY
            </td>
            <td>
                年份值
            </td>
        </tr>
        <tr>
            <td width="10%">
                DATETIME
            </td>
            <td width="10%">
                8
            </td>
            <td width="40%">
                1000-01-01 00:00:00/9999-12-31 23:59:59
            </td>
            <td>
                YYYY-MM-DD HH:MM:SS
            </td>
            <td>
                混合日期和时间值
            </td>
        </tr>
        <tr>
            <td width="10%">
                TIMESTAMP
            </td>
            <td width="10%">
                4
            </td>
            <td width="40%">
                <p>1970-01-01 00:00:00/2038 </p>

                <p>结束时间是第 <strong>2147483647</strong> 秒，北京时间 <strong>2038-1-19 11:14:07</strong>，格林尼治时间 2038年1月19日 凌晨
                    03:14:07</p>
            </td>
            <td>
                YYYYMMDD HHMMSS
            </td>
            <td>
                混合日期和时间值，时间戳
            </td>
        </tr>
    </tbody>
</table>

### 字符串类型

CHAR、VARCHAR、BINARY、VARBINARY、BLOB、TEXT、ENUM和SET

<table class="reference">
    <tbody>
        <tr>
            <th width="20%">
                类型
            </th>
            <th width="25%">
                大小
            </th>
            <th width="55%">
                用途
            </th>
        </tr>
        <tr>
            <td>
                CHAR
            </td>
            <td>
                0-255字节
            </td>
            <td>
                定长字符串
            </td>
        </tr>
        <tr>
            <td>
                VARCHAR
            </td>
            <td>
                0-65535 字节
            </td>
            <td>
                变长字符串
            </td>
        </tr>
        <tr>
            <td>
                TINYBLOB
            </td>
            <td>
                0-255字节
            </td>
            <td>
                不超过 255 个字符的二进制字符串
            </td>
        </tr>
        <tr>
            <td>
                TINYTEXT
            </td>
            <td>
                0-255字节
            </td>
            <td>
                短文本字符串
            </td>
        </tr>
        <tr>
            <td>
                BLOB
            </td>
            <td>
                0-65 535字节
            </td>
            <td>
                二进制形式的长文本数据
            </td>
        </tr>
        <tr>
            <td>
                TEXT
            </td>
            <td>
                0-65 535字节
            </td>
            <td>
                长文本数据
            </td>
        </tr>
        <tr>
            <td>
                MEDIUMBLOB
            </td>
            <td>
                0-16 777 215字节
            </td>
            <td>
                二进制形式的中等长度文本数据
            </td>
        </tr>
        <tr>
            <td>
                MEDIUMTEXT
            </td>
            <td>
                0-16 777 215字节
            </td>
            <td>
                中等长度文本数据
            </td>
        </tr>
        <tr>
            <td>
                LONGBLOB
            </td>
            <td>
                0-4 294 967 295字节
            </td>
            <td>
                二进制形式的极大文本数据
            </td>
        </tr>
        <tr>
            <td>
                LONGTEXT
            </td>
            <td>
                0-4 294 967 295字节
            </td>
            <td>
                极大文本数据
            </td>
        </tr>
    </tbody>
</table>

## 创建数据表

创建MySQL数据表需要以下信息

* 表名
* 表字段名
* 定义每个表字段

> CREATE TABLE table_name (column_name column_type)

```sql
CREATE TABLE IF NOT EXISTS `test_tbl`(
   `runoob_id` INT UNSIGNED AUTO_INCREMENT,
   `runoob_title` VARCHAR(100) NOT NULL,
   `runoob_author` VARCHAR(40) NOT NULL,
   `submission_date` DATE,
   PRIMARY KEY ( `runoob_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

## 删除数据表

> DROP TABLE table_name

## 插入数据

> INSERT INTO table_name ( field1, field2,...fieldN ) VALUES ( value1, value2,...valueN );

```sql
INSERT INTO test_tbl 
    (runoob_title, runoob_author, submission_date) 
    VALUES 
    ("学习", "haha", NOW())

```

## 查询数据

* 查询语句中你可以使用一个或者多个表，表之间使用逗号(,)分割，并使用WHERE语句来设定查询条件。
* SELECT 命令可以读取一条或者多条记录。
* 你可以使用星号（*）来代替其他字段，SELECT语句会返回表的所有字段数据
* 你可以使用 WHERE 语句来包含任何条件。
* 你可以使用 LIMIT 属性来设定返回的记录数。
* 你可以通过OFFSET指定SELECT语句开始查询的数据偏移量。默认情况下偏移量为0。

```sql
SELECT column_name,column_name
FROM table_name
[WHERE Clause]
[LIMIT N][ OFFSET M]
```

## MySQL WHERE 子句

* 查询语句中你可以使用一个或者多个表，表之间使用逗号, 分割，并使用WHERE语句来设定查询条件。
* 你可以在 WHERE 子句中指定任何条件。
* 你可以使用 AND 或者 OR 指定一个或多个条件。
* WHERE 子句也可以运用于 SQL 的 DELETE 或者 UPDATE 命令。
* WHERE 子句类似于程序语言中的 if 条件，根据 MySQL 表中的字段值来读取指定的数据。

```sql
SELECT field1, field2,...fieldN FROM table_name1, table_name2...
[WHERE condition1 [AND [OR]] condition2.....
```

```sql
SELECT * FROM video_server.test_tbl where runoob_author="haha"
```

## MySQL UPDATE 更新

* 你可以同时更新一个或多个字段。
* 你可以在 WHERE 子句中指定任何条件。
* 你可以在一个单独表中同时更新数据。

```sql
UPDATE table_name SET field1=new-value1, field2=new-value2
[WHERE Clause]
```
```sql
UPDATE video_server.test_tbl SET runoob_title="学习4" WHERE runoob_author="haha";
```

## MySQL DELETE

* 如果没有指定 WHERE 子句，MySQL 表中的所有记录将被删除。
* 你可以在 WHERE 子句中指定任何条件
* 您可以在单个表中一次性删除记录。

```sql
DELETE FROM table_name [WHERE Clause]
```

```sql
DELETE FROM video_server.test_tbl WHERE runoob_author="haha";
```

## MySQL LIKE

WHERE 子句可以使用 = 来设定获取数据的条件，如 "runoob_author = 'RUNOOB.COM'"。
但是有时候我们需要获取 runoob_author 字段含有 "COM" 字符的所有记录，这时我们就需要在 WHERE 子句中使用 SQL LIKE 子句。
SQL LIKE 子句中使用百分号 %字符来表示任意字符，类似于UNIX或正则表达式中的星号 *。
如果没有使用百分号 %, LIKE 子句与等号 = 的效果是一样的。

以下是 SQL SELECT 语句使用 LIKE 子句从数据表中读取数据的通用语法：

```sql
SELECT field1, field2,...fieldN 
FROM table_name
WHERE field1 LIKE condition1 [AND [OR]] filed2 = "somevalue"
```

```sql
select * from test_tbl where runoob_author like "%COM"
```

* 你可以在 WHERE 子句中指定任何条件。
* 你可以在 WHERE 子句中使用LIKE子句。
* 你可以使用LIKE子句代替等号 =。
* LIKE 通常与 % 一同使用，类似于一个元字符的搜索。
* 你可以使用AND或者OR指定一个或多个条件
* 你可以在DELETE或UPDATE 命令中使用WHERE...LIKE 子句来指定条件

## MySQL UNION 操作符
MySQL UNION 操作符用于连接两个以上的 SELECT 语句的结果组合到一个结果集合中。多个 SELECT 语句会删除重复的数据

* 语法
* MySQL UNION 操作符语法格式：
```sql
SELECT expression1, expression2, ... expression_n
FROM tables
[WHERE conditions]
UNION [ALL | DISTINCT]
SELECT expression1, expression2, ... expression_n
FROM tables
[WHERE conditions]
```

* expression1, expression2, ...expression_n: 要检索的列
* tables: 要检索的数据表
* WHERE conditions: 可选，检索条件
* DISTINCT: 可选，删除结果集中重复的数据。默认情况下 UNION 操作符已经删除了重复数据，所以 DISTINCT 修饰符对结果没啥影响
* ALL: 可选，返回所有结果集，包含重复数据

```sql
SELECT country FROM Websites
UNION 
SELECT country FROM apps
ORDER BY country
```

## MySQL 排序
我们知道从 MySQL 表中使用 SQL SELECT 语句来读取数据

如果我们需要对读取的数据进行排序，我们就可以使用MySQL 的 ORDER BY 子句来设定你想按哪个字段哪种方式来进行排序，再返回搜索结果

```sql
SELECT field1, field2,...fieldN FROM table_name1, table_name2... ORDER BY field1 [ASC [DESC][默认ASC]],
[field2...] [ASC [DESC][默认 ASC]]
```

```sql
SELECT * FROM test_tbl ORDER BY runoob_id DESC
```
* 你可以使用任何字段来作为排序的条件，从而返回排序后的查询结果
* 你可以设定多个字段来排序
* 你可以使用 ASC 或 DESC 关键字来设置查询结果是按升序或降序排列。默认情况下，它是按升序排列的
* 你可以添加 WHERE...LIKE 子句来设置条件

## MySQL GROUP BY 语句
GROUP BY 语句根据一个或多个列对结果集进行分组
在分组的列上我们可以使用 COUNT，SUM，AVG，等函数

```sql
SELECT column_name, function(column_name)
FROM table_name 
WHERE column_name operator value 
GROUP BY column_name 

```

## MySQL 连接的使用

在前几章节中，我们已经学会了如何在一张表中读取数据，这是相对简单的，但是在真正的应用中经常需要从多个数据表中读取数据。
本章节我们将向大家介绍如何使用 MySQL 的 JOIN 在两个或多个表中查询数据。
你可以在 SELECT, UPDATE 和 DELETE 语句中使用 Mysql 的 JOIN 来联合多表查询。

JOIN 按照功能大致分为如下三类:

* INNER JOIN(内连接，或等值连接)：获取两个字段匹配关系的记录
* LEFT JOIN (左连接): 获取左表所有记录，即使右表没有对应匹配的记录
* RIGHT JOIN (右连接): 与LEFT JOIN 相反，用于获取右表所有记录，即使左表没有对应匹配的记录

### INNER JOIN

```sql
SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a
INNER JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;
```
等价于
```sql
SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl
a, tcount_tbl b WHERE a.runoob_author = b.runoob_author
```

### LEFT JOIN

```sql
SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl
a LEFT JOIN tcount_tbl b ON a.runoob_author = b.runoob_author
```
该语句会读取左边的数据表 runoob_tbl 的所有选取的字段数据，即便在右侧表 tcount_tbl中 没有对应的 runoob_author 字段值。

## MySql NULL 值处理
我们已经知道 MySQL 使用 SQL SELECT 命令及 WHERE 子句来读取数据表中的数据,但是当提供的查询条件字段为 NULL 时，该命令可能就无法正常工作。
为了处理这种情况，MySQL提供了三大运算符:

* IS NULL: 当列的值是NULL，此运算符返回true
* IS NOT NULL: 当列的值不为NULL，运算符返回true
* <=>: 比较操作符 （不同于=操作符），当比较的两个值为NULL，运算符返回true


## MySQL 正则表达式
MySQL可以通过 LIKE ...% 来进行模糊匹配
MySQL 同样也支持其他正则表达式的匹配， MySQL中使用 REGEXP 操作符来进行正则表达式匹配

<table class="reference">
    <tbody>
        <tr>
            <th style="width:30%">模式</th>
            <th style="width:70%">描述</th>
        </tr>
        <tr>
            <td style="width:30%">^</td>
            <td style="width:70%">匹配输入字符串的开始位置。如果设置了 RegExp 对象的 Multiline 属性，^ 也匹配 '\n' 或 '\r' 之后的位置。</td>
        </tr>
        <tr>
            <td style="width:30%">$</td>
            <td style="width:70%">匹配输入字符串的结束位置。如果设置了RegExp 对象的 Multiline 属性，$ 也匹配 '\n' 或 '\r' 之前的位置。</td>
        </tr>
        <tr>
            <td style="width:30%">.</td>
            <td style="width:70%">匹配除 "\n" 之外的任何单个字符。要匹配包括 '\n' 在内的任何字符，请使用象 '[.\n]' 的模式。</td>
        </tr>
        <tr>
            <td style="width:30%">[...]</td>
            <td style="width:70%">字符集合。匹配所包含的任意一个字符。例如， '[abc]' 可以匹配 "plain" 中的 'a'。</td>
        </tr>
        <tr>
            <td style="width:30%">[^...]</td>
            <td style="width:70%">负值字符集合。匹配未包含的任意字符。例如， '[^abc]' 可以匹配 "plain" 中的'p'。</td>
        </tr>
        <tr>
            <td style="width:30%">p1|p2|p3</td>
            <td style="width:70%">匹配 p1 或 p2 或 p3。例如，'z|food' 能匹配 "z" 或 "food"。'(z|f)ood' 则匹配 "zood" 或 "food"。</td>
        </tr>
        <tr>
            <td style="width:30%">*</td>
            <td style="width:70%">匹配前面的子表达式零次或多次。例如，zo* 能匹配 "z" 以及 "zoo"。* 等价于{0,}。</td>
        </tr>
        <tr>
            <td style="width:30%">+</td>
            <td style="width:70%">匹配前面的子表达式一次或多次。例如，'zo+' 能匹配 "zo" 以及 "zoo"，但不能匹配 "z"。+ 等价于 {1,}。</td>
        </tr>
        <tr>
            <td style="width:30%">{n}</td>
            <td style="width:70%">n 是一个非负整数。匹配确定的 n 次。例如，'o{2}' 不能匹配 "Bob" 中的 'o'，但是能匹配 "food" 中的两个 o。</td>
        </tr>
        <tr>
            <td style="width:30%">{n,m}</td>
            <td style="width:70%">m 和 n 均为非负整数，其中n &lt;= m。最少匹配 n 次且最多匹配 m 次。</td>
        </tr>
    </tbody>
</table>

* 查找name字段中以 'st' 为开头的所有数据:
```sql
SELECT name FROM person_tbl WHERE name REGEXP '^st';
```

* 查找 name 字段中以 'ok' 为结尾的所有数据
```sql
SELECT name FROM person_tbl WHERE name REGEXP 'ok$';
```

* 查找name字段中包含'mar'字符串的所有数据：
```sql
 SELECT name FROM person_tbl WHERE name REGEXP 'mar';
```

* 查找name字段中以元音字符开头或以'ok'字符串结尾的所有数据：

```sql
SELECT name FROM person_tbl WHERE name REGEXP '^[aeiou]|ok$';
```

## MySQL 事务

MySQL 事务主要用于处理操作量大，复杂度高的数据。比如说，在人员管理系统中，你删除一个人员，你即需要删除人员的基本资料，也要删除和该人员相关的信息，如信箱，文章等等，这样，这些数据库操作语句就构成一个事务

* 在 MySQL 中只有使用了 lnnodb 数据库引擎的数据库或表才支持事务
* 事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么都不执行
* 事务用来管理 insert, update, delete 语句

一般来说，事务是必须满足4个条件（ACID）：：原子性(Atomicity, 或称不可分割性)、一致性(Consistency)、隔离性(Isolation, 又称独立性)、持久性(Durability)

* 原子性：一个事务(transaction) 中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会陪回滚（Rollback）到事务开始前的状态，就像这个事务从来没有被执行过一样

* 一致性：在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精度、串联性以及后续数据库可以自发性完成预定的工作

* 隔离性: 数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读取未提交（Read uncommitted）、读提交 (read committed)、可重复读(repeatable read)和串行化(Serializable)

* 持久性：事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失

在 MySQL 命令行的默认设置下，事务都是自动提交的，即执行SQL语句后就会马上执行 COMMIT 操作。因此要显式地开启一个事务务须使用命令 BEGIN 或 START TRANSACTION，或者执行命令 SET AUTOCOMMIT = 0，用来禁止使用当前回话的自动提交

### 事务控制语句：

* BEGIN 或 START TRANSACTION 显式地开启一个事务
* COMMIT 也可以使用 COMMIT WORK，不过二者是等价的。COMMIT 会提交事务，并使已对数据库进行的所有修改成永久性的；
* ROLLBACK 也可以使用 ROLLBACK WORK，不过二者是等价的。回滚会结束用户的事务，并撤销正在进行的所有未提交的修改；
* SAVEPOINT identifier，SAVEPOINT 允许在事务中创建一个保存点，一个事务中可以有多个 SAVEPOINT；
* RELEASE SAVEPOINT identifier 删除一个事务的保存点，当没有指定的保存点时，执行该语句会抛出一个异常；
* ROLLBACK TO identifier 把事务回滚到标记点；
* SET TRANSACTION 用来设置事务的隔离级别。InnoDB 存储引擎提供事务的隔离级别有READ UNCOMMITTED、READ COMMITTED、REPEATABLE READ 和 SERIALIZABLE。

### MySQL 事务处理主要有两种方法:
1)  用 BEGIN，ROLLBACK，COMMIT 来实现
* BEGIN 开始一个事务
* ROLLBACK 事务回滚
* COMMIT 事务确认

2) 直接使用 SET 来改变 MySQL 的自动提交模式:
* SET AUTOCOMMIT = 0 禁止自动提交
* SET AUTOCOMMIT = 1 开启自动提交

## MySQL 索引
MySQL 索引的建立对于 MySQL 的高效运行时很重要的，索引可以大大提高MySQL的检索速度。
打个比方，如果合理的设计且使用索引的MySQL是一辆兰博基尼的话，那么没有设计和使用索引的MySQL就是一个人力三轮车。
拿汉语字典的目录页（索引）打比方，我们可以按拼音、笔画、偏旁部首等排序的目录（索引）快速查找到需要的字。
索引分单列索引和组合索引。单列索引，即一个索引只包含单个列，一个表可以有多个单列索引，但这不是组合索引。组合索引，即一个索引包含多个列。

创建索引时，你需要确保该索引是应用在	SQL 查询语句的条件(一般作为 WHERE 子句的条件)。

实际上，索引也是一张表，该表保存了主键与索引字段，并指向实体表的记录。

上面都在说使用索引的好处，但过多的使用索引将会造成滥用。因此索引也会有它的缺点：虽然索引大大提高了查询速度，同时却会降低更新表的速度，如对表进行INSERT、UPDATE和DELETE。因为更新表时，MySQL不仅要保存数据，还要保存一下索引文件。

建立索引会占用磁盘空间的索引文件。

## MySQL 临时表
MySQL 临时表在我们需要保存一些临时数据时是非常有用的。临时表只在当前连接可见，当关闭连接时，Mysql会自动删除表并释放所有空间。

临时表在MySQL 3.23版本中添加，如果你的MySQL版本低于 3.23版本就无法使用MySQL的临时表。

MySQL临时表只在当前连接可见，如果你使用PHP脚本来创建MySQL临时表，那每当PHP脚本执行完成后，该临时表也会自动销毁。

如果你使用了其他MySQL客户端程序连接MySQL数据库服务器来创建临时表，那么只有在关闭客户端程序时才会销毁临时表，当然你也可以手动销毁。