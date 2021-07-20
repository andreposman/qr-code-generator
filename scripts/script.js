import http from "k6/http";
import { check} from "k6";

export let options = {
  discardResponseBodies: false,
  scenarios: {
    ext_controlled_scenario: {
      executor: 'externally-controlled',
      exec: 'createQRCode',
      vus: 0,
      maxVUs: 100,
      duration: '10m',
    },
  },
};
    // healthcheck_scenario: {
    //   executor: 'constant-vus',
    //   exec: 'healthcheck',
    //   vus: 10,
    //   duration: '10s',
    // },
    // ramping_qrcode_creation_scenario: {
    //   startTime: '30s',
    //   executor: 'ramping-vus',
    //   exec: 'rampingQRCodeCreation',
    //   startVUs: 0,
    //   stages: [
    //     { duration: '20s', target: 50 },
    //     { duration: '20s', target: 100 },
    //     { duration: '20s', target: 150 },
    //     { duration: '20s', target: 50 },
    //     { duration: '5s', target: 0 },
    //   ],
    //   gracefulRampDown: '5s',
    // },
// export function healthcheck() {
//   let response = http.get('http://localhost:9000/ping')
//     check(response, {'Response status code was 200': (resp) => resp.status == 200})
// }

// export function createQRCode() {
//   const url = 'http://localhost:9000/qrcode/create'
//   const payload = JSON.stringify({
//     data: "https://github.com/andreposman",
//     file_name: "qrcode-k6-test",
//   })
//   const params = {
//     headers: {
//       "Content-Type": "application/json",
//     },
//   }

//   let response = http.post(url, payload, params)

//   check(response, {'Response status code was 200': (resp) => resp.status == 200})
// }

export function createQRCode() {
  const url = 'http://192.168.15.30/qrcode/create'
  const payload = JSON.stringify({
    data: "https://github.com/andreposman",
    file_name: "qrcode-k6-ramping-test",
  })
  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  }

  let response = http.post(url, payload, params)

  check(response, {'Response status code was 200': (resp) => resp.status == 200})
}
