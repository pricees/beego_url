<div class="container">
  <h2>Shorten that url!</h2>
  <form action="/" method="post">
    <fieldset>
      <div class="form-group">
        <label for="longurl">Enter a url:</label>
        <input type="text" id="longurl" name="longurl" class="form-control"
        placeholder="http://google.com" />
      </div>
      <button class="btn btn-defualt" type="submit">Shorten!</button>
    </fieldset>
  </form>
<div>
'{{ .longurl }}' => '{{ .shorturl }}'
