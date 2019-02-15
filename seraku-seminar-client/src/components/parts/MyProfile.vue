<template>
  <div>
    <v-card>
      <div class="profile">
        <h3 class="headline mb-0">{{ name }}<span class="title">（Coin: {{ coin }}）</span></h3>
        <div class="subheading">{{ address }}</div>  
      </div>
    </v-card>
  </div>
</template>
<script>
import getUtil from '@/util/getUtil'
import encoding from '@/util/encoding'
import nacl from 'tweetnacl'

export default {
  data () {
    return {
      name: 'TEMP',
      address: 'TEMP',
      coin: 100,
      secretKey: null
    }
  },
  created () {
    this.secretKey = localStorage.getItem('secretKey')
    this.signIn()
  },
  methods: {
    signIn() {
      const secret = encoding.hex2ab(this.secretKey)
      const publicKey = encoding.toHexString(nacl.sign.keyPair.fromSecretKey(secret).publicKey).toUpperCase()
      window.localStorage.setItem('secretKey', this.secretKey)
      getUtil(
        '/user',
        {
          publicKey: nacl.util.encodeBase64(encoding.hex2ab(publicKey)) // PubKey
        }
      ).then((data) => {
        this.name = data.data.name
        this.address = data.data.public_key
        this.coin = data.data.remain_coin
      })
    }
  }
}
</script>
<style>
.profile {
  margin: 0 50px;
  padding: 10% 0;
  text-align: left;
  overflow-wrap: break-word;
  height: 284px;
}
</style>
