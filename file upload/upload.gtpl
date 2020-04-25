<html>
  <head>
    <title>Upload file</title>
    <style>
        form{
          text-align: center;
          padding: 50px;
        }
        #buton {
              height: 55px;
              width: 78px;
              font-size: 21px;
        }
    </style>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="/upload"
      method="post"
    >
      <input style="font-size: 27px;" type="file" name="uploadfile" />
      <input type="hidden" name="token" value="{{.}}" />
      <input id="buton" type="submit" value="upload" />
    </form>
  </body>
</html>
