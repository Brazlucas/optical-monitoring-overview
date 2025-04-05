<template>
  <v-row class="pa-5">
    <v-col cols="12">
      <v-card elevation="8" class="rounded-lg">
        <v-card-title class="text-h5 font-weight-bold">Registros</v-card-title>
        <v-divider></v-divider>

        <v-row class="ml-1 justify-end" no-gutters>
          <v-select
            color="red"
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

          <v-btn @click="$router.push('/manage-form')" color="red" class="mt-4 mb-5 mr-2">
            <v-icon left>mdi-cog</v-icon> Gerenciar
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
                <td>
                  <a v-bind="props" class="anchor-modal" @click="abrirAnotacoes(registro.Anotacoes)">
                    {{ registro.Anotacoes }}
                  </a>
                </td>
              </tr>
            </template>
          </template>
          
          <!-- Totais com Tooltip e Modal -->
          <template v-slot:body.append>
            <tr class="font-weight-bold bg-yellow-lighten-3">
              <td><strong>Total</strong></td>

              <td v-for="(total, campo) in totais" :key="total">
                <span
                  @click="abrirModal(campo, total)"
                >
                  <b>{{ total }}</b>
                </span>
              </td>
              <td></td>
            </tr>
          </template>
        </v-data-table>
      </v-card>
    </v-col>

    <!-- Modal de Detalhes -->
    <v-dialog v-model="modalAberto" max-width="500px">
      <v-card>
        <v-card-title>
          <span class="text-h5">{{ tituloModal }}</span>
        </v-card-title>
        <v-card-text>
          <p>{{ totalModal }}</p>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="red" text @click="modalAberto = false">Fechar</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import { ref, computed, onMounted, reactive } from "vue";
import { ListarRegistros } from "../../wailsjs/go/main/App";

export default {
  setup() {
    const registros = ref([]);
    const periodoFiltro = ref("semanal");
    const anoDefault = new Date().getFullYear();
    const anoFiltro = ref(anoDefault);

    const modalAberto = ref(false);
    const tituloModal = ref("");
    const totalModal = ref(0);

    const periodos = ["semanal", "mensal", "anual"];

    const anosDisponiveis = computed(() => {
      const anos = registros.value.map((registro) => new Date(registro.Data).getFullYear());
      return [...new Set(anos)].sort((a, b) => b - a);
    });

    const headers = [
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
      const date = new Date(`${data}T00:00:00`);
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

    const calcularTotal = (campo) =>
      computed(() =>
        Object.values(registrosAgrupados.value)
          .flat()
          .reduce((sum, r) => sum + (r[campo] || ''), 0)
      );

    const totais = reactive({
      totalOrcamentoReceita: calcularTotal("OrcamentoReceita"),
      totalOrcamentoSemReceita: calcularTotal("OrcamentoSemReceita"),
      totalOculosSolar: calcularTotal("OculosSolar"),
      totalAjuste: calcularTotal("Ajuste"),
      totalEntrega: calcularTotal("Entrega"),
      totalAssistencia: calcularTotal("Assistencia"),
      totalLenteContato: calcularTotal("LenteContato"),
      totalFluxo: calcularTotal("Fluxo"),
      totalVenda: calcularTotal("Venda"),
    });

    const abrirModal = (campo, total) => {
      totalModal.value = total;
      const campoToLowerCase = campo.toLowerCase();
      const headerToLowerCase = headers.map((h) => `total${h.key.toLowerCase()}`);

      const toLowerCase = headerToLowerCase.find((h) => h === campoToLowerCase);

      if (toLowerCase) {
        tituloModal.value = headers[headerToLowerCase.indexOf(toLowerCase)].title;
        modalAberto.value = true;
      }
    };

    const abrirAnotacoes = (anotacoes) => {
      tituloModal.value = "Anotações";
      modalAberto.value = true;
      totalModal.value = anotacoes;
    };

    onMounted(carregarRegistros);

    return {
      registros,
      headers,
      periodoFiltro,
      periodos,
      anoFiltro,
      anosDisponiveis,
      registrosAgrupados,
      totais,
      modalAberto,
      tituloModal,
      totalModal,
      abrirAnotacoes,
      abrirModal,
    };
  },
};
</script>

<style scoped>
a {
  cursor: pointer !important;
  color: red !important;
  transition: color 0.3s ease !important;
}

a:hover {
  color: gray !important;
  text-decoration: underline !important;
}
</style>
