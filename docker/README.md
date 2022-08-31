**There is a little info about docker-containerization of the ImGo project for the demo or tests.**  

# Motivation
Some scripts for docker-containerization of the ImGo project have been added especially for the demo and test avoiding any --like postgres or other DB, Rredis or Memcached, etc-- installation some related components on the developer's workstation. There was choose to use docker-compose and beforehanded built docker images independently. That provides a little freedom. It is possible to build and run images from docker-compose.yml, but that will be more rely only on docker-compose. Independently built the images might be used somewhere else, for instance under k8s at least.  

# Build images
Being in the catalog 'docker'  you need to run only build.sh which is located on the same level with docker-compose.yml file (not in components subdirs) like below
```bash
./build.sh
```
That is all!

# Run and test
Now you can run and test whole service with the dependent components (postgres and redis).
```bash
docker-compse up
```
Wait for a little while and try to test or might be to use :)
To stop press 'Ctrl + C'  

In order to use that in a daemon mode you can start all by command
```bash
docker-compose up -d
````
Stopping by command
```bash
docker-compose down
````
**How to test and use you can find a few test requests on the main README.md** ([Usage](https://github.com/tdv/imgo/blob/master/README.md#usage))  


# Wrapping up
Good luck and I will appreciate any feedback :)