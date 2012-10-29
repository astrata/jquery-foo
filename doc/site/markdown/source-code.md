# Get the source

Jsfoo is an [Open Source][1] project sponsored by [Astrata Software][2].

The FastCGI daemon is written in [Go][4] and the source code is [available at our github page][3].


## Hacking *$.foo*

We are open to contributions, if you have any just make a pull request ;-).

Quality plugins licensed as Open Source, such as with MIT or GPL, could be
accepted into the project following some simple conventions on directory
structure plus an special ``index.yaml`` file.

### Directory structure

<table class="table">
  <thead>
    <tr>
      <th>Directory</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>jquery-foo-plugins</code></td>
      <td>Project directory</td>
    </tr>
    <tr>
      <td><code>jquery-foo-plugins/some-plugin</code></td>
      <td>Container directory for the <code>some-plugin</code> jQuery plugin.</td>
    </tr>
    <tr>
      <td><code>jquery-foo-plugins/some-plugin/index.yaml</code></td>
      <td>A specially crafted [YAML][5] file that contains plugin data and code paths.</td>
    </tr>
    <tr>
      <td><code>jquery-foo-plugins/some-plugin/1.0.0</code></td>
      <td>Directory that contains the 1.0.0 release of the plugin.</td>
    </tr>
    <tr>
      <td><code>jquery-foo-plugins/some-plugin/2.0.0</code></td>
      <td>Directory that contains the 2.0.0 release of the plugin.</td>
    </tr>
    <tr>
      <td><code>jquery-foo-plugins/another-plugin</code></td>
      <td>Container directory for the <code>another-plugin</code> jQuery plugin.</td>
    </tr>
  </tbody>
</table>

Do not include documentation, HTML pages or demo files into the plugin directory, only
cleanly-coded uncompressed Javascript files, CSS files (that may contain a media directory) and
LICENSE/README files will be accepted.

The ``.min.js`` files will be generated automatically based on the original ``.js`` file, please avoid including them too.

### Format of the ``index.yaml`` file.

```yaml
---
  # Package name.
  name: "Twitter Bootstrap Javascript Plugins"

  # Stable version.
  stable: "2.1.1"

  # Copyright note.
  copyright: "Copyright 2012 Twitter, Inc."

  # Project's license.
  license: "http://www.apache.org/licenses/LICENSE-2.0"

  # Project's author.
  author: "Twitter, Inc."

  # Project's main site.
  url: "http://twitter.github.com/bootstrap"

  # Latest version (may not be stable yet).
  latest: "2.1.1"

  # Available packages.
  packages:
    # Package version.
    2.1.1:
      # Script files.
      source:
        - "2.1.1/bootstrap.js"
    # Package version.
    2.0.1:
      # Script files.
      source:
        - "2.0.1/bootstrap-alert.js"
        - "2.0.1/bootstrap-button.js"
        - "2.0.1/bootstrap-carousel.js"
        - "2.0.1/bootstrap-collapse.js"
        - "2.0.1/bootstrap-dropdown.js"
        - "2.0.1/bootstrap-modal.js"
        - "2.0.1/bootstrap-tooltip.js"
        - "2.0.1/bootstrap-popover.js"
        - "2.0.1/bootstrap-scrollspy.js"
        - "2.0.1/bootstrap-tab.js"
        - "2.0.1/bootstrap-transition.js"
        - "2.0.1/bootstrap-typeahead.js"
```

[1]: http://en.wikipedia.org/wiki/Open_source
[2]: http://astrata.mx
[3]: https://github.com/Astrata/jquery-foo
[4]: http://golang.org
[5]: http://yaml.org
