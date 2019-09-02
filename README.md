# ImGo is a service written in Go to download, convert and distribute downloaded images.

# Motivation
ImGo is an ambiguous abbreviation for “Images in Go”, although rather it might means “I'm going to learn GoLang”  
This name has been chosen coz it's my first own project in Go after more than 15 years of experience in C++. I like C++, and I hope that this language will have new features, such as reflection, network, etc., although now it is not. It became a cause for me to try something else from programming languages. If I wish to learn something, I'll choose a task that will be close to my real job, and try to solve it in new. Some time ago I was solving a very close task. I implemented an image service in my major programming language. After that I tried to do it in Go.  
It was an interesting experience for me, and I hope that this project might be a step towards learning Go for you like as for me, or this will have become a part of your own web project.

# Introduction
The aim of the service is not to frequently upload images by resources' administrators or topic creators, and majour purpouse is the fast getting stored and cached images from a service by resource users.  
The whole service consists directly of the ImGo service, storage and cache. I like the dependency injection  technique. This is a quite flexible approach for implementing software without hard dependencies and easy to enlarge applications by new features. I'm using this approach through out where I might to do it, but nothing more than enough for my needs. Therefore, in the service, I used dependency injection, and I have a flexible architecture for adding new storage and cache implementations. You can enlarge service by your own storage, cache and other entities implementations for your needs. And I hope that it'll be effortless for you.  

# Features
- Uploading images
- Getting images
- Converting images to the one common format
- Supported storage implementations: PostgreSQL based, MySQL based
- Supported cache implementations: Redis based, Memcached based

# Plans
- Logging
- Daemon mode
- Examples in docker
- Testing

# OS and Compiller
I built this project in Go 1.12 on Ubuntu 16.04 / 18.04, and I hope that project will be able to build in other OS and compler's versions.

# Dependencies
- ImageMagick (for image convertation)

# Build
```bash
git clone https://github.com/tdv/imgo.git  
cd go/src/imgo
go get
go build
```

# Usage
**Note**  
Since the service in its basic configuration uses PostgreSQL-based storage and Redis-based cache, you must have these services in your environment. Of course, you can change the basic options in the configuration file or, perhaps, add your own implementation of one of the interfaces and use it.  
(PostgreSQL database schema in db/postgres/schema.sql)  

**Run**  
After the build you can run application by followed command
```bash
./imgo
```
**Note**  
The service starts with the configuration file, which should be placed next to the application.  
Allowed configuration formats:
- json
- xml
- yaml  

The format is determined by the configuration file extension.  

**Upload images**  
For testing images uploading, you can use curl.  
For example, run the below command from go/src/imgo directory  
```bash
curl -is -XPUT "http://localhost:55555/put?format=JPG" --data-binary @./images/1.jpg
```
For uploading an image with a custom size, you can add extra parameters, like as in the command below.  
```bash
curl -is -XPUT "http://localhost:55555/put?format=JPG&width=600&height=300" --data-binary @./images/1.jpg
```
**Note**  
You have to specify the required format parameter to understand the input format in the image converter (ImageMagick library).  
```bash
Response example:
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Mon, 26 Aug 2019 21:02:44 GMT
Content-Length: 40

5966f327301f3922fce598f0574fa518d492f808
```
The returned identifier is intended to obtain the image, as shown in the example below.  

**Getting images**  
After image uploading, you can get it via a link similar to  
http://localhost:55555/get/5966f327301f3922fce598f0574fa518d492f808

# Tests
Very often for basic load testing I use the ab utility. Such a simple test give me to do basic evaluation of the capabilities of the service.
```bash
ab -c 800 -n 1000000 -r -k "http://localhost:55555/get/645b92e65c697d9f97a61f81b9a8739dd18e5a1f"


Server Software:        
Server Hostname:        localhost
Server Port:            55555

Document Path:          /get/645b92e65c697d9f97a61f81b9a8739dd18e5a1f
Document Length:        86416 bytes

Concurrency Level:      800
Time taken for tests:   111.834 seconds
Complete requests:      1000000
Failed requests:        0
Keep-Alive requests:    0
Total transferred:      86497000000 bytes
HTML transferred:       86416000000 bytes
Requests per second:    8941.79 [#/sec] (mean)
Time per request:       89.468 [ms] (mean)
Time per request:       0.112 [ms] (mean, across all concurrent requests)
Transfer rate:          755310.84 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   12  39.0     10    1013
Processing:     1   78   7.4     78     138
Waiting:        0   13   3.8     13     131
Total:          1   89  39.4     88    1112

```
Let's look at the line from the results
```bash
Requests per second:    8941.79 [#/sec] (mean)
```
I think, that is a good result. If your owns service really used with loads close to 10k rps, you'll have a lot of interesting tasks in future and money :)  

**Many thanks for your attention!**
