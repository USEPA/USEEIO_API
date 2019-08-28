const clean = require('gulp-clean');
const gulp = require('gulp');

gulp.task('build', (cb) => {
    require('bootprint')
        .load(require('bootprint-openapi'))
        .merge({ /* Any other configuration */ })
        .build('apidoc.yaml', 'build')
        .generate()
        .done(console.log);
    if (cb) {
        cb();
    }
});

gulp.task('clean', () => {
    return gulp.src('build', {
        read: false,
        allowEmpty: true,
    }).pipe(clean({ force: true }));
});

gulp.task('default', gulp.series('clean', 'build'));
