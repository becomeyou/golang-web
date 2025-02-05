<!DOCTYPE html>

<html>
<head>
  <title>user edit Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>

    <form action="/user/doEdit" method="post">
        用户名: <input type="text" name="username"></br></br>
        密码: <input type="password" name="password"></br></br>
        爱好: <input type="checkbox" value=1 id="lable1" name="hobby"><label for="lable1">吃饭</label>
            <input type="checkbox" value=2 id="lable2" name="hobby"><label for="lable2">睡觉</label>
            <input type="checkbox" value=3 id="lable3" name="hobby"><label for="lable3">打豆豆</label></br></br>
        <input type="submit" value="提交">
    </form>

</body>
</html>