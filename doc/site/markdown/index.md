## Quick start

Link <a href="http://get.jsfoo.org/jquery.js" target="_blank">``jquery.js``</a> and <a href="http://get.jsfoo.org/jquery.foo.js" target="_blank">``jquery.foo.js``</a> directly from our server, then use ``$.foo.pull`` to fetch the plugins you need.

All the plugins you pull from ``jQuery.foo``, dependencies and .CSS files will be automatically merged into one single compressed and minified Javascript file before being served.

```html
<head>
  <!-- jQuery library -->
  <script type="text/javascript" src="http://get.jsfoo.org/jquery.js"></script>

  <!-- jQuery.foo plugin -->
  <script type="text/javascript" src="http://get.jsfoo.org/jquery.foo.js"></script>

  <script type="text/javascript">
    // Loading Fancybox and jQuery-UI
    $.foo.pull('fancybox', 'ui');
  </script>
</head>
<body>
  <script type="text/javascript">
    $(document).ready(
      function() {
        // At this point both plugins are now be available.
      }
    );
  </script>
</body>
</code>
```

<!--

## Alternate local setup

A remote loading of the script may not be always convenient, an alternate setup of ``jQuery.foo`` includes copying and running it in your own PHP-capable server and CDN.

Download the ``jquery-foo`` package from github and update the plugins submodule.

<code>
cd ~/www/cdn.example.com
git clone git://github.com/xiam/jquery-foo.git jquery-foo
cd jquery-foo
git submodule update --init --recursive
</code>

Make the ``temp`` directory writable by the webserver.

<code>
chmod -R 777 temp
</code>

Run the ``Makefile`` to generate the ``.min.js`` files (it uses [YUI Compressor](http://developer.yahoo.com/yui/compressor/)).

<code>
make
</code>

Copy the ``conf.php.example`` file into ``conf.php`` and edit it to suit your needs.

<code>
cp conf.php.example conf.php
vim conf.php
</code>

Point your website's ``webroot`` (or ``documentroot``) to ``jquery-foo/public``.

It's done! now you have a local running instance of ``$.foo``.
-->

## API reference

### $.foo.pull(plugin1, plugin2, ...)

Loads the given list of plugins. An updated list of available plugins can be always found at <a href="https://github.com/xiam/jsfoo-plugins" target="_blank">plugin repository</a>.

<code>
$.foo.pull('fancybox', 'bootstrap');
</code>


### $.foo.plugin(methods, name)

Shortcut for creating jQuery plugins.

<code>
(
  function($) {
    var methods = {
      'init': function() {
        console.log('Hello world!');
      }
    };

    $.foo.plugin(methods, 'hello');
  }
)(jQuery);

$.hello();
</code>


