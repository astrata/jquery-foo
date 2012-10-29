## Unleash *$.foo*

Add ``jquery.js`` and ``jquery.foo.js`` to the ``HEAD`` of your template.

```html
<!-- Loading the jQuery library -->
<script type="text/javascript" src="http://get.jsfoo.org/jquery.js"></script>

<!-- Loding the jQuery.foo plugin -->
<script type="text/javascript" src="http://get.jsfoo.org/jquery.foo.js"></script>
```

## Pull the plugins you need

Use ``$.foo.pull`` to load jQuery plugins into your web app.

```html
<!-- Add to HEAD -->
<script type="text/javascript">
  // This will load Fancybox and jQuery-UI
  $.foo.pull('fancybox', 'ui');
  // NOTE: If any plugin depends on another it will be pulled too, there's
  // no need to solve dependencies manually.
</script>
```


## Initialize your web app

Your web app will continue loading normally while the plugins are downloaded in the background, you can rely on the
well-known ``$(document).ready()`` function to initialize your plugins when they're all loaded.

```html
<!-- Add to BODY -->
  <script type="text/javascript">
    $(document).ready(
      function() {
        // Add your code here :-).
      }
    );
  </script>
</body>
```

