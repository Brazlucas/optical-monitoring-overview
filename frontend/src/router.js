import { createRouter, createWebHistory } from "vue-router";
import Form from "./components/Form.vue";
import ListaRegistros from "./components/Lista.vue";
import ManageFormVue from "./components/ManageForm.vue";

const routes = [
  { path: "/form", component: Form },
  { path: "/manage-form", component: ManageFormVue },
  { path: "/", component: ListaRegistros },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
