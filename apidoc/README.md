# apidoc
The files in this folder generate the documentation files of the USEEIO API from
a [Swagger (OpenAPI)](https://swagger.io/specification/) document:
[apidoc.yaml](./apidoc.yaml). You can directly copy and edit the `apidoc.yaml`
file in the [Swagger Editor](http://editor.swagger.io/) (currently it shows some
semantic errors). We use [Gulp](https://gulpjs.com/) and 
[bootprint-openapi](https://github.com/bootprint/bootprint-openapi) to build the
HTML page (with CSS and JavaScript) from this specification. So if you have
[npm]() and [Gulp](https://gulpjs.com/) installed, you can switch to the `apidoc`
folder, install the respective node modules, and finally run the build:

```bash
cd apidoc
npm install
gulp
```

The API documentation files are then located in the `apidoc/build` folder.
