import axios from 'axios'
import config from '@/config/config'

export default async function getRequest(
  endpoint,
  urlParams
) {
  let headers = {
    'Content-Type': 'text/plain',
    'Accept': 'application/json-rpc'
  }

  let options = {
    url: config.getUrl + endpoint,
    method: 'GET',
    headers: headers,
    params: urlParams
  }

  return axios(options)
}