<template>
  <v-layout column justify-center align-center>
    <v-flex xs12 sm8 md6>
      <v-card>
        <v-card-title class="headline">Register</v-card-title>
        <v-card-text>
          <form @submit.prevent="registerUser">
            <v-text-field label="Name Surname" v-model="fullName" required></v-text-field>
            <v-text-field label="User Name" v-model="userName" required></v-text-field>
            <v-text-field label="Password" v-model="password" type="password" required></v-text-field>
            <v-btn type="submit">Sign Up</v-btn>Already Have a account? Login
            <nuxt-link to="/login">here</nuxt-link>
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
      fullName: '',
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
    registerUser() {
      const { fullName, userName, password } = this
      const URL = 'http://localhost:8000/api/register'
      this.$axios({
        method: 'POST',
        url: URL,
        data: {
          user_name: userName,
          full_name: fullName,
          password
        }
      })
        .then(res => {
          console.log(res)
          localStorage.setItem('token', res.data.user.token)
          this.$router.push('/app')
        })
        .catch(err => {
          // eslint-disable-next-line
          console.log(err)
        })
    }
  }
}
</script>