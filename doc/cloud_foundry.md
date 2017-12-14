## Deploying to Cloud Foundry
Once the application is [built](https://github.com/USEPA/USEEIO_API), it can be deployed in a [Cloud Foundry](https://cloudfoundry.org/)
instance using the [binary build pack](https://docs.cloudfoundry.org/buildpacks/binary/index.html).
The script `apidoc/cfdist.bat` compiles the back-end for Linux and packages all the
resources into the folder `apidoc/cfdist` that can be then deployed using the
[Cloud Foundry Command Line Interface](https://docs.cloudfoundry.org/cf-cli/).
Thus, you need to have [Go](https://golang.org/) installed in order to run the
script.

The `cfdist.bat` script expects that the `build` folder exists within the 'apidoc' directory and that
you put the data that should be deployed with the server into the folder
`build/data`, otherwise it will exit with an error message. You should
then be able to run the script. You probably need to change the meta-data in the
`cfdist/manifest.yaml` file before deploying the application but in general the
workflow should look like this:

```bash
cd apidoc
cfdist.bat
cd cfdist
# change meta-data in manifest.yaml
cf login -a api.run.pivotal.io   # login
# deploy the application as binary package:
cf push -b https://github.com/cloudfoundry/binary-buildpack.git
```
