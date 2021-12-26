# Course Project: Simple Library Managment System, Databases, Fudan University, Spring 2020
项目地址 https://github.com/codeplay0314/Library-Managment-System

---

## 目录
#### [安装配置](#0)
#### [用户手册](#1)
#### [管理员手册](#2)
#### [root手册](#3)
#### [功能实现和解释](#4)

---

## <h3 id="0">安装配置</h3>
- MySQL
- Go
---

## <h3 id="1">用户手册</h3>
### 简介
这是一个简单的图书馆管理程序，您可以借助它进行借书、还书、查询书籍等操作。

使用该系统，你需要一个学生账号。如果没有，请联系图书馆管理员注册。

你可以同时借阅任意本书，从借阅当天开始，每本书默认须在 21 天内归还。超过三本书逾期您的账号将会被封锁，此时您需要联系管理员才能还书和重新激活账号。每本书有三次延期（每次一周）的机会。

请注意使用本系统时，书籍 ID 均为图书馆唯一指定的数字，并非ISBN码。

若书籍遗失，请联系管理员。

### 启动程序
1. 完成安装配置
2. 保证 MySQL 账户及密码与文件 library.go 中 `User` 和 `Password` 保持一致，使得程序能顺利链接数据库
3. 进入SLMS文件夹，运行终端
4. 键入 `go run library.go`
5. 程序运行成功，则显示

```
Welcome to the Library Management System!
Please choose login mode:
1. student
2. admin
3. root
0. exit
Please select a number(0-3):
```

### 登陆账号
1. 键入 `1` 进入学生登陆界面

```
Please log in to a student account
Student ID:
```

2. 输入账号和密码

```
Please log in to a student account
Student ID: 18307130163
Password: ********
```

3. 登录状态  
- 失败

```
Please log in to a student account
Student ID: 18307130163
Password: ********

Login Failed!
Please enter correct ID and password...
```

- 成功：进入功能界面

```
Welcome Codeplay!
What do you want to do?
 1: Borrow a book
 2: Return a book
 3: Search books in library
 4: Query borrow histroy info of a student
 5: Query all overdue books a student has borrowed
 6: Query all books a student has borrowed and not returned yet
 7: Check the deadline of a borrowed book
 8: Extend the deadline of returning a book
 0: Log out
You need to do(0-8):
```

### 选择功能
1. 键入相应数字选择功能  
 1: 借阅一本书  
 2: 退还一本书  
 3: 查询书籍的信息  
 4: 查询此账户的借书信息  
 5: 查询此账户过期的借阅书籍  
 6: 查询此账户未退还的借阅书籍  
 7: 查询此账户借阅书籍的应还日期  
 8: 延长此账户借阅书籍的退还日期  
 0: 登出  

2. 键入 `1` 选择开始；键入`0` 选择退出

```
[(function)]
1. Begin
0. Exit
You want to(1/0):  
```

### 功能介绍
####  1 借阅一本书
1. 键入书籍唯一识别号 ID 号

```
[Borrow a book]
Please enter ID number of the book: 2

```

2. 查看信息并确认

```
[Borrow a book]


Please check
-------------
[ID]  2
[ISBN]  9780739360392
[Title]  Harry Potter and the Deathly Hallows
[Author]  J. K. Rowling
[Status]  available
[Info]  borrowed by student 233
-------------
Are you sure to borrow this book? (y/N)

```

3. 键入 `y` 确认， `N` 取消，区分大小写

#### 2 退还一本书
1. 查看所有已借阅书籍

```
[Return a book]


Please check the book you have borrowed:

NO.1
-------------
[Record ID]  11
[Book ID]  1
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2019-01-08
[Expected Return Time]  2019-01-29
[Delayed Times]  0
[Status]  being used
[Info]  
-------------


NO.2
-------------
[Record ID]  13
[Book ID]  3
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2018-05-08
[Expected Return Time]  2018-05-29
[Delayed Times]  0
[Status]  being used
[Info]  
-------------


NO.3
-------------
[Record ID]  14
[Book ID]  4
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2020-03-18
[Expected Return Time]  2020-04-08
[Delayed Times]  0
[Status]  being used
[Info]  
-------------
You want to return book NO.

```

2. 选择并键入要退还的书籍的号码（以NO.2）为例

```
You want to return book NO.2
Are you sure to return the book NO.2? (y/N)
```

3. 键入 `y` 确认， `N` 取消，区分大小写

