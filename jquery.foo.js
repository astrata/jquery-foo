/*!
  Copyright (c) 2012
  Written by José Carlos Nieto <xiam@menteslibres.org>

  Licensed under the MIT License
  Redistributions of files must retain the above copyright notice.
*/

(
  function($) {

    $.foo = function() {
      $.foo.pull.apply($.foo, arguments);
    };

    /* $.foo.pull(script1, script2, ...) */
    $.foo.pull = function() {
      var names = [];
      
      for (var i = 0; i < arguments.length; i++) {
        names.push(arguments[i]);
      };

      var url = $.foo.root + '?load='+names.join(',');

      $.holdReady(true);

      $.getScript(url, function() {
        $.holdReady(false);  
      });

    };

    /* $.foo.styles(style1, style2, ...) */
    $.foo.styles = function() {
      for (var i = 0; i < arguments.length; i++) {
        $('head').prepend($('<link />', {
          'type': 'text/css',
          'rel': 'stylesheet',
          'href': $.foo.root + arguments[i]
        }));
      };
    }
    
    /* directly taken from http://docs.jquery.com/Plugins/Authoring */
    $.foo.plugin = function(name, methods) {
      $.fn[name] = function(method) {
        if (methods[method]) {
          return methods[method].apply(this, Array.prototype.slice.call(arguments, 1));
        } else if (typeof method === 'object' || !method) {
          return methods.init.apply(this, arguments);
        } else {
          $.error('Method '+method+' does not exist on jQuery.'+name);
        };
      };
    };

    /* startup */
    $.foo.start = function() {
      var scripts = $('head script');
      
      for (var i = 0; i < scripts.length; i++) {
        var src = $(scripts[i]).attr('src');
        var match = src.match(/^(.*)jquery\.foo\.js$/);
        if (match) {
          if (!$.foo.root) {
            $.foo.root    = match[1] || '/';
          };
          $.foo.parent  = scripts[i];
        };
      };
    };
  
    /* alias for $.foo.pull */
    if (typeof $.pull == 'undefined') {
      $.pull = $.foo.pull;
    };

  }
)(jQuery);
