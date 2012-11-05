/*!
  Copyright (c) 2012 Astrata Software
  Written by José Carlos Nieto <xiam@menteslibres.org>

  Licensed under the MIT License
  Redistributions of files must retain the above copyright notice.
*/
(function(a){a.foo=function(){a.foo.pull.apply(a.foo,arguments)};a.foo.pull=function(){var d=[];for(var c=0;c<arguments.length;c++){d.push(arguments[c])}var b=a.foo.root+"?load="+d.join(",");a.holdReady(true);a.ajaxSetup({cache:true});a.getScript(b,function(){a.holdReady(false)});a.ajaxSetup({cache:false})};a.foo.styles=function(){for(var b=0;b<arguments.length;b++){a("head").prepend(a("<link />",{type:"text/css",rel:"stylesheet",href:a.foo.root+arguments[b]}))}};a.foo.plugin=function(c,b){a.fn[c]=function(d){if(b[d]){return b[d].apply(this,Array.prototype.slice.call(arguments,1))}else{if(typeof d==="object"||!d){return b.init.apply(this,arguments)}else{a.error("Method "+d+" does not exist on jQuery."+c)}}}};a.foo.start=function(){var b=a("head script");for(var d=0;d<b.length;d++){var e=a(b[d]).attr("src");var c=e.match(/^(.*)jquery\.foo\.js$/);if(c){if(!a.foo.root){a.foo.root=c[1]||"/"}a.foo.parent=b[d]}}};if(typeof a.pull=="undefined"){a.pull=a.foo.pull}a.foo.start()})(jQuery);