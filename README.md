# Jaeger implementation on Go Language

Try on local:
1. Run jaeger all-in-one on your local 
	docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest
	Once the container starts, open http://127.0.0.1:16686/ in the browser to access the Jaeger UI.

2. Run all cmd/http/main.go on all microservices
	Available url:
	- Showtime -> /showtime?id=1
	- Showname -> /showname?id=1
	- Book     -> /book?id=1

3. Check Jaeger UI for output


