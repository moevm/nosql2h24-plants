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
          <h2 style="color: #89A758; font-family: 'Century Gothic', sans-serif">{{ species }}</h2>
          <img :src="plantImage" alt="Plant Photo" class="specific-plant-photo">
        </div>
        <div>
          <h2 style="color: #000000; font-family: 'Century Gothic', sans-serif; margin-bottom: 0">{{ price }}</h2>
          <p style="color: #7E7E7E; font-size: 14px; margin-bottom: 0; margin-top: 0">{{ place }}</p>
          <p style="color: #000000; font-size: 16px; font-weight: bold; margin-bottom: 0">Характеристики</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Размер: {{size}}</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Условия освещения: {{lighting}}</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Частота полива: {{wateringFrequency}}</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Температурный режим: {{temperature}}</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">Сложность ухода: {{careLevel}}</p>
          <p style="color: #000000; font-size: 16px; font-weight: bold; margin-bottom: 0">Описание</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">{{ description }}</p>
          <p style="color: #000000; font-size: 16px; font-weight: bold; margin-bottom: 0">Тип</p>
          <p style="color: #000000; font-size: 16px; margin-bottom: 0; margin-top: 0">{{ type }}</p>
          <div>
            <button class="white-button-green-text">Купить</button>
            <button class="white-button-green-text">Обменяться</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import axios from "axios";

export default {
  name: "SpecificPlant",
  components: { Navbar },

  data() {
    return {
      lastName: '',
      firstName: '',
      patronymic: '',
      photo: '',
      plantID: '',
      type: '',
      species: '',
      size: '',
      lighting: '',
      wateringFrequency: '',
      temperature: '',
      careLevel: '',
      description: '',
      place: '',
      plantImage: '',
      price: null,
      userId: ''
    };
  },

  mounted() {
    this.plantID = sessionStorage.getItem("specificPlant");
    this.getPlant();
  },

  methods: {
    async getPlant() {
      axios
          .get(`/api/plants/${this.plantID}`)
          .then((response) => {
            this.species = response.data.plant.species;
            this.type = response.data.plant.type;
            this.size = response.data.plant.size;
            this.lighting = response.data.plant.lightCondition;
            this.wateringFrequency = response.data.plant.wateringFrequency;
            this.temperature = response.data.plant.temperatureRegime;
            this.careLevel = response.data.plant.careComplexity;
            this.description = response.data.plant.description;
            this.place = response.data.plant.place;
            this.plantImage = response.data.plant.image;
            this.price = response.data.plant.price;
            this.firstName = response.data.user.name;
            this.lastName = response.data.user.surname;
            this.patronymic = response.data.user.fatherName;
            this.photo = response.data.user.photo;
          })
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
    }
  }
}
</script>

<style>
@import "../../../main.css";
@import "../../../plants.css";
@import "../../../user.css";
</style>