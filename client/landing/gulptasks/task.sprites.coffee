gulp        = require 'gulp'
spritesmith = require 'gulp.spritesmith'
argv        = require('minimist') process.argv
site        = SITE_NAME or argv.site or 'landing'
base        = "#{__dirname}/.."
tinypng     = require 'gulp-tinypng'

{ BUILD_PATH } = require './helper.constants'
{ log } = require './helper.logger'

module.exports = (pixelRatio = 1, version = '') ->

  suffix = if pixelRatio > 1 then "__#{pixelRatio}x" else ''

  stream = gulp.src "#{base}/site.#{site}/sprites@#{pixelRatio}x/**/*.png"
    .pipe spritesmith
      imgName   : "sprite@#{pixelRatio}x.png"
      cssName   : "sprite@#{pixelRatio}x.styl"
      imgPath   : "/a/site.#{site}/images/sprite@#{pixelRatio}x.png?#{ version }"
      algorithm : 'binary-tree'
      padding   : 5
      cssFormat : 'stylus'
      cssVarMap : require('./helper.namestylusvars').bind stream, suffix

  stream.css
    .pipe gulp.dest "./../site.#{site}/styl/"

  stream.img
    .pipe tinypng 'jmodUCq84QJZAvbQKDxC3ct95PmQQXZi'
    .pipe gulp.dest "#{BUILD_PATH}/images/"

  return stream
