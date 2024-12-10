<template>
  <Navbar />
  <div class="page-layout">
    <div class="sidebar">
      <div class="side-form">
        <img src="../../../public/logo.png" alt="Plant Shop Logo" class="logo"/>
        <h2 style="color: #7E7E7E; font-family: 'Century Gothic', sans-serif">Продавец</h2>
        <div class="user-info">
          <img :src="photo" alt="User Photo" class="user-photo">
          <div class="user-details">
            <p>{{ lastName }}</p>
            <p>{{ firstName }}</p>
            <p>{{ patronymic }}</p>
          </div>
        </div>
      </div>
    </div>

    <div class="plant-container">
      <div class="search-plants">
        <input class="search-input" type="text" placeholder="Поиск по объявлениям"/>
        <button class="green-button-white-text" id="search-button">Найти</button>
      </div>

      <div class="plant-info-grid">
        <div>
          <h2 style="color: #89A758; font-family: 'Century Gothic', sans-serif">{{ accepterPlant.species }}</h2>
          <img :src="accepterPlant.image" alt="Plant Photo" class="specific-plant-photo">
        </div>
        <div style="margin-left: 2%">
          <h2 style="color: #000000; font-family: 'Century Gothic', sans-serif; margin-bottom: 0">{{ accepterPlant.price }}</h2>
          <p style="color: #7E7E7E; font-size: 14px; margin-bottom: 0; margin-top: 0">{{ accepterPlant.place }}</p>
          <p style="color: #000000; font-size: 16px; font-weight: bold; margin-bottom: 0">Характеристики</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Размер: {{accepterPlant.size}}</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Условия освещения: {{accepterPlant.lighting}}</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Частота полива: {{accepterPlant.wateringFrequency}}</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Температурный режим: {{accepterPlant.temperature}}</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Сложность ухода: {{accepterPlant.careLevel}}</p>
          <p style="color: #000000; font-size: 16px; font-weight: bold; margin-bottom: 0">Описание</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">{{ accepterPlant.description }}</p>
          <p style="color: #000000; font-size: 16px; font-weight: bold; margin-bottom: 0">Тип</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">{{ accepterPlant.type }}</p>
          <div>
            <button class="white-button-green-text-sale">Купить</button>
            <button class="white-button-green-text-sale">Обменяться</button>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div v-if="isOpen" class="modal-overlay" @click="close">
    <div class="trade-modal-content" @click.stop>
      <header class="trade-modal-header">
        <div style="display: flex; justify-content: space-between; width: 100%;">
          <h2 style="font-family: 'Century Gothic', sans-serif; color: black">Обмен</h2>
          <button @click="closeModal" class="close-button">X</button>
        </div>
        <div style="margin-top: 2%;">
            <span :style="{ fontWeight: 'bold', color: 'black', fontSize: '14px'}">
              Выберете одно из своих объявлений для совершения обмена
            </span>
        </div>
      </header>

      <section class="modal-body" v-for="(ads, index) in userAds" :key="index">
        <div style="display: flex;flex-direction: row">
          <div style="display: flex;flex-direction: row">
            <input v-model="trade_plant_id" type="radio" value="Холодостойкие (до 15°C)" />
            <img v-if="ads.image" :src="ads.image" alt="Plant Image" class="trade-plant-image" />
          </div>
          <div style="margin-left: 4%">
            <p :style="{fontWeight: 'bold', color: '#89A758'}">{{ ads.species }}</p>
            <p :style="{fontWeight: 'bold'}">{{ ads.price }}</p>
            <p>{{ ads.place }}</p>
            <p>{{ ads.data }}</p>
          </div>
        </div>
        <hr style="margin-top: 10px; border: 1px solid #ccc;" />
      </section>
    </div>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import axios from "axios";

export default {
  name: "Plant",
  components: { Navbar },

  data() {
    return {
      lastName: '',
      firstName: '',
      patronymic: '',
      photo: '',
      accepterPlant: {
        id: '',
        type: '',
        species: '',
        size: '',
        lighting: '',
        wateringFrequency: '',
        temperature: '',
        careLevel: '',
        description: '',
        place: '',
        image: '',
        price: null
      },
      accepterID: '',
      offererID: '',
      isOpen: false,
      userAds: [],
      trade_plant_id: '',
      plants_for_trade: []
    };
  },

  mounted() {
    this.accepterPlant.id = sessionStorage.getItem("plant");
    this.offererID = sessionStorage.getItem("id");
    this.getPlant();
  },

  methods: {
    errorBuy() {
      this.$notify({
        title: "Ошибка!",
        text: "Произошла ошибка при покупке, попробуйте еще раз.",
        type: 'error'
      });
    },

    successBuy() {
      this.$notify({
        title: "Получилось!",
        text: "Вы приобрели новое растение.",
        type: 'success'
      });
    },

    async getPlant() {
      axios
          .get(`/api/plants/${this.accepterPlant.id}`)
          .then((response) => {
            this.accepterPlant.species = response.data.plant.species;
            this.accepterPlant.type = response.data.plant.type;
            this.accepterPlant.size = response.data.plant.size;
            this.accepterPlant.lighting = response.data.plant.lightCondition;
            this.accepterPlant.wateringFrequency = response.data.plant.wateringFrequency;
            this.accepterPlant.temperature = response.data.plant.temperatureRegime;
            this.accepterPlant.careLevel = response.data.plant.careComplexity;
            this.accepterPlant.description = response.data.plant.description;
            this.accepterPlant.place = response.data.plant.place;
            this.accepterPlant.image = response.data.plant.image;
            this.accepterPlant.price = response.data.plant.price;
            this.firstName = response.data.user.name;
            this.lastName = response.data.user.surname;
            this.patronymic = response.data.user.fatherName;
            this.photo = response.data.user.photo;
          })
    },

    async getPlantsForTrade() {
      this.plants_for_trade = [];
      axios
          .get(`/api/plants/${this.plantID}`)
          .then((response) => {
            response.data.plants.forEach(elem => {
              let plant = {
                id: elem.id,
                image: elem.image,
                species: elem.species,
                price: elem.price,
                createdAt: elem.createdAt,
                place: elem.place
              };
              this.plants.push(plant);
            });
            this.plantsCount = parseInt(response.data.count);
          })
    },

    async buyPlant() {
      const buyData ={
        sellerId: this.accepterID,
        buyerId: this.offererID,
        plantId: this.accepterPlant.id,
        species: this.accepterPlant.species,
        price: this.accepterPlant.id
      }

      try {
        await axios.post(`/api/plants/buy`, buyData);
        this.successBuy();
        this.$router.push('/plants/sale');
      } catch (error) {
        this.errorBuy();
      }
    },

    formatDate(date) {
      return new Date(date).toLocaleDateString("ru-RU", {
        year: 'numeric',
        month: 'long',
        day: 'numeric'
      });
    },

    formatPrice(price) {
      return `${price} ₽`;
    },

    openModal() {
      this.isOpen = true;
      this.getIn();
      this.getOut();
    },

    closeModal() {
      this.isOpen = false;
    },
  }
}
</script>

<style>
@import "../../../main.css";
@import "../../../plants.css";
@import "../../../user.css";

.trade-modal-content {
  background: #fff;
  width: 60%;
  max-width: 500px;
  padding: 20px;
  border-radius: 8px;
  position: relative;
}

.trade-modal-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  border-bottom: 1px solid #ccc;
  padding-bottom: 10px;
  margin-bottom: 10px;
  flex-direction: column;
}

.trade-modal-header h2 {
  font-size: 1.5em;
  margin: 0;
}
</style>