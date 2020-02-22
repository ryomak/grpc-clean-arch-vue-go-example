# grpc-cleanarch-vue-go-example
BE: grpc + clearn architecture + Go
FE: grpc-web + Vuejs

## simple gps map (WIP)
display member's coordinate of same room

<img width="685" alt="スクリーンショット 2019-10-13 17 52 21" src="https://user-images.githubusercontent.com/21288308/67107581-b06ae880-f207-11e9-95e8-3def795f9514.png">


### I/F
```
/radar/:roomname/:nickname
ex)/radar/roomA/ryoma
```

## USAGE
1. vue proto build
``` $ cd front && sh protobuild.sh && cd ..  ```

2. go proto build
``` cd server && make proto-build && cd ..```

3. go db setting
```$ make migrate/init && make migrate/up ```

4. running server
```$ make run ```

5. running vue-server
```$ npm run serve ``` 

display member point of same room 

## TODO
- [ ] to set docker mysql

## LICENSE
MIT
