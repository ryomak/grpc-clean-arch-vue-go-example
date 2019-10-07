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
      num: 0 ,
      client:null,
    }
  },
  created: function () {
    // eslint-disable-next-line
    this.client = new RoomHandlerClient('http://localhost:8080', null, null)
    this.getTotalNum()
  },
  methods: {
    getTotalNum: function () {
      // eslint-disable-next-line
      let getRequest = new getTotalNumParams()
      // eslint-disable-next-line
      this.client.getTotalNum(getRequest, {}, (err, response) => {
        this.num = response.toObject()
        console.log(this.num)
      })
    },
    addNum: function () {
      // eslint-disable-next-line
      let request = new addNumParams()
      request.setNumber(Number(this.inputField))
      // eslint-disable-next-line
      this.client.addNum(request, {}, (err, response) => {
        this.inputField = ''
        this.num = response.toObject()
      })
    }
  }
}
</script>
