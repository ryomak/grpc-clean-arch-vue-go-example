<template>
  <div class="about">
    <h1>This is an about page</h1>
  </div>
</template>
<script>
import { RoomHandlerClient } from '@/assets/room_grpc_web_pb.js'
import { RoomReqest } from '@/assets/room_pb.js'
export default {
  name: 'about',
  components: {},
  data: function () {
    return {
      inputField: '',
      data:{},
      client:null,
    }
  },
  created: function () {
    // eslint-disable-next-line
    this.client = new RoomHandlerClient('http://localhost:8080', null, null)
    this.updateRoom()
  },
  methods: {
    updateRoom: function () {
      // eslint-disable-next-line
      let roomReqest = new RoomReqest()
      // eslint-disable-next-line
      this.client.GetRoom(roomReqest, {}, (err, response) => {
        this.data = response.toObject()
        console.log(this.data)
      })
    }
  }
}
</script>
