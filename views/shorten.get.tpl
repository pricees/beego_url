
<form action="/" method="post">
  <input type="text" name="longurl" />
  <input type="submit" value="Shorten!" />
</form>

<div>
'{{ .longurl }}' => '{{ .shorturl }}'
