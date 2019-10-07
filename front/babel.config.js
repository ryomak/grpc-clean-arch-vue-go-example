module.exports = {
  presets: [
    '@vue/app',
    ['env',{
      'modules':'commonjs'
    }]
  ],
  'plugins':['add-module-exports']
}