#### 3 查询书籍的信息
1. 选择查询模式并键入相应号码

```
[Query info of a book in library]
1. Search by book ID
2. Search by book ISBN
3. Search by book title
4. Search by author
0. Exit
You want to(1/0):

```

2. 输入查询信息（以按作者查询为例）

```
Please enter author name:
J. K. Rowling
```

3. 查看所有书籍信息

```
[Query info of a book in library]


All books found:
-------------
[ID]  1
[ISBN]  9780739360385
[Title]  Harry Potter and the Deathly Hallows
[Author]  J. K. Rowling
[Status]  borrowed
[Info]  borrowed by student 233
-------------

-------------
[ID]  2
[ISBN]  9780739360392
[Title]  Harry Potter and the Deathly Hallows
[Author]  J. K. Rowling
[Status]  borrowed
[Info]  borrowed by student 18307130163
-------------

-------------
[ID]  3
[ISBN]  9780747532743
[Title]  Harry Potter and the Sorcerer's Stone
[Author]  J. K. Rowling
[Status]  available
[Info]  borrowed by student 233
-------------

-------------
[ID]  4
[ISBN]  9780439064866
[Title]  Harry Potter and the Chamber of Secrets
[Author]  J. K. Rowling
[Status]  borrowed
[Info]  borrowed by student 233
-------------

Type enter to conitnue...

```

#### 4 查询此账户的借书信息
1. 查看

```
[Query borrow histroy info of a student]


Your borrow history:
-------------
[Record ID]  6
[Book ID]  2
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2020-05-07
[Returned Time]  2020-05-07
[Expected Return Time]  2020-05-28
[Delayed Times]  0
[Status]  returned
[Info]  
-------------

-------------
[Record ID]  8
[Book ID]  2
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2020-04-18
[Returned Time]  2020-05-08
[Expected Return Time]  2020-05-09
[Delayed Times]  0
[Status]  returned
[Info]  
-------------

-------------
[Record ID]  11
[Book ID]  1
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2019-01-08
[Expected Return Time]  2019-01-29
[Delayed Times]  0
[Status]  being used
[Info]  
-------------

-------------
[Record ID]  13
[Book ID]  3
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2018-05-08
[Returned Time]  2020-05-08
[Expected Return Time]  2018-05-29
[Delayed Times]  0
[Status]  returned
[Info]  
-------------

-------------
[Record ID]  14
[Book ID]  4
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2020-03-18
[Expected Return Time]  2020-04-08
[Delayed Times]  0
[Status]  being used
[Info]  
-------------

Type enter to conitnue...
```

#### 5 查询此账户过期的借阅书籍

```
[Query all overdue books a student has borrowed]


Your overdue borrowing:
-------------
[Record ID]  11
[Book ID]  1
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2019-01-08
[Expected Return Time]  2019-01-29
[Delayed Times]  0
[Status]  being used
[Info]  
-------------

-------------
[Record ID]  14
[Book ID]  4
[Borrowed By] 233 (Student ID)
[Borrowed Time]  2020-03-18
[Expected Return Time]  2020-04-08
[Delayed Times]  0
[Status]  being used
[Info]  
-------------

Type enter to conitnue...
```

#### 6 查询此账户未退还的借阅书籍

```
[Query all books a student has borrowed and not returned yet]


All your borrowed books:
-------------
[Record ID]  15
[Book ID]  2
[Borrowed By] 18307130163 (Student ID)
[Borrowed Time]  2020-05-08
[Expected Return Time]  2020-05-29
[Delayed Times]  0
[Status]  being used
[Info]  
-------------

Type enter to conitnue...
```

#### 7 查询此账户借阅书籍的应还日期
1. 键入书籍 ID

```
[Check the deadline of a borrowed book]
Book ID:  

```
2. 查看

```
[Check the deadline of a borrowed book]
Book ID: 2
The book is expected to be return on 2020-05-29

Type enter to conitnue...
```

#### 8 延长此账户借阅书籍的退还日期
1. 键入书籍 ID

```
[Extend the deadline of returning a book]
Book ID:

```
2. 成功

```
[Extend the deadline of returning a book]
Book ID: 2
Successfully extended! Now the book should be returned by 2020-06-05

Type enter to conitnue...
```

#### 0 登出

---

## <h3 id="2">管理员手册</h3>
### 简介
这是一个简单的图书馆管理程序，您可以借助它进行图书馆管理。

