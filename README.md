git clone https://github.com/DMax-w/calculator.git

cd calculator

docker build . -t calc

docker run calc

http://172.17.0.2:8080/api/div?a=25&b=5