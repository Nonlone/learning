const gulp = require('gulp');

 const browserSync = require('browser-sync').create();// 静态服务器


exports.browserSync= function (){
    browserSync.init({
        server: {
            baseDir: "./"
        }
    });
    gulp.watch("*.html").on("change",browserSync.reload);
};
