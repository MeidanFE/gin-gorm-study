<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>模版继承</title>
    <style>
      * {
        margin: 0;
      }

      .nav {
        height: 50px;
        position: fixed;
        top: 0;
        background-color: burlywood;
      }

      .main {
        margin-top: 50px;
      }

      .menu {
        width: 20%;
        height: 100px;
        position: fixed;
        left: 0;
        background-color: cornflowerblue;
      }
    </style>
  </head>
  <body>
    <div class="nav"></div>
    <div class="main">
      <div class="menu"></div>
      <div class="content">
        <h1>这是Home页面</h1>
        <p>{{.}}</p>
      </div>
    </div>
  </body>
</html>
