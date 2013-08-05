<div class="container">
  <div class="row">
  <div class="col-lg-6 col-offset-3">
      <h2>Shorten that url!</h2>
      <form action="/" method="post">
        <fieldset>
          <div class="form-group">
            <label for="longurl">Long url:</label>
            <input type="text" id="longurl" name="longurl" class="form-control" placeholder="{{.longurl}}" />
          </div>
          <button class="btn btn-defualt" type="submit">Shorten!</button>
        </fieldset>
      </form>
      {{if compare .shorturl nil}}
      {{else}}
      <p class="lead">Shortened Url: 
        <a href="{{ .shorturl }}" target="_blank">{{ .shorturl }}</a></p>
      {{end}}
    </div>
  </div>
<div>
