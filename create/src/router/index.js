import Vue from "vue";
import VueRouter from "vue-router";

import Songs from "../pages/songs.vue";
import InputSongs from "../pages/input-songs.vue";

Vue.use(VueRouter);

const router = new VueRouter({
  mode: "history",
  base: __dirname,
  routes: [
    { path: "/", component: Songs },
    { path: "/songs", component: Songs },
    { path: "/input-songs", component: InputSongs }
  ]
});

export default router;
