<template>
  <div class="navbar">

    <div class="menu-icons">
      <i v-if="role === 0" class="icon fa fa-chart-bar" @click="navigate(0)"></i>
      <i class="icon fa fa-search" @click="navigate(role === 0 ? 1 : 0)"></i>
      <i class="icon fa fa-leaf" @click="navigate(role === 0 ? 2 : 1)"></i>
      <i class="icon fa fa-bell" @click="openModal" ></i>
      <i class="icon fa fa-user-circle" @click="navigate(role === 0 ? 3 : 2)"></i>
    </div>

    <button class="my-ads-button" @click="navigate(role === 0 ? 4 : 3)">Мои объявления</button>

    <div v-if="isOpen" class="modal-overlay" @click="close">
      <div class="modal-content" @click.stop>
        <header class="modal-header">
          <div style="display: flex; justify-content: space-between; width: 100%;">
            <h2 color="black">Запросы</h2>
            <button @click="closeModal" class="close-button">X</button>
          </div>
          <div style="margin-top: 2%;">
            <span :style="{ fontWeight: selected === 'in' ? 'bold' : 'normal', color: 'black', fontSize: '14px'}" @click="select('in')" class="clickable">
              Входящие
            </span>
            <span :style="{ fontWeight: selected === 'out' ? 'bold' : 'normal', color: 'black', fontSize: '14px' }" @click="select('out')" class="clickable">
              Исходящие
            </span>
          </div>
        </header>
        <section v-if="selected === 'in'" class="modal-body" v-for="(trades, index) in inT" :key="index">
          <p :style="{fontWeight: 'bold'}">{{ trades.offerer }} предлагает</p>
          <p>{{ trades.plant_in }}</p>
          <p :style="{fontWeight: 'bold'}">В обмен на</p>
          <p>{{ trades.plant_out }}</p>
          <div style="margin-top: 2%; display: flex; gap: 10px;">
            <button class="trade-button">Согласиться</button>
            <button class="trade-button">Отказаться</button>
          </div>
        </section>
        <section v-if="selected === 'out'" class="modal-body" v-for="(trades, index) in outT" :key="index">
          <p :style="{fontWeight: 'bold'}">Вы предложили</p>
          <p>{{ trades.plant_out }}</p>
          <p :style="{fontWeight: 'bold'}">В обмен на</p>
          <p>{{ trades.plant_in }}</p>
          <p :style="{fontWeight: 'bold', color: '#89A758'}">Статус</p>
          <p>{{ trades.status }}</p>
          <hr style="margin-top: 10px; border: 1px solid #ccc;" />
        </section>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Navbar",

  data() {
    return {
      role: '',
      id: '',
      currentType: 'in',
      inT: [],
      outT: [],
      selected: 'in',
      isOpen: false,
      adminItems: [
        { path: '/plants/statistics' },
        { path: '/plants/sale' },
        { path: '/plants/care'},
        { path: '/plants/user' },
        { path: '/plants/ads' }
      ],
      regularItems: [
        { path: '/plants/sale' },
        { path: '/plants/care'},
        { path: '/plants/user' },
        { path: '/plants/ads' }
      ]
    };
  },

  beforeMount() {
    this.role = sessionStorage.getItem("role");
    this.id = sessionStorage.getItem("id");
    this.getIn();
    this.getOut();
  },

  computed: {
    menuItems() {
      return this.role === 0 ? this.adminItems : this.regularItems;
    }
  },

  methods: {
    navigate(index) {
      const target = this.menuItems[index];
      if (target) {
        this.$router.push(target.path);
      }
    },

    openModal() {
      this.isOpen = true;
    },

    closeModal() {
      this.isOpen = false;
    },

    select(type) {
      this.selected = type;
    },

    async getIn() {
      axios
          .get(`/api/trade/in/${this.id}`)
          .then((response) => {
            response.data.trade.forEach(elem => {
              let in_trade = {
                offerer: elem.offerer.name,
                plant_in: elem.offerer.plant.name,
                plant_out: elem.accepter.plant.name,
              };
              this.inT.push(in_trade)
            })
          })
    },

    async getOut() {
      axios
          .get(`/api/trade/out/${this.id}`)
          .then((response) => {
            response.data.trade.forEach(elem => {
              let trade_status = '';
              switch (elem.status) {
                case 'STATUS_UNSPECIFIED':
                  trade_status = 'Неизвестный статус';
                  break;
                case 'STATUS_PENDING':
                  trade_status = 'Ожидает подтверждения';
                  break;
                case 'STATUS_ACCEPTED':
                  trade_status = 'Согласовано';
                  break;
                case 'STATUS_REJECTED':
                  trade_status = 'Отклонено';
                  break;
              }
              let out_trade = {
                plant_out: elem.offerer.plant.name,
                plant_in: elem.accepter.plant.name,
                status: trade_status
              };
              this.outT.push(out_trade)
            })
          })
    },
  }
}
</script>

<style scoped>
@import "../../main.css";
@import "../../modal.css";

.navbar {
  display: flex;
  align-items: center;
  padding: 10px 20px;
  background-color: #333;
  color: #fff;
  position: fixed;
  top: 0;
  left: 0;
  width: 98%;
  z-index: 1000;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
}

.menu-icons {
  display: flex;
  gap: 15px;
  align-items: center;
  margin-left: auto;
}

.icon {
  font-size: 18px;
  cursor: pointer;
  color: #fff;
}

.my-ads-button {
  padding: 8px 16px;
  background-color: #89A758;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
  margin-left: 20px;
}

.my-ads-button:hover {
  background-color: #77934a;
}

.clickable {
  cursor: pointer;
  margin-right: 10px;
}

.modal-body p {
  margin: 1px 0;
  line-height: 1.2;
  font-size: 14px;
  color: black;
}

.trade-button {
  font-size: 12px;
  font-weight: bold;
  color: #2E2E2E;
  background-color: transparent;
  border: 2px solid #89A758;
  border-radius: 10px;
  padding: 10px 15px;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  position: relative;
}
</style>
