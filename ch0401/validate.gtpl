<html>
    <head>
        <title>Form输入值的Server端验证</title>
    </head>
    <body>
        <form action="/val" method="POST">
        用户名（必填）：<input type="text" name="username" /><br/>
        年龄：<input type="number" name="age" /><br/>
        名字（中文）：<input type="text" name="realname" /> <br/>
        英文名：<input type="text" name="english" /> <br/>
        邮箱：<input type="email" name="email" /> <br/>
        水果：<select name="fruit">
            <option value="apple">苹果</option>
            <option value="pear">梨</option>
            <option value="banana">香蕉</option>
        </select> <br/>
        性别：<input type="radio" name="gender" value="1" />男<input type="radio" name="gender" value="2">女<br/>
        爱好：<input type="checkbox" name="interest" value="football">足球
            <input type="checkbox" name="interest" value="basketball">篮球
            <input type="checkbox" name="interest" value="tennis">网球
            <br/>
        <input type="submit" name="submit" value="提交">
        </form>
    </body>
</html>