<!DOCTYPE html>
<html lang="en">

  <head>

    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />

    {{ if .IsHome }}
        <title>{{ setting "page/head/title" }}</title>
    {{ else }}
      {{ if .Title }}
        <title>
          {{ .Title }} {{ if setting "page/head/title" }} // {{ setting "page/head/title" }} {{ end }}</title>
      {{ else }}
        <title>{{ setting "page/head/title" }}</title>
      {{ end }}
    {{ end }}

    <link rel="shortcut icon" href="/favicon.ico" />

    <!-- Le HTML5 shim, for IE6-8 support of HTML elements -->
    <!--[if lt IE 9]>
    <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->

    <script type="text/javascript" src="http://get.jsfoo.org/jquery.js"></script>
    <script type="text/javascript" src="http://get.jsfoo.org/jquery.foo.js"></script>

    <link rel="stylesheet" href="http://static.hckr.org/normalize/normalize.css" />

    <link rel="stylesheet" href="http://static.hckr.org/bootstrap/css/bootstrap.css" />
    <link rel="stylesheet" href="http://static.hckr.org/bootstrap/css/bootstrap-responsive.css" />

    <link rel="stylesheet" href="http://static.hckr.org/google-code-prettify/prettify.css" />
    <script type="text/javascript" src="http://static.hckr.org/google-code-prettify/prettify.js"></script>

    <link rel="stylesheet" href="/css/style.css" />

    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <script type="text/javascript">
      $(document.body).ready(
        function() {
          // Code (marking code blocks for prettyPrint)
          var code = $('code');

          for (var i = 0; i < code.length; i++) {
            var el = $(code[i])
            var className = el.attr('class');
            if (className) {
              el.addClass('prettyprint');
              el.addClass('lang-'+className);
            }
          };

          // An exception, LaTeX blocks.
          var code = $('code.latex');

          for (var i = 0; i < code.length; i++) {
            var el = $(code[i])
            var img = $('<img>', { 'src': 'http://phibin.com/api/render?snippet='+encodeURIComponent(el.html()) });
            img.insertBefore(el);
            el.hide();
          };

          // Starting prettyPrint.
          prettyPrint();

          // Tables without class

          $('table').each(
            function(i, el) {
              if (!$(el).attr('class')) {
                $(el).addClass('table');
              };
            }
          );

          // Navigation
          var links = $('ul.menu li').removeClass('active');

          for (var i = 0; i < links.length; i++) {
            var a = $(links[i]).find('a');
            if (a.attr('href') == document.location.pathname) {
              $(links[i]).addClass('active');
            };
          };

        }
      );
    </script>

    <style type="text/css">
      .navbar .brand {
        margin-left: 0px;
      }
      body {
        padding-top: 50px;
      }
    </style>

  </head>

  <body>

    <div class="container" id="container">

      <div class="navbar navbar-fixed-top navbar-inverse">
        <div class="navbar-inner">
          <div class="container">

            <a class="brand" href="{{ url "/" }}">{{ setting "page/brand" }}</a>

            <div class="nav-collapse">
              {{ if settings "page/body/menu" }}
                <ul id="nav" class="nav menu">
                  {{ range settings "page/body/menu" }}
                    <li>{{ link .url .text }}</li>
                  {{ end }}
                </ul>
              {{ end }}
            </div>

            <ul id="nav" class="nav pull-right menu">
              <li><a id="astrata-link" href="http://astrata.mx" target="_blank" />Astrata</a></li>
            </ul>

          </div>
        </div>
      </div>


    {{ if .IsHome }}

      <div class="project-header hero-unit">

        <h1>jQuery.foo</h1>

        <p>
          The <code>jQuery.foo</code> plugin lets you include many great <a href="/plugins">jQuery plugins</a> from the cloud
          into your web app with ease.
        </p>

      </div>

      <div class="container-fluid">
        <div class="row">
          <div class="span12">
            {{ .ContentHeader }}

            {{ .Content }}

            {{ .ContentFooter }}
          </div>
        </div>
      </div>

    {{ else }}

      {{ if .BreadCrumb }}
        <ul class="breadcrumb menu">
          {{ range .BreadCrumb }}
            <li><a href="{{ url .link }}">{{ .text }}</a> <span class="divider">/</span></li>
          {{ end }}
        </ul>
      {{ end }}

      <div class="container-fluid">

        <div class="row">
          {{ if .SideMenu }}
            {{ if .Content }}
              <div class="span3">
                  <ul class="nav nav-list menu">
                    {{ range .SideMenu }}
                      <li>
                        <a href="{{ url .link }}">{{ .text }}</a>
                      </li>
                    {{ end }}
                  </ul>
              </div>
              <div class="span8">
                {{ .ContentHeader }}

                {{ .Content }}

                {{ .ContentFooter }}
              </div>
            {{ else }}
              <div class="span12">
                {{ if .CurrentPage }}
                  <h1>{{ .CurrentPage.text }}</h1>
                {{ end }}
                <ul class="nav nav-list menu">
                  {{ range .SideMenu }}
                    <li>
                      <a href="{{ url .link }}">{{ .text }}</a>
                    </li>
                  {{ end }}
                </ul>
              </div>
            {{ end }}
          {{ else }}
            <div class="span12">
              {{ .ContentHeader }}

              {{ .Content }}

              {{ .ContentFooter }}
            </div>
          {{ end }}
        </div>

      </div>

    {{ end }}

    {{ if setting "page/body/scripts/footer" }}
      <script type="text/javascript">
        {{ setting "page/body/scripts/footer" | jstext }}
      </script>
    {{ end }}

    <hr />

    <footer>
      <i>jsfoo</i> by <a href="http://astrata.mx" target="_blank">astrata</a>.
    </footer>

    <script type="text/javascript">

      var _gaq = _gaq || [];
      _gaq.push(['_setAccount', 'UA-29633552-1']);
      _gaq.push(['_trackPageview']);

      (function() {
        var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
        ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
        var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
      })();

    </script>

  </body>
</html>
