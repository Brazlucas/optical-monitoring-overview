<template>
  <v-row class="pa-5">
    <v-col cols="12">
      <v-card>
        <v-card-title>Controle da Loja</v-card-title>
        <v-card-text>
          <v-form>
            <v-row>
              <!-- Seleção de Vendedor -->
              <v-col cols="12" md="6">
                <v-select
                  v-model="vendedorSelecionado"
                  :items="vendedores"
                  item-title="Nome"
                  item-value="ID"
                  label="Vendedor"
                  outlined
                ></v-select>
              </v-col>
              <!-- Data -->
              <v-col cols="12" md="6">
                <v-text-field v-model="data" label="Data" type="date" outlined></v-text-field>
              </v-col>
              <!-- Orçamentos -->
              <v-col cols="12" md="6">
                <v-text-field v-model="orcamentoReceita" label="Orçamento c/ Receita" type="number" outlined></v-text-field>
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field v-model="orcamentoSemReceita" label="Orçamento s/ Receita" type="number" outlined></v-text-field>
              </v-col>
              <!-- Óculos Solar -->
              <v-col cols="12" md="6">
                <v-text-field v-model="oculosSolar" label="Óculos Solar" type="number" outlined></v-text-field>
              </v-col>
              <!-- Ajuste -->
              <v-col cols="12" md="6">
                <v-text-field v-model="ajuste" label="Ajuste" type="number" outlined></v-text-field>
              </v-col>
              <!-- Entrega -->
              <v-col cols="12" md="6">
                <v-text-field v-model="entrega" label="Entrega" type="number" outlined></v-text-field>
              </v-col>
              <!-- Assistência -->
              <v-col cols="12" md="6">
                <v-text-field v-model="assistencia" label="Assistência" type="number" outlined></v-text-field>
              </v-col>
              <!-- Anotações -->
              <v-col cols="12">
                <v-textarea v-model="anotacoes" label="Anotações" rows="2" outlined></v-textarea>
              </v-col>
            </v-row>
            <!-- Botões -->
            <v-row>
              <v-col cols="12" md="6">
                <v-btn color="primary" block @click="adicionarRegistro">Adicionar</v-btn>
              </v-col>
              <v-col cols="12" md="6">
                <v-btn color="grey" block outlined @click="$router.go(-1)">Cancelar</v-btn>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { AdicionarRegistro, ListarVendedores } from "../../wailsjs/go/main/App";

export default {
  setup() {
    const router = useRouter();
    const data = ref("");
    const orcamentoReceita = ref(0);
    const orcamentoSemReceita = ref(0);
    const oculosSolar = ref(0);
    const ajuste = ref(0);
    const entrega = ref(0);
    const assistencia = ref(0);
    const anotacoes = ref("");
    const vendedores = ref([]);
    const vendedorSelecionado = ref(null);

    const carregarVendedores = async () => {
      try {
        vendedores.value = await ListarVendedores();
      } catch (error) {
        console.error("Erro ao carregar vendedores:", error);
      }
    };

    onMounted(carregarVendedores);

    const adicionarRegistro = async () => {

    if (!vendedorSelecionado.value) {
      alert("Selecione um vendedor antes de adicionar o registro!");
      return;
    }

    try {
      // Converter a data para o formato que o backend espera
      const dataFormatada = new Date(data.value).toISOString().split("T")[0]; // YYYY-MM-DD
      
      // Converter vendedor para número, garantindo que não seja string
      const vendedorID = parseInt(vendedorSelecionado.value);

      await AdicionarRegistro(
        dataFormatada, // Data convertida
        parseInt(orcamentoReceita.value) || 0,
        parseInt(orcamentoSemReceita.value) || 0,
        parseInt(oculosSolar.value) || 0,
        parseInt(ajuste.value) || 0,
        parseInt(entrega.value) || 0,
        parseInt(assistencia.value) || 0,
        vendedorID, // ID convertido para número
        anotacoes.value
      );
        router.push("/");
      } catch (error) {
        console.error("Erro ao adicionar registro:", error);
      }
    };

    return {
      data,
      orcamentoReceita,
      orcamentoSemReceita,
      oculosSolar,
      ajuste,
      entrega,
      assistencia,
      anotacoes,
      vendedores,
      vendedorSelecionado,
      adicionarRegistro,
    };
  },
};
</script>
