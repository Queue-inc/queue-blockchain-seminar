<template>
  <div>
    <v-card>
      <v-list two-line>
        <div class="userlist">
          <div v-for="user in users" v-bind:key="user._id">
            <div class="list-content">
              <h3 class="headline mb-0">{{ user.name }}<span class="title">（Coin: {{ user.remain_coin }}）</span></h3>
              <div class="subheading">{{ user.public_key }}</div>  
            </div>
            <v-divider></v-divider>
          </div>
        </div>
      </v-list>
    </v-card>
  </div>
</template>
<script>
import getUtil from "@/util/getUtil"
import encoding from '@/util/encoding'
import nacl from 'tweetnacl'

export default {
  data () {
    return {
      users: [],
      secretKey: null
    }
  },
  methods: {
    getUsers() {
      const secret = encoding.hex2ab(this.secretKey)
      const publicKey = encoding.toHexString(nacl.sign.keyPair.fromSecretKey(secret).publicKey).toUpperCase()

      getUtil(
        '/users',
        {
          publicKey: nacl.util.encodeBase64(encoding.hex2ab(publicKey)) // PubKey
        }
      ).then((data) => {
        this.users = data.data
      })
    }
  },
  created () {
    this.secretKey = localStorage.getItem('secretKey')
    this.getUsers()
  }
}
</script>
<style>
.userlist {
  text-align: left;
}
.list-content {
  margin: 20px;
}
</style>
