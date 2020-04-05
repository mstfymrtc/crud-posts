<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm8 md6>
      <v-card>
        <v-card-title class="headline">Login Here</v-card-title>
        <v-card-text>
          <form @submit.prevent="loginUser">
            <v-text-field v-model="userName" label="Username" required />
            <v-text-field v-model="password" label="Password" type="password" required />
            <v-btn type="submit">Sign In</v-btn>Want to Register? Register
            <nuxt-link to="/register">here</nuxt-link>
          </form>
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>
<script>
export default {
  data() {
    return {
      userName: '',
      password: ''
    }
  },
  mounted() {
    //auth middleware kullanılması gerekiyor, temp çözüm
    const token = localStorage.getItem('token')
    if (token) {
      this.$router.push('/app')
    }
  },
  methods: {
    loginUser() {
      const { userName, password } = this
      const URL = 'http://localhost:8000/api/authenticate'
      this.$axios({
        method: 'POST',
        url: URL,
        data: {
          user_name: userName,
          password
        }
      })
        .then(res => {
          if (!res.data.status) {
            throw res
            return
          }
          localStorage.setItem('token', res.data.user.token)
          this.$router.push('/app')
        })
        .catch(err => {
          alert(err.data.message)
          // eslint-disable-next-line
          console.log(err)
        })
    }
  }
}
</script>