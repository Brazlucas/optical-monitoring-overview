<template>
  <div class="container">
    <h1>Controle de Vendas</h1>
    <div class="form-container">
      <div class="form-group">
        <label>Data:</label>
        <input type="date" v-model="data" />
      </div>
      <div class="form-group">
        <label>Orçamento c/ Receita:</label>
        <input type="number" v-model="orcamentoReceita" />
      </div>
      <div class="form-group">
        <label>Orçamento s/ Receita:</label>
        <input type="number" v-model="orcamentoSemReceita" />
      </div>
      <div class="form-group">
        <label>Óculos Solar:</label>
        <input type="number" v-model="oculosSolar" />
      </div>
      <div class="form-group">
        <label>Ajuste:</label>
        <input type="number" v-model="ajuste" />
      </div>
      <div class="form-group">
        <label>Entrega:</label>
        <input type="number" v-model="entrega" />
      </div>
      <div class="form-group">
        <label>Assistência:</label>
        <input type="number" v-model="assistencia" />
      </div>
      <div class="form-group">
        <label>Anotações:</label>
        <textarea rows="2" v-model="anotacoes"></textarea>
      </div>
    </div>
    <button @click="adicionarRegistro">Adicionar</button>
  </div>
</template>

<script>
import { ref } from "vue";
import { useRouter } from "vue-router"; // Import correto do useRouter
import { AdicionarRegistro } from "../../wailsjs/go/main/App";

export default {
  setup() {
    const router = useRouter(); // Inicializando o router
    const data = ref("");
    const orcamentoReceita = ref(0);
    const orcamentoSemReceita = ref(0);
    const oculosSolar = ref(0);
    const ajuste = ref(0);
    const entrega = ref(0);
    const assistencia = ref(0);
    const anotacoes = ref("");

    const adicionarRegistro = async () => {
      try {
        await AdicionarRegistro(
          data.value,
          orcamentoReceita.value,
          orcamentoSemReceita.value,
          oculosSolar.value,
          ajuste.value,
          entrega.value,
          assistencia.value,
          1, // ID do vendedor (poderia ser selecionado dinamicamente)
          anotacoes.value
        );
        // Redireciona para a página de registros
        router.push("/registros");
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
      adicionarRegistro,
    };
  },
};
</script>
