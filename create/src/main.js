import Vue from "vue";
import router from "./router";
import VueMaterial from "vue-material";
import "vue-material/dist/vue-material.min.css";
import "vue-material/dist/theme/default-dark.css";

import App from "./App.vue";

Vue.config.productionTip = false;

Vue.use(VueMaterial);

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
