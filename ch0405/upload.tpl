<html>
    <head>
        <title>文件上传</title>
    </head>
    <body>
        <!-- 
            要使表单能够上传文件，首先第一步就是要添加form的enctype属性，enctype属性有如下三种情况:
            application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
            multipart/form-data	  不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
            text/plain	  空格转换为 "+" 加号，但不对特殊字符编码。
        -->
        <form enctype="multipart/form-data" action="/up" method="POST">
            上传文件：<input type="file" name="upfile" /> <br/>
            <input type="hidden" name="token" value="{{.}}"/>
            <input type="submit" value="上传" />
        </form>

    </body>
</html>