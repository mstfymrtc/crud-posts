<template>
  <v-layout column justify-center align-center>
    <v-card class="mb-3" width="100%">
      <v-card-text>
        <form @submit.prevent="addPost">
          <v-textarea v-model="content" label="Content" required />
          <v-btn color="success" type="submit">Add Post</v-btn>
        </form>
      </v-card-text>
    </v-card>
    <v-card class="mb-3" width="100%">
      <v-data-table
        :rows-per-page-items="[10, 20, 30, 40]"
        disable-initial-sort
        :headers="headers"
        :items="listData"
        class="elevation-0"
      >
        <template slot="items" slot-scope="props">
          <tr>
            <td>
              <span>{{new Date(props.item.created_at).toLocaleDateString("tr")}}</span>
            </td>

            <td class="text-truncate" style="max-width:200px">{{ props.item.content }}</td>
            <td>
              <v-btn color="info" @click="showPostDetails(props.item.id)">DETAILS</v-btn>
              <v-btn
                v-if="currentUserId==props.item.author_id"
                color="error"
                @click="deletePost(props.item.id)"
              >DELETE</v-btn>
            </td>
          </tr>
        </template>
      </v-data-table>
    </v-card>
    <v-row justify="center">
      <v-dialog v-model="dialog" max-width="500px">
        <v-card class="mx-auto" style="overflow-x:hidden" color="info" dark>
          <v-card-actions>
            <v-list-tile class="grow">
              <v-list-tile-avatar color="grey darken-3">
                <v-img
                  class="elevation-6"
                  src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                ></v-img>
              </v-list-tile-avatar>

              <v-list-tile-content>
                <v-list-tile-title>{{post&&post.author&&post.author.full_name}}</v-list-tile-title>
                <v-list-tile-text
                  style="font-size:11px"
                >{{post&&post.author&&post.author.user_name}}</v-list-tile-text>
              </v-list-tile-content>

              <v-layout align-center justify-end>
                <span class="subheading">{{new Date(post.created_at).toLocaleString("tr")}}</span>
              </v-layout>
            </v-list-tile>
          </v-card-actions>
          <v-card-text
            style="word-break:break-all"
            class="headline font-weight-bold"
          >{{ post.content }}</v-card-text>
        </v-card>
      </v-dialog>
    </v-row>
  </v-layout>
</template>
<script>
import decode from 'jwt-decode'
export default {
  data() {
    return {
      currentUserId: 0,
      dialog: false,
      post: {},
      content: '',
      listData: [],
      headers: [
        {
          text: 'Creation Time',
          value: 'created_at'
        },
        {
          text: 'Content',
          value: 'content',
          sortable: false
        },
        {
          text: 'Action',
          value: 'action',
          sortable: false
        }
      ]
    }
  },
  mounted() {
    //auth middleware kullanılması gerekiyor, temp çözüm
    const token = localStorage.getItem('token')
    if (!token) {
      localStorage.removeItem('token')
      this.$router.push('/login')
    }
    var decodedToken = decode(token)
    this.currentUserId = decodedToken.UserId

    this.fetchPosts()
  },
  methods: {
    fetchPosts() {
      const token = localStorage.getItem('token')
      const URL = 'http://localhost:8000/api/posts'
      this.$axios({
        method: 'GET',
        url: URL,
        headers: {
          Authorization: `Bearer ${token}`,
          accept: 'application/json'
        }
      })
        .then(res => {
          this.listData = res.data.data
        })
        .catch(err => {
          console.log(err)
        })
    },
    fetchPost(id) {
      const token = localStorage.getItem('token')
      const URL = `http://localhost:8000/api/posts/${id}`
      return this.$axios({
        method: 'GET',
        url: URL,
        headers: {
          Authorization: `Bearer ${token}`,
          accept: 'application/json'
        }
      })
        .then(res => Promise.resolve(res))
        .catch(err => Promise.reject(err))
    },

    addPost() {
      // eslint-disable-next-line
      const { content } = this
      const data = { content }
      const token = localStorage.getItem('token')
      const URL = 'http://localhost:8000/api/posts'
      this.$axios({
        method: 'POST',
        url: URL,
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`
        },
        data: data
      })
        .then(_ => {
          // this.$router.push('/app')
          this.listData = [_.data.post, ...this.listData]
          this.content = ''
        })
        .catch(err => {
          // eslint-disable-next-line
          console.log(err)
        })
    },
    deletePost(id) {
      const token = localStorage.getItem('token')
      const URL = `http://localhost:8000/api/posts/${id}`
      this.$axios({
        method: 'DELETE',
        url: URL,
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`
        }
      })
        .then(_ => {
          this.fetchPosts()
        })
        .catch(err => {
          // eslint-disable-next-line
          console.log(err)
        })
    },
    showPostDetails(id) {
      this.fetchPost(id)
        .then(res => {
          this.post = res.data.data
          this.dialog = true
        })
        .catch(err => alert(err))
      // alert(item)
    }
  }
}
</script>
