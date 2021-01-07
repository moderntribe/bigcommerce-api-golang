import gulp from 'gulp';
import cleaner from 'gulp-clean';
import { exec } from 'child_process';
import log from 'fancy-log';
import mustache from 'gulp-mustache';
import rename from 'gulp-rename';
import download from 'gulp-download2';
import fs from 'fs';
import { find } from 'lodash';
import merge from 'gulp-merge-json';

const apis = [
  {
    name: 'Wishlists',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/wishlists/wishlists.v3.json',
    moduleName: 'wishlists',
  },
  {
    name: 'Widgets',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/widgets/widgets.v3.json',
    moduleName: 'widgets',
    patch: 'widgets.patch',
  },
  {
    name: 'Themes',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/themes/themes.v3.json',
    moduleName: 'themes',
  },
  // segfault
  // {
  //   name: 'Catalog',
  //   url: 'https://developer.bigcommerce.com/api-reference/store-management/catalog/catalog.v3.json',
  //   moduleName: 'catalog',
  // },
  {
    name: 'Subscribers',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/subscribers/subscribers.v3.json',
    moduleName: 'subscribers',
  },
  {
    name: 'StoreInfo',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/store-information-api/store_information.v2.json',
    moduleName: 'storeinfo',
    patch: 'storeinfo.patch',
  },
  {
    name: 'Scripts',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/scripts/scripts.v3.json',
    moduleName: 'scripts',
    patch: "scripts.patch",
  },
  {
    name: 'PriceLists',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/price-lists/price_lists.v3.json',
    moduleName: 'pricelists',
    patch: 'pricelists.patch',
  },
  {
    name: 'OrdersV2',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/orders/orders.v2.oas2.json',
    moduleName: 'ordersv2',
  },
  {
    name: 'OrdersV3',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/order-transactions/orders.v3.json',
    moduleName: 'ordersv3',
    patch: 'ordersv3.patch',
  },
  {
    name: 'Sites',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/sites/sites.v3.json',
    moduleName: 'sites',
    patch: 'sites.patch',
  },
  {
    name: 'Channels',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/channels/channels.v3.json',
    moduleName: 'channels',
    patch: 'channels.patch',
  },
  {
    name: 'Carts',
    url: 'https://developer.bigcommerce.com/api-reference/store-management/carts/carts.v3.json',
    moduleName: 'carts',
    patch: 'carts.patch',
  },
];

const install = api => done => {
  // gem build clients/StoreInfo/swagger_client_store_info.gemspec
  // gem install --user-install clients/StoreInfo/swagger_client_store_info-1.0.0.gem
  // add this to the Gemfile: gem 'swagger_client_store_info', '~> 1.0.0'
  return exec(
    `rm -f clients/${api.name}/${api.gemName}-1.0.0.gem \
    && gem build clients/${api.name}/${api.gemName}.gemspec \
    && gem install --user-install clients/${api.name}/${api.gemName}-1.0.0.gem`,
    { cwd: '.' },
    (err, stdout, stderr) => {
      log(stdout);
      log(stderr);
      done(err);
    },
  );
};

const clean = api => (
  () => (
    gulp
      .src([
          `clients/${api.name}/*`,
          `!clients/${api.name}/node_modules`,
          `!clients/${api.name}/test`,
        ],
        { read: false, allowEmpty: true })
      .pipe(cleaner())
  )
)

const gomodule = api => done => {
  console.info(api);
  if (!fs.existsSync(`clients/${api.moduleName}`)) {
    fs.mkdirSync(`clients/${api.moduleName}`);
  }
  return exec(
    `/usr/local/go/bin/go mod init bigcommerce.com/apis/clients/${api.moduleName}`,
    { cwd: `clients/${api.moduleName}`, env: { GO111MODULE: "on", HOME: "/users/woodj", GOPATH: "/user/woodj/go" } },
    (err, stdout, stderr) => {
      log(stdout);
      log(stderr);
      done(err);
    },
  );
}

const codegen = api => done => {
  console.info(api);
  return exec(
    `/usr/local/bin/swagger generate client \
    -f ${api.url} \
    --skip-validation \
    -t clients/${api.moduleName}`,
    { cwd: '.' },
    (err, stdout, stderr) => {
      log(stdout);
      log(stderr);
      done(err);
    },
  );
};

