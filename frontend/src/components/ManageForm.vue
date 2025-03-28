<template>
  <v-row class="pa-5">
    <v-col cols="12">
      <v-card elevation="8" class="rounded-lg">
        <v-card-title class="text-h5 font-weight-bold">Gerenciamento do sistema</v-card-title>
        <v-card-text>
          <v-row>
            <v-col v-for="item in menuItems" :key="item.title" cols="12" md="4">
              <v-card 
                class="pa-5 card-hover fixed-size d-flex flex-column justify-space-between align-center"
                @click="navigateTo(item.route)"
                elevation="8"
              >
                <v-icon size="40" class="mb-3">mdi-cog</v-icon>
                <v-card-title class="text-center text-h6 font-weight-bold">{{ item.title }}</v-card-title>
                <v-card-text class="text-center">{{ item.description }}</v-card-text>
              </v-card>
            </v-col>
          </v-row>
          <v-row class="mt-10" justify="end">
            <v-col cols="12" md="6">
              <v-btn color="grey darken-2" block outlined @click="$router.go(-1)">Cancelar</v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { useRouter } from "vue-router";
import { ref } from "vue";

export default {
  setup() {
    const router = useRouter();
    const menuItems = ref([
      { title: "Vendedores", description: "Gerencie os vendedores cadastrados", route: "/vendedores" },
      { title: "Entidades", description: "Controle as entidades do sistema", route: "/entidades" },
      { title: "Datas", description: "Acompanhe datas importantes", route: "/datas" }
    ]);

    const navigateTo = (route) => {
      router.push(route);
    };

    return {
      menuItems,
      navigateTo
    };
  }
};
</script>

<style scoped>
.card-hover {
  background: linear-gradient(135deg, #ff4d4d, #ff1a1a);
  color: white;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  cursor: pointer;
  border-radius: 12px;
}

.card-hover:hover {
  transform: scale(1.08);
  box-shadow: 0 12px 24px rgba(255, 26, 26, 0.6);
}

.fixed-size {
  height: 180px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>