import { createRouter, createWebHistory } from "vue-router";
import Form from "./components/Form.vue";
import ListaRegistros from "./components/Lista.vue";

const routes = [
  { path: "/", component: Form },
  { path: "/registros", component: ListaRegistros },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
