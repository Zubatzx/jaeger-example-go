# Jaeger implementation on Go Language

Try on your local:
1. Run jaeger all-in-one on your local.

  ```
  docker run -d -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest
  ```
  
  
  Once the container starts, open http://127.0.0.1:16686/ in the browser to access the Jaeger UI.


2. Run all cmd/http/main.go on all microservices.

Available url:
---
  - Showtime -> /showtime?id=1
  - Showname -> /showname?id=1
  - Book     -> /book?id=1


3. Check Jaeger UI for output



Credits:
---
- https://medium.com/@masroor.hasan/tracing-infrastructure-with-jaeger-on-kubernetes-6800132a677
- https://medium.com/opentracing/take-opentracing-for-a-hotrod-ride-f6e3141f7941
- https://medium.com/velotio-perspectives/a-comprehensive-tutorial-to-implementing-opentracing-with-jaeger-a01752e1a8ce
- https://medium.com/@carlosedp/instrumenting-go-for-tracing-c5bdabe1fc81
- https://github.com/jaegertracing/jaeger-kubernetes
- https://github.com/jaegertracing/jaeger/tree/master/examples/hotrod
- https://github.com/yurishkuro/opentracing-tutorial
