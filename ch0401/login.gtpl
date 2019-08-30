<html>
    <head>
        <title>Form登录</title>
    </head>
    <body>
        <!-- action表示把form提交到 /login -->
        <form action="/login" method="POST">
            用户名：<input type="text" name="username" /> <br/>
            密码：<input type="password" name="password" /> <br/>
            <input type="submit" value="登录" />
        </form>
    </body>
</html>