xk6-dns
=====

This is a [k6](https://github.com/k6io/k6) extension using the
[xk6](https://github.com/k6io/xk6) system.

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- If you're using SQLite, a build toolchain for your system that includes `gcc` or
  another C compiler. On Debian and derivatives install the `build-essential`
  package. On Windows you can use [tdm-gcc](https://jmeubank.github.io/tdm-gcc/).
  Make sure that `gcc` is in your `PATH`.
- Git

Then:

1. Install `xk6`:
  ```shell
  go install github.com/k6io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  xk6 build v0.33.0 --with github.com/yjuba/xk6-dns
  ```

## Example

```javascript
// example.js
import dns from 'k6/x/dns';

export let options = {
    discardResponseBodies: true,
    scenarios: {
        my_awesome_dns_test: {
            executor: 'constant-vus',
            vus: 1,
            duration: '5s',
        },
    },
};

export default function () {
    dns.setReadTimeout('1s');
    dns.resolve('127.0.0.1:53', 'example.com.', 'A');
}
```

Result output:

```shell
$ ./k6 run example.js

          /\      |??| /??/   /??/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ??\
   /          \   |  |\  \ |  (?)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: example.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 35s max duration (incl. graceful stop):
           * my_awesome_dns_test: 1 looping VUs for 5s (gracefulStop: 30s)


running (05.0s), 0/1 VUs, 1938 complete and 0 interrupted iterations
my_awesome_api_test ↓ [======================================] 1 VUs  5s

     data_received........: 87 kB 17 kB/s
     data_sent............: 56 kB 11 kB/s
     dns.dial.count.......: 1938  387.409151/s
     dns.request.count....: 1938  387.409151/s
     dns.response.rtt.....: avg=2.04ms min=0s       med=1ms   max=356ms    p(90)=1ms    p(95)=1ms
     iteration_duration...: avg=2.56ms min=794.36μs med=1.5ms max=357.11ms p(90)=1.64ms p(95)=1.68ms
     iterations...........: 1938  387.409151/s
     vus..................: 1     min=1        max=1
     vus_max..............: 1     min=1        max=1
```
