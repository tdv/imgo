# ImGo is an image service written in Go.

# Introduction
ImGo might be considered like an ambiguous abbreviation for “Images in Go” or “I'm going to learn GoLang” (your choice). The major purpose of the service is to be fast for getting stored images rather fast uploading.

# Features
- Upload, download and convert images
- Supported: PostgreSQL, MySQL, Redis, Memcached
- Based on: ImageMagick and Go standard library

# OS and Compiller
The project was built with Go 1.18 on Ubuntu 20.04. Hopefully, the project will be able to build within other OS and compiler version.

# Dependencies
- ImageMagick

# Build
```bash
git clone https://github.com/tdv/imgo.git  
cd imgo
go build .
```
**Note**
If you have a build issue on Ubuntu related on ImageMagick, you can try to install the package with a command like below
```bash
sudo apt-get install libmagickwand-dev
```

# Usage
**Note**  
Ones the service was built you need to have the installed PostgreSQL and Redis on your workstation. Nevertheless, that is possible to try all in one out of the box having used a self-contained solution based on docker and docker-compose (see the folder 'docker'). Definitely recommended!

**Upload images**  
For testing image uploads, you can use curl.  
```bash
curl -is -XPUT "http://localhost:55555/put?format=JPG" --data-binary @./images/1.jpg
```
To upload an image with a custom size, you can add extra parameters
```bash
curl -is -XPUT "http://localhost:55555/put?format=JPG&width=600&height=300" --data-binary @./images/1.jpg
```
**Note**  
You need to specify the format parameter for the converter (ImageMagick library).  
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
After the image was uploaded, you can get the one by a link similar to  
http://localhost:55555/get/5966f327301f3922fce598f0574fa518d492f808

# Tests
The ab utility might be used to assess the speed.
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
I think, that is a good result. If your service really used with loads close to 10k rps, you'll have a lot of interesting tasks...
