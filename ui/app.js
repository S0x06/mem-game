const path = require('path');
const views = require('koa-views');
const router = require('koa-router')();
const logger = require('koa-logger');
const server = require('koa-static');
const Koa = require('koa');

const app = new Koa();

app.use(logger());
app.use(views(path.join(__dirname, '/views'), {extension: 'ejs'}));
app.use(server(path.join(__dirname, '/node_modules')));
app.use(server(path.join(__dirname, '/statics')));
app.use(server(path.join(__dirname)));

router.get('/', index);

app.use(router.routes());

async function index(ctx) {
  await ctx.render("index")
}

console.log('listen 3000');
app.listen(3000);