使用该系统，你需要一个管理员账号。如果没有，请联系老板注册。

请注意使用本系统时，书籍 ID 均为图书馆唯一指定的数字，并非ISBN码。

请仔细阅读本手册。

### 启动程序
1. 完成安装配置
2. 保证 MySQL 账户及密码与文件 library.go 中 `User` 和 `Password` 保持一致，使得程序能顺利链接数据库
3. 进入SLMS文件夹，运行终端
4. 键入 `go run library.go`
5. 程序运行成功，则显示

```
Welcome to the Library Management System!
Please choose login mode:
1. student
2. admin
3. root
0. exit
Please select a number(0-3):
```

### 登陆账号
1. 键入 `2` 进入管理员登陆界面

```
Please log in to an admin account
Student ID:
```

2. 输入账号和密码

```
Please log in to an admin account
Admin ID: 1
Password: ******
```

3. 登录状态  
- 失败

```
Please log in to an admin account
Admin ID: 1
Password: ******

Login Failed!
Please enter correct ID and password...
```

- 成功：进入功能界面

```
Welcome Codeplay!
What do you want to do?
 1: Add student accounts
 2: Change student account status
 3: Add book to library
 4: Remove book from library
 5: Search books in library
 6: Query borrow histroy info of a student
 7: Query all overdue books a student has borrowed
 8: Query all books a student has borrowed and not returned yet
 9: Check the deadline of borrowed books
10: Extend the deadline of returning book
 0: Log out
You need to do(0-10):
```

### 选择功能
1. 键入相应数字选择功能  
 1: 添加学生账户  
 2: 更改学生账户状态  
 3: 添加书籍  
 4: 移除书籍  
 5: 查询书籍的信息  
 6: 查询此账户的借书信息  
 7: 查询此账户过期的借阅书籍  
 8: 查询此账户未退还的借阅书籍  
 9: 查询此账户借阅书籍的应还日期  
10: 延长此账户借阅书籍的退还日期  
 0: 登出
2. 键入 `1` 选择开始；键入`0` 选择退出

```
 [(function)]
1. Begin
0. Exit
You want to(1/0):  
```

### 功能介绍
#### 1 添加学生账户
1. 键入信息

```
[Add student accounts]
Student ID: 18307130000
Name: ZHANG San
Password: **********
```

2. 确认

```
[Add student accounts]


Please check:
[ID]: 18307130000
[Name]: ZHANG San
[Password]: **********
Are you sure to add this student account? (y/N)y

Successfully added!
```

#### 2 更改学生账户状态
1. 查看所有学生信息

```
[Change student account status]

Account Lists:
-------------
[ID]  18307130000
[Name]  ZHANG San
[Status]  active
[info]  add by admin 1
-------------

-------------
[ID]  18307130163
[Name]  Codeplay
[Status]  active
[info]  active by admin 1 because Revival
-------------

-------------
[ID]  233
[Name]  233
[Status]  active
[info]  active by admin 0 because to return books
-------------
1. Begin
0. Exit
You want to(1/0):
```

2. 选择学生

```
[Change student account status]
The ID of the student account that you want to change: 18307130163
```

3. 键入信息

```
[Change student account status]
The ID of the student account that you want to change: 18307130163
You want to change the status of this account to: (suspended/active/...) suspended
The reason for changing account status: graduated
```

4. 确认

```
[Change student account status]


Please check:
You want to change the student account from
-------------
[ID]  18307130163
[Name]  Codeplay
[Status]  active
[info]  active by admin 1 because Revival
-------------

to
-------------
[ID]  18307130163
[Name]  Codeplay
[Status]   suspended
[info]  set suspended by admin 1 because graduated
-------------
Do you confirm? (y/N)

```

#### 3 添加书籍
1. 键入信息

```
[Add book to library]
ISBN: 9780262033848
Title: Introduction to Algorithms
Authour: Thomas H. Cormen    
```

2. 确认

```
[Add book to library]


Please check
-------------
[ID]  5
[ISBN]  9780262533058
[Title]  Introduction to Algorithms
[Author]  Thomas H. Cormen
[Status]  available
[Info]  
-------------
Are you sure to add this book? (y/N)

```

#### 4 移除书籍
1. 键入书籍 ID

```
[Remove book from library]
Please enter ID number of the book:  
```

2. 确认

