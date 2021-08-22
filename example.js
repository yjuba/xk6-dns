import dns from 'k6/x/dns';

export let options = {
    discardResponseBodies: true,
    scenarios: {
        my_awesome_api_test: {
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
