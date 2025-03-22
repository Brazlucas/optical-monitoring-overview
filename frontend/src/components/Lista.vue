<template>
  <v-row class="pa-5">
    <v-col cols="12">
      <v-card elevation="3">
        <v-row class="ml-1 justify-end">

          <!-- <v-col cols="2">
            <v-select
              v-model="anoFiltro"
              :items="anosDisponiveis"
              label="Filtrar por Ano"
              outlined
            ></v-select>
          </v-col> -->

          <v-col cols="2">
            <v-select
              color="red"
              v-model="periodoFiltro"
              :items="periodos"
              label="Filtrar por intervalo"
            ></v-select>
          </v-col>

          <v-col cols="2">
            <v-btn @click="$router.push('/form')" color="red" class="mt-4 mb-5">
              <v-icon left>mdi-plus</v-icon> Novo Registro
            </v-btn>
          </v-col>

        </v-row>
        <v-card-title class="text-h5">Registros</v-card-title>
        <v-divider></v-divider>
      
        <v-data-table
          :items="registrosFiltrados"
          :headers="headers"
          item-value="id"
          density="comfortable"
        >
          <template v-slot:item.Data="{ item }">
            {{ formatData(item.Data) }}
          </template>
          <!-- Linha de Totais no corpo da tabela -->
          <template v-slot:body.append>
            <tr>
              <td><strong>Total</strong></td>
              <td></td>
              <td><b>{{ totalOrcamentoReceita }}</b></td>
              <td><b>{{ totalOrcamentoSemReceita }}</b></td>
              <td><b>{{ totalOculosSolar }}</b></td>
              <td><b>{{ totalAjuste }}</b></td>
              <td><b>{{ totalEntrega }}</b></td>
              <td><b>{{ totalAssistencia }}</b></td>
              <td></td>
            </tr>
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

    // Gerando os anos disponíveis a partir dos registros
    const anosDisponiveis = computed(() => {
      const anos = registros.value.map((registro) => new Date(registro.Data).getFullYear());
      return [...new Set(anos)].sort((a, b) => b - a); // Ordenando de forma decrescente
    });

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

    const filtrarRegistros = (filtro, ano) => {
      const hoje = new Date();
      const inicioPeriodo = new Date();

      if (filtro === "semanal") {
        inicioPeriodo.setDate(hoje.getDate() - 7); // 7 dias atrás
      } else if (filtro === "mensal") {
        inicioPeriodo.setMonth(hoje.getMonth() - 1); // 1 mês atrás
      } else if (filtro === "anual") {
        inicioPeriodo.setFullYear(hoje.getFullYear() - 1); // 1 ano atrás
      }

      return registros.value.filter((registro) => {
        const dataRegistro = new Date(registro.Data);
        const anoRegistro = dataRegistro.getFullYear();

        // Filtro por intervalo e por ano
        return (
          dataRegistro >= inicioPeriodo &&
          (!ano || anoRegistro === ano)
        );
      });
    };

    // Computed properties para calcular totais
    const totalOrcamentoReceita = computed(() =>
      registrosFiltrados.value.reduce((sum, r) => sum + r.OrcamentoReceita, 0)
    );

    const totalOrcamentoSemReceita = computed(() =>
      registrosFiltrados.value.reduce((sum, r) => sum + r.OrcamentoSemReceita, 0)
    );

    const totalOculosSolar = computed(() =>
      registrosFiltrados.value.reduce((sum, r) => sum + r.OculosSolar, 0)
    );

    const totalAjuste = computed(() =>
      registrosFiltrados.value.reduce((sum, r) => sum + r.Ajuste, 0)
    );

    const totalEntrega = computed(() =>
      registrosFiltrados.value.reduce((sum, r) => sum + r.Entrega, 0)
    );

    const totalAssistencia = computed(() =>
      registrosFiltrados.value.reduce((sum, r) => sum + r.Assistencia, 0)
    );

    const registrosFiltrados = computed(() => filtrarRegistros(periodoFiltro.value, anoFiltro.value));

    onMounted(carregarRegistros);

    return {
      registros,
      headers,
      formatData,
      periodoFiltro,
      periodos,
      anoFiltro,
      anosDisponiveis,
      registrosFiltrados,
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
