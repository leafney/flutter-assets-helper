## protobuf 编译

### golang

```
cd pkg

protoc --go_out=./ ./protocol/message.proto
```

或者，使用命令：

```
cd <project_dir>

./pkg/protocol/build.sh
```

### vue3 + ts

```
cd web

pbjs -t static-module --es6 -w es6 -o ./src/proto/proto.js ../pkg/protocol/message.proto

pbts -o ./src/proto/index.d.ts ./src/proto/proto.js
```

或者，在 `package.json` 文件中添加：

```
  "scripts": {
    "proto": "pbjs -t static-module --es6 -w es6 -o src/proto/proto.js ../pkg/protocol/message.proto && pbts -o src/proto/index.d.ts src/proto/proto.js"
  },
```

然后执行命令：

```
cd <project_dir/web>

pnpm run proto
```

---


