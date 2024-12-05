<template>
  <Navbar />
  <div class="page-layout">
    <div class="sidebar">
      <div class="side-form">
        <img src="../../../public/logo.png" alt="Plant Shop Logo" class="logo"/>
        <form @submit.prevent="submitForm">
          <div class="inputs-labels">
            Тип растения
            <input class="inputs" v-model="formData.type" placeholder="Наименование типа" />
          </div>

          <div class="inputs-labels">Размер</div>
          <label class="checkbox-labels"><input v-model="formData.size" type="checkbox" value="Маленькие (до 20 см)" /> Маленькие (до 20 см)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.size" type="checkbox" value="Средние (от 20 до 50 см)" /> Средние (от 20 до 50 см)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.size" type="checkbox" value="Большие (более 50 см)" /> Большие (более 50 см)</label>

          <div class="inputs-labels">Условия освещения</div>
          <label class="checkbox-labels"><input v-model="formData.lighting" type="checkbox" value="Тенелюбивые" /> Тенелюбивые</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.lighting" type="checkbox" value="Полутеневые" /> Полутеневые</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.lighting" type="checkbox" value="Светолюбивые" /> Светолюбивые</label>

          <div class="inputs-labels">Частота полива</div>
          <label class="checkbox-labels"><input v-model="formData.wateringFrequency" type="checkbox" value="Редкий полив (раз в 2 недели)" /> Редкий полив (раз в 2 недели)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.wateringFrequency" type="checkbox" value="Средний полив (1-2 раза в неделю)" /> Средний полив (1-2 раза в неделю)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.wateringFrequency" type="checkbox" value="Частый полив (ежедневно)" /> Частый полив (ежедневно)</label>

          <div class="inputs-labels">Температурный режим</div>
          <label class="checkbox-labels"><input v-model="formData.temperature" type="checkbox" value="Холодостойкие (до 15°C)" /> Холодостойкие (до 15°C)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.temperature" type="checkbox" value="Средний режим (15-22°C)" /> Средний режим (15-22°C)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.temperature" type="checkbox" value="Теплолюбивые (более 22°C)" /> Теплолюбивые (более 22°C)</label>

          <div class="inputs-labels">Сложность ухода</div>
          <label class="checkbox-labels"><input v-model="formData.careLevel" type="checkbox" value="Для начинающих" /> Для начинающих</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.careLevel" type="checkbox" value="Требует среднего ухода" /> Требует среднего ухода</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.careLevel" type="checkbox" value="Для опытных цветоводов" /> Для опытных цветоводов</label>
          <br>

          <div class="inputs-labels">
            Слова в описании
            <textarea class="inputs" v-model="formData.description" placeholder="Что-то важное для вас"></textarea>
          </div>

          <div class="inputs-labels">
            Вид растения
            <input class="inputs" v-model="formData.species" type="text" placeholder="Наименование вида" />
          </div>

          <div class="inputs-labels">
            Город
            <input class="inputs" v-model="formData.city" type="text" placeholder="Введите город" />
          </div>

          <div class="inputs-labels">
            Цена, Р
            <div style="display: flex; justify-content: space-between">
              <input class="inputs" style="margin-right: 2%" v-model="formData.priceFrom" type="number" placeholder="От" />
              <input class="inputs" v-model="formData.priceTo" type="number" placeholder="До" />
            </div>
          </div>

          <button type="submit" style="margin-top: 2%" class="green-button-white-text">Показать объявления</button>
        </form>
      </div>
    </div>

    <div class="plant-container">
      <div class="search-plants">
        <input class="search-input" v-model="search" type="text" placeholder="Поиск по объявлениям" />
        <button class="green-button-white-text" id="search-button">Найти</button>
      </div>

      <div class="select-sort">
        <select class="custom-select" v-model="sort_type">
          <option disabled value="">Сортировка</option>
          <option>По умолчанию</option>
          <option>Дешевле</option>
          <option>Дороже</option>
          <option>По дате</option>
        </select>
      </div>

      <div class="plant-grid">
        <div v-for="plant in plants" class="plant-card">
          <div class="plant-content">
            <img v-if="plant.image" :src="plant.image" alt="Plant Image" class="plant-image" />
            <div class="plant-info">
              <div v-if="plant.species" class="plant-title">{{ plant.species }}</div>
              <div v-if="plant.price" class="plant-price">{{ formatPrice(plant.price) }}</div>
              <div v-if="plant.place" class="plant-place">{{ plant.place }}</div>
              <div v-if="plant.createdAt" class="plant-date">{{ formatDate(plant.createdAt) }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import axios from "axios";

const PLANTS_URL = '/api/plants';
const NEW_PLANT_URL = '/api/plants/add';

export default {
  name: "Sale",
  components: { Navbar },

  data() {
    return {
      plants: [],
      formData: {
        type: '',
        species: '',
        size: [],
        lighting: [],
        wateringFrequency: [],
        temperature: [],
        careLevel: [],
        description: '',
        city: '',
        priceFrom: null,
        priceTo: null,
        image: '',
      },
      search: '',
      sort_type: '',
      userId: ''
    };
  },

  mounted() {
    this.getPlants();
    this.userId = sessionStorage.getItem("id");
  },

  methods: {
    async getPlants() {
      axios
          .get(PLANTS_URL)
          .then((response) => {
            response.data.plants.forEach(elem => {
              let plant = {
                image: elem.image,
                species: elem.species,
                price: elem.price,
                createdAt: elem.createdAt,
                place: elem.place
              };
              this.plants.push(plant)
            })
          })
    },

    addImage(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.formData.image = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },

    submitForm() {
      console.log(this.formData.image)
      this.createPlant();
    },

    async createPlant() {
      const plantData = {
        image: 'https://i.pinimg.com/736x/c2/ad/d9/c2add9a552ba76ebe2c1c42e487766f7.jpg',
        place: this.formData.city,
        size: this.formData.size,
        price: this.formData.price,
        lightCondition: this.formData.lighting,
        wateringFrequency: this.formData.wateringFrequency,
        temperatureRegime: this.formData.temperature,
        careComplexity: this.formData.careLevel,
        description: this.formData.description,
        type: this.formData.type,
        species: this.formData.species,
        createdAt: new Date(),
        userId: this.userId
      };

      try {
        await axios.post(NEW_PLANT_URL, plantData);
        alert('Объявление успешно добавлено!');
        this.clearForm();
        location.reload();
      } catch (error) {
        alert('Произошла ошибка при добавлении объявления. Попробуйте снова.');
      }
    },

    clearForm() {
      this.formData.city = '';
      this.formData.size = '';
      this.formData.price = '';
      this.formData.lighting = '';
      this.formData.wateringFrequency = '';
      this.formData.temperature = '';
      this.formData.careLevel = '';
      this.formData.description = '';
      this.formData.type = '';
      this.formData.species = '';
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

<style scoped>
@import "../../../main.css";
@import "../../../plants.css";

#search-button {
  margin-left: 1%;
  width: 8%;
}

.custom-select {
  font-family: 'Century Gothic', sans-serif;
  font-size: 13px;
  color: #000000;
  width: 200px;
  padding: 10px;
  border-color: transparent;
  background-color: #FFFFFF;
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
  cursor: pointer;
}

.custom-select:focus {
  border-color: transparent;
  outline: none;
}

.custom-select::after {
  position: absolute;
  border: 2px solid transparent;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  pointer-events: none;
}

.custom-select option {
  font-weight: 400;
  color: #000000;
  padding: 10px;
  background-color: #FFFFFF;
}

.plant-price {
  font-family: 'Century Gothic', sans-serif;
  color: black;
  font-weight: bold;
}

.plant-place,
.plant-date {
  font-size: 13px;
  font-family: 'Century Gothic', sans-serif;
  color: #7E7E7E;
}
</style>