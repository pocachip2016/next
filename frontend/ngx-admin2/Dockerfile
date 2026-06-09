FROM node:18-alpine
RUN apk add --no-cache python3 make g++
WORKDIR /app
COPY package*.json ./
RUN npm install --legacy-peer-deps --ignore-scripts && \
    npm uninstall node-sass --legacy-peer-deps && \
    npm install sass@1.69.7 @types/ws@8.5.4 --legacy-peer-deps --save-dev && \
    rm -rf /app/node_modules/@angular-devkit/build-angular/node_modules/@babel/runtime && \
    node -e "const fs=require('fs'); \
      const dirs=['/app/node_modules/@babel/runtime/helpers/esm','/app/node_modules/@babel/runtime/helpers']; \
      const esmContent='export default function regeneratorValues(o){var it=typeof Symbol!==\"undefined\"&&o[Symbol.iterator];return it?it.call(o):o;}'; \
      const cjsContent='\"use strict\";Object.defineProperty(exports,\"__esModule\",{value:true});exports[\"default\"]=regeneratorValues;function regeneratorValues(o){var it=typeof Symbol!==\"undefined\"&&o[Symbol.iterator];return it?it.call(o):o;}'; \
      if(!fs.existsSync(dirs[0]+'/regeneratorValues.js'))fs.writeFileSync(dirs[0]+'/regeneratorValues.js',esmContent); \
      if(!fs.existsSync(dirs[1]+'/regeneratorValues.js'))fs.writeFileSync(dirs[1]+'/regeneratorValues.js',cjsContent);"
COPY . .
EXPOSE 4200
CMD ["npx", "ng", "serve", "--host", "0.0.0.0", "--poll", "2000", "--disable-host-check"]
