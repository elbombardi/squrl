package routes

const RESPONSE_404 = `<!DOCTYPE html>
<html>
<head>
  <title>404 Page Not Found</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      text-align: center;
      background-color: #f2f2f2;
      margin: 0;
      padding: 50px;
    }

    h1 {
      color: #333;
    }

    p {
      color: #666;
    }

    .container {
      max-width: 600px;
      margin: 0 auto;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>Oops! Page Not Found</h1>
    <p>The page you are looking for could not be found.</p>
    <p>Please check the URL or navigate back <form>
      <input type="button" value="<" onclick="history.back()">
     </form>
     </p>
  </div>
</body>
</html>
`

const RESPONSE_500 = `<!DOCTYPE html>
<html>
<head>
  <title>500 Internal Server Error</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      text-align: center;
      background-color: #f2f2f2;
      margin: 0;
      padding: 50px;
    }

    h1 {
      color: #333;
    }

    p {
      color: #666;
    }

    .container {
      max-width: 600px;
      margin: 0 auto;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>Oops! Internal Server Error</h1>
    <p>Sorry, something went wrong on the server.</p>
    <p>Please try again later or contact support.</p>
  </div>
</body>
</html>
`
