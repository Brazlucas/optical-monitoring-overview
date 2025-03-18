import { createRouter, createWebHistory } from "vue-router";
import Form from "./components/Form.vue";
import ListaRegistros from "./components/Lista.vue";

const routes = [
  { path: "/form", component: Form },
  { path: "/", component: ListaRegistros },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
