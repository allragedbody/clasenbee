<!DOCTYPE <html>
  <head>
    <meta charset="utf-8">
    <title>创建配置模板</title>
  </head>
  <body>
  <div>
    <form action="/web/createmoduleresult" method="post" enctype ="multipart/form-data" >
      <table>
        <tr>
          <td> 模板名称: </td>
          <td><input  name="name" value=""/></td>
        </tr>
        <tr>
        <td>模板中文名称:</td>
        <td><input  name="chname" value=""/></td>
        </tr>
        <tr>
        <td>模板目录:</td>
        <td><input  name="path" value=""/></td>
        </tr>
        <tr>
        <td>模板备注:</td>
        <td><input  name="desc" value=""/></td>
        </tr>
        <tr>
        <td><input type="submit" name="Button1" value="提交" id="Button1" /></td>
        </tr>
      </table>
    </form>
    {{template "footer.html"}}
  </div>
  </body>
</html>>