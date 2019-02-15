<template>
  <div>
    <v-card>
      <div class="sendfunds">
        <v-text-field
          v-model="name"
          label="Name"
          required
        ></v-text-field>
        <v-text-field
          v-model="secretKey"
          label="Secret Key"
          required
          readonly
        ></v-text-field>
        <v-text-field
          v-model="publicKey"
          label="Public Key"
          required
          readonly
        ></v-text-field>
        <v-btn color="success" @click="request">Sign Up</v-btn>
      </div>
    </v-card>
  </div>
</template>
<script>
import encoding from '@/util/encoding'
import nacl from 'tweetnacl'
import makeRPC from '@/util/rpcUtils'
import txBody from '@/util/txBody'
import bson from 'bson'

export default {
  data () {
    return {
      publicKey: null,
      secretKey: null,
      name: null,
      response: null
    }
  },
  created () {
    this.createKeyPair()
  },
  methods: {
    createKeyPair () {
      let keys = nacl.sign.keyPair()
      this.publicKey = encoding.toHexString(keys.publicKey).toUpperCase()
      this.secretKey = encoding.toHexString(keys.secretKey).toUpperCase()
    },
    request () {
      if ( this.name === '' || this.name === undefined ) {
        window.alert('Please Fill In Your Name.')
        return
      }
      makeRPC(
        txBody.createUser(new bson.ObjectID().toString(), this.name), // enbtityの生成
        nacl.util.encodeBase64(encoding.hex2ab(this.publicKey)), // PubKey
        encoding.hex2ab(this.secretKey) // SecretKey
      ).then((data) => {
        if (data.result.check_tx.code) {
          window.alert('Something Went Wrong Please Try Again.')
          return 
        }
        this.response = data
        window.localStorage.setItem('secretKey', this.secretKey)
        window.location.reload()
      })
    }
  }
}
</script>