const processTestSupportFiles = api => () => (
  gulp.src('./resources/test/**/*.mustache')
    .pipe(mustache('./gulpfile.config.json'))
    .pipe(rename(path => ({
      dirname: path.dirname,
      basename: path.basename,
      extname: ""
    })))
    .pipe(gulp.dest(`clients/${api.name}/test`))
);

const processSourceFiles = api => () => {
  // all the apis and models
  const apiFilenames = fs.readdirSync(`clients/${api.name}/src/api`);
  const modelFilenames = fs.readdirSync(`clients/${api.name}/src/model`);
  const files = [
    ...apiFilenames.map(f => ({ name: f.replace('.js', ''), path: `./api/${f.replace('.js', '')}`})),
    ...modelFilenames.map(f => ({ name: f.replace('.js', ''), path: `./model/${f.replace('.js', '')}`})),
  ];

  // index file
  gulp.src('./resources/src/index.js.mustache')
    .pipe(mustache({ files }))
    .pipe(rename(path => ({
      dirname: path.dirname,
      basename: path.basename,
      extname: ""
    })))
    .pipe(gulp.dest(`clients/${api.name}/src`))

  return gulp.src(['./resources/src/*.js', '!resources/src/*.mustache'])
    .pipe(gulp.dest(`clients/${api.name}/src`));
};

const copyConfigFiles = api => () => (
  gulp.src('resources/.babelrc')
    .pipe(gulp.dest(`clients/${api.name}`))
);

const postCleanup = api => () => (
  gulp
    .src([
        `clients/${api.name}/.travis.yml`,
        `clients/${api.name}/git_push.sh`,
        `clients/${api.name}/mocha.opts`,
      ],
      { read: false, allowEmpty: true })
    .pipe(cleaner())
);

const applyPatch = api => (done) => {
  // duplicate the clients/{api.name} dir to clients/{api.name}.orig
  // make the fix in clients/{api.name}
  // diff -ur {api.name}.orig {api.name} > ../resources/patches/{api.name}.patch
  // to patch an unfixed src: cd {api.name} && patch -p1 -i ../resources/patches/{api.name}.patch
  if (api.patch) {
    exec(
      `patch -p1 -i ../../resources/patches/${api.patch}`,
      { cwd: `./clients/${api.name}` },
      (err, stdout, stderr) => {
        log(stdout);
        log(stderr);
        done(err);
      },
    )
  } else {
    done();
  }
};

const buildClean = (api) => {
  return gulp.series(
    Object.assign(clean(api), { displayName: 'clean' }),
    Object.assign(gomodule(api), { displayName: 'gomodule' }),
    Object.assign(codegen(api), { displayName: 'codegen' }),
    // Object.assign(processSourceFiles(api), { displayName: 'processSourceFiles' }),
    // Object.assign(processTestSupportFiles(api), { displayName: 'processTestSupportFiles' }),
    // Object.assign(copyConfigFiles(api), { displayName: 'copyConfigFiles' }),
    Object.assign(applyPatch(api), { displayName: 'applyPatch' }),
    Object.assign(postCleanup(api), { displayName: 'postCleanup' }),
  );
};

const buildOne = async (done) => {
  const [,,,arg,value] = process.argv;
  if (arg !== '--name' || !value) {
    console.info('Usage: gulp buildOne --name [api.name]');
    return done();
  }
  const api = find(apis, { name: value });
  if (!api) {
    console.warn(`No such api: ${value}`);
    console.info('Usage: gulp buildOne --name [api.name]');
    return done();
  }
  await buildClean(api)();
  return done();
};
const buildAll = gulp.series(apis.map(api => buildClean(api)));

// given some env props, creates a master config, then creates all the test support config files, so they can run
// just keep your gulpfile.config.json around, and you don't have to run this
const testSupport = async () => {
  const config = {
    clientId: process.env.clientId,
    accessToken: process.env.accessToken,
    storeId: process.env.storeId,
  };
  return fs.writeFileSync('gulpfile.config.json', JSON.stringify(config));
};

// export tasks
export {
  buildAll, buildOne, testSupport,
};

export default buildAll;
