import encoding from './encoding'
import config from '@/config/config'
import request_promise from 'request-promise'
import nacl from 'tweetnacl'
nacl.util = require('tweetnacl-util')

export default async function makeRPC (txBody, publicKey, secret) {
  let signature = nacl.sign.detached(nacl.util.decodeUTF8(JSON.stringify(txBody)), secret)

  let tx = {
    body: JSON.stringify(txBody),
    signature: encoding.toHexString(signature),
    publicKey: publicKey
  }

  let base64Data = nacl.util.encodeBase64(encoding.str2ab(JSON.stringify(tx)))

  let headers = {
    'Content-Type': 'text/plain',
    'Accept': 'application/json-rpc'
  }

  let options = {
    url: config.rpcUrl,
    method: 'POST',
    headers: headers,
    json: true,
    body: {
      jsonrpc: '2.0',
      method: 'broadcast_tx_commit',
      params: {
        tx: base64Data
      },
      id: 'anything'
    }
  }
  return request_promise(options)
}
