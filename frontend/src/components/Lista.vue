<template>
  <div class="container">
    <h1>Registros de Vendas</h1>
    <table>
      <thead>
        <tr>
          <th>Data</th>
          <th>Orçamento c/ Receita</th>
          <th>Orçamento s/ Receita</th>
          <th>Óculos Solar</th>
          <th>Ajuste</th>
          <th>Entrega</th>
          <th>Assistência</th>
          <th>Anotações</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="registro in registros" :key="registro.id">
          <td>{{ registro.Data }}</td>
          <td>{{ registro.OrcamentoReceita }}</td>
          <td>{{ registro.OrcamentoSemReceita }}</td>
          <td>{{ registro.OculosSolar }}</td>
          <td>{{ registro.Ajuste }}</td>
          <td>{{ registro.Entrega }}</td>
          <td>{{ registro.Assistencia }}</td>
          <td>{{ registro.Anotacoes }}</td>
        </tr>
      </tbody>
    </table>
    <button @click="$router.push('/')">Adicionar Novo</button>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import { ListarRegistros } from "../../wailsjs/go/main/App";

export default {
  setup() {
    const registros = ref([]);

    const carregarRegistros = async () => {
      registros.value = await ListarRegistros();
    };

    onMounted(carregarRegistros);

    return { registros };
  },
};
</script>

<style scoped>
/* Adicione estilos conforme necessário */
</style>
