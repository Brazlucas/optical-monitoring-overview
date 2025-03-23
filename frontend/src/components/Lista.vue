<template>
  <v-row class="pa-5">
    <v-col cols="12">
      <v-card elevation="3">
        <v-card-title class="text-h5">Registros</v-card-title>
        <v-divider></v-divider>

        <v-row class="ml-1 justify-end" no-gutters>
          <v-select
            v-model="anoFiltro"
            :items="anosDisponiveis"
            label="Filtrar por Ano"
            outlined
            class="mr-10 ml-2"
          ></v-select>

          <v-select
            color="red"
            v-model="periodoFiltro"
            :items="periodos"
            label="Filtrar por intervalo"
            class="mr-10"
          ></v-select>

          <v-btn @click="$router.push('/form')" color="red" class="mt-4 mb-5 mr-2">
            <v-icon left>mdi-plus</v-icon> Novo Registro
          </v-btn>
        </v-row>

        <v-data-table :headers="headers" :items="[]" item-value="id" density="comfortable">
          <template v-slot:body>
            <template v-for="(registros, data) in registrosAgrupados" :key="data">
              <tr>
                <td colspan="12" class="text-center font-weight-bold bg-grey-lighten-2">
                  {{ data }}
                </td>
              </tr>

              <tr v-for="registro in registros" :key="registro.id">
                <!-- <td>{{ formatData(registro.Data) }}</td> -->
                <td>{{ registro.VendedorNome }}</td>
                <td>{{ registro.OrcamentoReceita }}</td>
                <td>{{ registro.OrcamentoSemReceita }}</td>
                <td>{{ registro.OculosSolar }}</td>
                <td>{{ registro.Ajuste }}</td>
                <td>{{ registro.Entrega }}</td>
                <td>{{ registro.Assistencia }}</td>
                <td>{{ registro.LenteContato }}</td>
                <td>{{ registro.Fluxo }}</td>
                <td>{{ registro.Venda }}</td>
                <td>{{ registro.Anotacoes }}</td>
              </tr>
            </template>
          </template>
        </v-data-table>
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
    const periodoFiltro = ref("semanal");
    const anoFiltro = ref(null);

    const periodos = ["semanal", "mensal", "anual"];

    const anosDisponiveis = computed(() => {
      const anos = registros.value.map((registro) => new Date(registro.Data).getFullYear());
      return [...new Set(anos)].sort((a, b) => b - a);
    });

    const headers = [
      // { title: "Data", key: "Data" },
      { title: "Vendedor", key: "VendedorNome" },
      { title: "Orçamento c/ Receita", key: "OrcamentoReceita" },
      { title: "Orçamento s/ Receita", key: "OrcamentoSemReceita" },
      { title: "Óculos Solar", key: "OculosSolar" },
      { title: "Ajuste", key: "Ajuste" },
      { title: "Entrega", key: "Entrega" },
      { title: "Assistência", key: "Assistencia" },
      { title: "L/C - Ñ compra", key: "LenteContato" },
      { title: "Fluxo", key: "Fluxo" },
      { title: "Venda", key: "Venda" },
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

    const filtrarRegistros = computed(() => {
      if (!registros.value.length) return [];

      const hoje = new Date();
      const inicioPeriodo = new Date();

      if (periodoFiltro.value === "semanal") {
        inicioPeriodo.setDate(hoje.getDate() - 7);
      } else if (periodoFiltro.value === "mensal") {
        inicioPeriodo.setMonth(hoje.getMonth() - 1);
      } else if (periodoFiltro.value === "anual") {
        inicioPeriodo.setFullYear(hoje.getFullYear() - 1);
      }

      return registros.value.filter((registro) => {
        const dataRegistro = new Date(registro.Data);
        const anoRegistro = dataRegistro.getFullYear();

        return (
          dataRegistro >= inicioPeriodo &&
          (!anoFiltro.value || anoRegistro === anoFiltro.value)
        );
      });
    });

    const registrosAgrupados = computed(() => {
      const agrupados = {};

      filtrarRegistros.value.forEach((registro) => {
        const dataFormatada = formatData(registro.Data);
        if (!agrupados[dataFormatada]) {
          agrupados[dataFormatada] = [];
        }
        agrupados[dataFormatada].push(registro);
      });

      return agrupados;
    });

    onMounted(carregarRegistros);

    return {
      registros,
      headers,
      formatData,
      periodoFiltro,
      periodos,
      anoFiltro,
      anosDisponiveis,
      filtrarRegistros,
      registrosAgrupados,
    };
  },
};
</script>