```
[Remove book from library]


Please check
-------------
[ID]  4
[ISBN]  9780439064866
[Title]  Harry Potter and the Chamber of Secrets
[Author]  J. K. Rowling
[Status]  borrowed
[Info]  borrowed by student 233
-------------
Are you sure to remove this book? (y/N)
```

#### 5 查询书籍的信息
[ 参见用户手册 功能 3 ]

#### 6 查询的借书信息
[ 参见用户手册 功能 4 ]

#### 7 查询账户过期的借阅书籍
[ 参见用户手册 功能 5 ]

#### 8 查询账户未退还的借阅书籍
[ 参见用户手册 功能 6 ]

#### 9 查询账户借阅书籍的应还日期
[ 参见用户手册 功能 7 ]
#### 10 延长账户借阅书籍的退还日期
[ 参见用户手册 功能 8 ]

#### 0 登出

## <h3 id="3">root手册</h3>
### 启动程序
1. 完成安装配置
2. 保证 MySQL 账户及密码与文件 library.go 中 `User` 和 `Password` 保持一致，使得程序能顺利链接数据库
3. 进入SLMS文件夹，在 `config.json` 更改账户和密码
4. 运行终端，键入 `go run library.go`
5. 程序运行成功，则显示

```
Welcome to the Library Management System!
Please choose login mode:
1. student
2. admin
3. root
0. exit
Please select a number(0-3):
```

### 登陆账号
1. 键入 `3` 进入管理员登陆界面

```
Please log in to an admin account
Student ID:
```

2. 输入密码

```
Please log in to root account
Password: ******
```

### 选择功能

```
Successfully log in to root account!

Please select service below:
 1: Add admin accounts
 2: Add student accounts
 3: Change admin account status
 4: Change student account status
 5: Add book to library
 6: Remove book from library
 7: Search books in library
 8: Query borrow histroy info of a student
 9: Query all overdue books a student has borrowed
10: Query all books a student has borrowed and not returned yet
11: Check the deadline of borrowed books
12: Extend the deadline of returning book
 0: Log out
You need to do(0-12):

```

## <h3 id="4">功能实现和解释</h3>
### 数据库
本系统使用四个模式存储所有信息

- admin 管理员账户

| #ID | ANAME | PASSWORD | STATUS |  INFO |
|---|---|---|---|---|
| 0 | root | FudanICS#2020| active | root account |
|1| Codeplay| 8hg*(hfauj | suspended | hello |

- students 学生账户

| #ID | SNAME | PASSWORD | STATUS |  INFO |
|---|---|---|---|---|
| 18307130163 | Codeplay | ?fad2FSdda| active | me |
|18307130000| Zhang San| 0duf+dhua2 | suspended | set suspended by admin 0 because of having 4 books overdue |

- books 书籍

| #ID | ISBN | TITLE | AUTHOR | STATUS |  INFO |
|---|---|---|---|---|---|
|1| 9780739360385 | Harry Potter and the Deathly Hallows | J. K. Rowling| borrowed | borrowed by student 233 |
|2| 9787111407010 | Introduction to Algorithms | Thomas H. Cormen, Charles E. Leiserson, Ronald Rivest, Clifford Stein | available |  |
|3 | 9780439064866 |  Harry Potter and the Chamber of Secrets | removed | J. K. Rowling | lost from student 18307130000 |

- borrow_record 借阅记录

| #RECORD | #BID | #STU | BORRORDATE | DELAY |  RETURENDATE |  STATUS | INFO |
|---|---|---|---|---|---|---|---|
|1| 5 | 18307130163 | 2018-03-28 | 1 | 2018-04-12| avaliable |  |
|2| 13 | 18307130000 | 2019-08-30 | 0 | / | removed | lost on 2019-09-08 |
|3 | 234 |  18307130163 | 2020-05-08 | 3 | / | borrowed |good book |

### 程序流程
参见代码 `library.go`
（懒得打了）
### 设计心得
很多想法没及时写下来就忘了
主要将时间花在了交互体验上（虽然这是一个不会有人用的程序），增加了自动跳转、自动检查用户逾期信息并进行封锁等鸡肋功能。  为实现一个小功能可能要花几个小时。
注重程序的鲁棒性，对于用户的违规操作进行了极大限制，排除有可能导致运行错误的可能。  
改进方向：
- 无法避免由直接修改数据库带来的运行错误
- 用户被封锁后无法还书（只能联系管理员还）
- books库和borrow_record在异常情况下有可能出现冲突
- 测试大规模数据可能造成的后果
- 代码分块
