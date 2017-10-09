const clean = require('gulp-clean');
const gulp = require('gulp');

gulp.task('default', ['clean'], () => {
    require('bootprint')
        .load(require('bootprint-openapi'))
        .merge({ /* Any other configuration */ })
        .build('apidoc.yaml', 'build')
        .generate();
});

gulp.task('clean', () => {
    return gulp.src(['build'], {
        read: false
    }).pipe(clean({ force: true }));
});
