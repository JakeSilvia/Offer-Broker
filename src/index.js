import Vue from 'vue';
import Buefy from 'buefy';
import App from './App.vue';
import router from './router';
import ajax from './requests/requests'

Vue.config.productionTip = false;
Vue.use(Buefy, {
  defaultIconPack: 'mdi',
});
let APP = {}
APP.Ajax = ajax
window.APP = APP

Vue.prototype.$ajax = ajax
Vue.prototype.$notification = function (type, message) {
  this.$buefy.toast.open({
    message: message,
    type: type
  })
}


new Vue({
  el: '#app',
  router,
  render: (h) => h(App),
}).$mount('#app');
