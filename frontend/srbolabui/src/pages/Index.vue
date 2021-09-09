<template>
  <q-page>
    <q-list>
      <q-table
        title="Templates"
        :data="templates"
        :columns="columns"
        row-key="enginenumber"
      />
    </q-list>
  </q-page>
</template>

<script>
  import axios from 'axios';

export default {
  name: 'PageIndex',

  data () {
    return {
      templates: [],
      columns: [
        {
          name: 'EngineNumber',
          required: true,
          label: 'Engine Number',
          field: 'EngineNumber',
          align: 'left',
          sortable: true
        },
        { name: 'ProductionYear', align: 'center', label: 'Production year', field: 'ProductionYear', sortable: true },
        { name: 'Manufacturer', label: 'Manufacturer', field: 'Manufacturer', sortable: true },
        { name: 'Color', label: 'Color', field: 'Color', sortable: true }
      ],
    }
  },
  mounted() {
    axios.get("http://127.0.0.1:8000/api/template/getall").then(result => {
      this.templates = JSON.parse(result.data.Data)
    }).catch(t => {
      alert('error')
    })
  }
}
</script>
