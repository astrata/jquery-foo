<?php

/*

  Copyright (c) 2012 <xiam@menteslibres.org>

  Permission is hereby granted, free of charge, to any
  person obtaining a copy of this software and associated
  documentation files (the "Software"), to deal in the
  Software without restriction, including without limitation
  the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the
  Software, and to permit persons to whom the Software is
  furnished to do so, subject to the following conditions:

  The above copyright notice and this permission notice
  shall be included in all copies or substantial portions of
  the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY
  KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
  WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
  PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS
  OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
  OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
  OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
  SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

*/

  require '../conf.php';
  
  header('Content-Type: text/javascript; charset=utf-8');

  $loaded = array();

  function upload($file) {
    require('../lib/S3/S3.php');

    $s3 = new S3(STORAGE_S3_KEY, STORAGE_S3_SECRET);

    $name = basename($file);
    
    if (ENABLE_GZIP) {
      $response = $s3->putObject(
        $s3->inputFile($file),
        STORAGE_S3_BUCKET,
        $name,
        S3::ACL_PUBLIC_READ,
        array(),
        array(
          'Content-Type' => 'text/javascript; charset=utf-8',
          'Content-Encoding' => 'gzip',
          'Vary' => 'Accept-Encoding'
        )
      );
    } else {
      $response = $s3->putObject(
        $s3->inputFile($file),
        STORAGE_S3_BUCKET,
        $name,
        S3::ACL_PUBLIC_READ
      );
    }

    if ($response) {
      return sprintf('http://%s/%s', STORAGE_S3_BUCKET, $name);
    } else {
      return false;
    }
  }
  
  function load_plugin($name, &$space) {

    if (isset($space[$name])) {
      return false;
    } else {
      $space[$name] = true;
    }

    $version = preg_match('/^(.+)\-([0-9\.]+)$/', $name, $match);

    if ($version) {
      $name    = $match[1];
      $version = $match[2];
    } else {
      $version = null;
    }

    if ($name) {

      $path = sprintf('../plugins/%s/index.json', $name);

      if (file_exists($path)) {

        $text = file_get_contents($path);
        $json = json_decode($text, true);

        if ($json) {

          if (!$version) {
            $version = $json['stable'];
          }

          $package =& $json['packages'][$version];

          if (!empty($package)) {

            if (!empty($package['requires'])) {
              foreach ($package['requires'] as $require) {
                load_plugin($require, $space);
              }
            }

            if (!empty($package['script'])) {
              foreach($package['script'] as $script) {
                $space['script'][] = array(
                  'copy'  => sprintf('%s (%s). %s %s', $json['plugin_name'], $version, $json['copyright'], $json['license']),
                  'file'  => sprintf('../plugins/%s/%s', $name, $script)
                );
              }
            }
            
            if (!empty($package['style'])) {
              foreach($package['style'] as $style) {
                $space['style'][] = sprintf('../plugins/%s/%s', $name, $style);
              }
            }

          } else {
            die(sprintf('/* Missing package "%s" */', htmlspecialchars($name)));
          }

        }
      
      }
    }
  }

  $load =& $_GET['load'];

  if (!empty($load)) {

    $space = array(
      'script' => array(),
      'style' => array()
    );
    
    $load = preg_replace('/[^0-9a-zA-Z\-\.,]/', '', $load);

    $etag = md5($load);
 
    if (ENABLE_CACHE) {
      if (isset($_SERVER['HTTP_IF_NONE_MATCH'])) {
        if ($_SERVER['HTTP_IF_NONE_MATCH'] == $etag) {
          header('HTTP/1.0 304 Not Modified');
          exit;
        }
      }
      header('Cache-Control: public');
      header('Expires: '.gmdate('D, d M Y H:i:s', time()+3600).' GMT');
      header('Etag: '.$etag);
    }
    
    $cache_file = CACHE_DIR.$etag.'.js';

    if (!ENABLE_CACHE || file_exists($cache_file) == false) {

      $load = explode(',', $load);
      sort($load, SORT_STRING);

      foreach ($load as $name) {
        load_plugin($name, $space);
      }
      
      if (ENABLE_GZIP) { 
        ob_start('ob_gzhandler');
      } else {
        ob_start();
      }
     
      echo sprintf('$.foo.styles(%s);', json_encode($space['style']));
      echo "\n\n";
      
      foreach ($space['script'] as $file) {
        echo '/* '.$file['copy'].' */';
        echo "\n";
        readfile($file['file']);
        echo "\n\n";
      }

      $buff = ob_get_clean();

      if (ENABLE_GZIP) {
        $size = strlen($buff);

        header('Content-Encoding: gzip');
        header('Vary: Accept-Encoding');

        $crc = crc32($buff);

        $buff = "\x1f\x8b\x08\x00\x00\x00\x00\x00".substr(gzcompress($buff, 9), 0, -4);
        $buff .= pack('V', $crc);
        $buff .= pack('V', $size);
      }

      $fh = fopen($cache_file, 'w');

      if ($fh) {
        fwrite($fh, $buff);
        fclose($fh);
      }

      $url = upload($cache_file);

      if ($url) {
        header('HTTP/1.1 301 Moved Permanently');
        header(sprintf('Location: %s', $url));
      } else {
        echo $buff;
      };

    } else {
    
      header('HTTP/1.1 301 Moved Permanently');
      header(sprintf('Location: http://%s/%s', STORAGE_S3_BUCKET, basename($cache_file)));

      /*
      if (ENABLE_GZIP) {
        header('Content-Encoding: gzip');
        header('Vary: Accept-Encoding');
      }

      readfile($cache_file);
      */
    }

  }
?>
