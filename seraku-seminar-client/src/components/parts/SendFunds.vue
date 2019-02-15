<template>
  <div>
    <v-card>
      <div class="sendfunds">
        <v-text-field
          v-model="to"
          label="To"
          required
        ></v-text-field>
        <v-text-field
          v-model="amount"
          label="Amount"
          required
        ></v-text-field>
        <v-btn color="success" @click="sendFunds">Send</v-btn>
      </div>
    </v-card>
  </div>
</template>
<script>
import encoding from '@/util/encoding'
import nacl from 'tweetnacl'
import makeRPC from '@/util/rpcUtils'
import txBody from '@/util/txBody'

export default {
  data () {
    return {
      to: '',
      amount: 0
    }
  },
  methods: {
    sendFunds() {
      const key = localStorage.getItem('secretKey')
      if (!key) {
        return
      }
      const secret = encoding.hex2ab(key)
      const publicKey = encoding.toHexString(nacl.sign.keyPair.fromSecretKey(secret).publicKey).toUpperCase()
      makeRPC(
        txBody.sendFunds(this.to, this.amount), // entityの生成
        nacl.util.encodeBase64(encoding.hex2ab(publicKey)), // PubKey
        secret // SecretKey
      ).then((data) => {
        if (data.result.check_tx.code) {
          window.alert('Something Went Wrong Please Try Again.')
          return 
        }
        this.response = data
      }).catch ((e) => {
        window.alert(e)
      })
    }
  }
}
</script>
<style>
.sendfunds {
  margin: 0px 50px;
  padding: 50px 0px;
}
</style>
