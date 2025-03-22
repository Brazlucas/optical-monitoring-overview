<template>
  <v-row class="pa-5">
    <v-col cols="12">
      <v-btn @click="$router.push('/form')" color="primary" class="mt-4 mb-5">
        <v-icon left>mdi-plus</v-icon> Adicionar Novo
      </v-btn>

      <v-card elevation="3">
        <v-card-title class="text-h5">Registros</v-card-title>
        <v-divider></v-divider>
        <v-data-table
          :items="registros"
          :headers="headers"
          item-value="id"
          density="comfortable"
        >
          <template v-slot:item.Data="{ item }">
            {{ formatData(item.Data) }}
          </template>
        </v-data-table>
        <!-- Linha de totais -->
        <!-- <v-footer class="bg-blue-darken-3 white--text">
          <v-row no-gutters class="pa-3">
            <v-col cols="12" md="2"><strong>Total</strong></v-col>
            <v-col cols="12" md="1">{{ totalOrcamentoReceita }}</v-col>
            <v-col cols="12" md="1">{{ totalOrcamentoSemReceita }}</v-col>
            <v-col cols="12" md="1">{{ totalOculosSolar }}</v-col>
            <v-col cols="12" md="1">{{ totalAjuste }}</v-col>
            <v-col cols="12" md="1">{{ totalEntrega }}</v-col>
            <v-col cols="12" md="1">{{ totalAssistencia }}</v-col>
            <v-col cols="12" md="2">—</v-col>
          </v-row>
        </v-footer> -->
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { ref, computed, onMounted } from "vue";
import { ListarRegistros } from "../../wailsjs/go/main/App";

export default {
  setup() {
    const registros = ref([]);

    const headers = [
      { title: "Data", key: "Data" },
      { title: "Vendedor", key: "VendedorNome" },
      { title: "Orçamento c/ Receita", key: "OrcamentoReceita" },
      { title: "Orçamento s/ Receita", key: "OrcamentoSemReceita" },
      { title: "Óculos Solar", key: "OculosSolar" },
      { title: "Ajuste", key: "Ajuste" },
      { title: "Entrega", key: "Entrega" },
      { title: "Assistência", key: "Assistencia" },
      { title: "Anotações", key: "Anotacoes" },
    ];

    const carregarRegistros = async () => {
      registros.value = await ListarRegistros();
    };

    const formatData = (data) => {
      if (!data) return "Data inválida";
      const date = new Date(data);
      return date.toLocaleDateString("pt-BR", {
        year: "numeric",
        month: "2-digit",
        day: "2-digit",
      });
    };

    // Computed properties para calcular totais
    const totalOrcamentoReceita = computed(() =>
      registros.value.reduce((sum, r) => sum + r.OrcamentoReceita, 0)
    );

    const totalOrcamentoSemReceita = computed(() =>
      registros.value.reduce((sum, r) => sum + r.OrcamentoSemReceita, 0)
    );

    const totalOculosSolar = computed(() =>
      registros.value.reduce((sum, r) => sum + r.OculosSolar, 0)
    );

    const totalAjuste = computed(() =>
      registros.value.reduce((sum, r) => sum + r.Ajuste, 0)
    );

    const totalEntrega = computed(() =>
      registros.value.reduce((sum, r) => sum + r.Entrega, 0)
    );

    const totalAssistencia = computed(() =>
      registros.value.reduce((sum, r) => sum + r.Assistencia, 0)
    );

    onMounted(carregarRegistros);

    return {
      registros,
      headers,
      formatData,
      totalOrcamentoReceita,
      totalOrcamentoSemReceita,
      totalOculosSolar,
      totalAjuste,
      totalEntrega,
      totalAssistencia,
    };
  },
};
</script>
