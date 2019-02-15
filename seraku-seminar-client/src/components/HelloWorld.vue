<template>
  <div class="hello">
    <v-app>
      <h1 class="wallet display-3">Wallet App</h1>
      <div v-if="toggle">
        <v-container grid-list-md text-xs-center>
          <v-layout row wrap>
            <v-flex xs6>
              <div>
                <profile></profile>
              </div>
            </v-flex>
            <v-flex xs6>
              <div>
                <send-funds></send-funds>
              </div>
            </v-flex>
            <v-flex xs12>
              <div>
                <users></users>
              </div>
            </v-flex>
          </v-layout>
        </v-container>
      </div>
      <div v-else>
        <div>
          <sign-up></sign-up>
        </div>
      </div>
    </v-app>
  </div>
</template>

<script>
import Profile from '@/components/parts/MyProfile'
import SendFunds from '@/components/parts/SendFunds'
import SignUp from '@/components/parts/SignUp'
import Users from '@/components/parts/Users'
export default {
  components: {
    Profile,
    SendFunds,
    SignUp,
    Users
  },
  name: 'HelloWorld',
  data () {
    return {
      toggle: false
    }
  },
  methods: {
  },
  created () {
    if (window.localStorage.getItem('secretKey')) {
      this.toggle = true
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.wallet {
  margin-top: 30px;
}
</style>
