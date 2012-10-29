/*!
  Copyright (c) 2012
  Written by José Carlos Nieto <xiam@menteslibres.org>

  Licensed under the MIT License
  Redistributions of files must retain the above copyright notice.
*/
(function(b){b.foo=function(){b.foo.pull.apply(b.foo,arguments)};b.foo.pull=function(){var e=[];for(var f=0;f<arguments.length;f++){e.push(arguments[f])}var a=b.foo.root+"?load="+e.join(",");b.holdReady(true);b.getScript(a,function(){b.holdReady(false)})};b.foo.styles=function(){for(var a=0;a<arguments.length;a++){b("head").prepend(b("<link />",{type:"text/css",rel:"stylesheet",href:b.foo.root+arguments[a]}))}};b.foo.plugin=function(d,a){b.fn[d]=function(c){if(a[c]){return a[c].apply(this,Array.prototype.slice.call(arguments,1))}else{if(typeof c==="object"||!c){return a.init.apply(this,arguments)}else{b.error("Method "+c+" does not exist on jQuery."+d)}}}};b.foo.start=function(){var a=b("head script");for(var g=0;g<a.length;g++){var f=b(a[g]).attr("src");var h=f.match(/^(.*)jquery\.foo\.js$/);if(h){if(!b.foo.root){b.foo.root=h[1]||"/"}b.foo.parent=a[g]}}};if(typeof b.pull=="undefined"){b.pull=b.foo.pull}b.foo.start()})(jQuery);