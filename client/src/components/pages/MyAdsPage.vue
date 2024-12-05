<template>
  <Navbar />
  <div class="page-layout">
    <div class="sidebar">
      <div class="side-form">
        <img src="../../../public/logo.png" alt="Plant Shop Logo" class="logo"/>
        <div class="inputs-labels" id="create-ads">
          Создать объявление
        </div>
        <form @submit.prevent="submitForm">
          <div class="inputs-labels">
            Тип растения
            <input class="inputs" v-model="formData.type" placeholder="Наименование типа" required/>
          </div>

          <div class="inputs-labels">
            Вид растения
            <input class="inputs" v-model="formData.species" type="text" placeholder="Наименование вида" required/>
          </div>

          <div class="inputs-labels">Размер</div>
          <label class="checkbox-labels"><input v-model="formData.size" type="radio" value="Маленькие (до 20 см)" /> Маленькие (до 20 см)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.size" type="radio" value="Средние (от 20 до 50 см)" /> Средние (от 20 до 50 см)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.size" type="radio" value="Большие (более 50 см)" /> Большие (более 50 см)</label>

          <div class="inputs-labels">Условия освещения</div>
          <label class="checkbox-labels"><input v-model="formData.lighting" type="radio" value="Тенелюбивые" /> Тенелюбивые</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.lighting" type="radio" value="Полутеневые" /> Полутеневые</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.lighting" type="radio" value="Светолюбивые" /> Светолюбивые</label>

          <div class="inputs-labels">Частота полива</div>
          <label class="checkbox-labels"><input v-model="formData.wateringFrequency" type="radio" value="Редкий полив (раз в 2 недели)" /> Редкий полив (раз в 2 недели)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.wateringFrequency" type="radio" value="Средний полив (1-2 раза в неделю)" /> Средний полив (1-2 раза в неделю)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.wateringFrequency" type="radio" value="Частый полив (ежедневно)" /> Частый полив (ежедневно)</label>

          <div class="inputs-labels">Температурный режим</div>
          <label class="checkbox-labels"><input v-model="formData.temperature" type="radio" value="Холодостойкие (до 15°C)" /> Холодостойкие (до 15°C)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.temperature" type="radio" value="Средний режим (15-22°C)" /> Средний режим (15-22°C)</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.temperature" type="radio" value="Теплолюбивые (более 22°C)" /> Теплолюбивые (более 22°C)</label>

          <div class="inputs-labels">Сложность ухода</div>
          <label class="checkbox-labels"><input v-model="formData.careLevel" type="radio" value="Для начинающих" /> Для начинающих</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.careLevel" type="radio" value="Требует среднего ухода" /> Требует среднего ухода</label>
          <br>
          <label class="checkbox-labels"><input v-model="formData.careLevel" type="radio" value="Для опытных цветоводов" /> Для опытных цветоводов</label>
          <br>

          <div class="inputs-labels">
            Описание
            <textarea class="inputs" v-model="formData.description" placeholder="Введите описание"></textarea>
          </div>

          <div class="inputs-labels">
            Город
            <input class="inputs" v-model="formData.city" type="text" placeholder="Введите город" required/>
          </div>

          <div class="inputs-labels">
            Цена, Р
            <input class="inputs" v-model="formData.price" type="number" placeholder="Введите цену" required/>
          </div>

          <button class="white-button-green-text" style="margin-bottom: 2%">Добавить изображения</button>
          <button type="submit" class="green-button-white-text">Опубликовать</button>
        </form>
      </div>
    </div>

    <div class="plant-container">
      <header class="modal-header">
        <h1 style="font-family: Century Gothic,sans-serif">Мои объявления</h1>
        <span
            :style="{ color: selected === 'active' ? '#000000' : '#7E7E7E', fontWeight: 'bold', fontSize: '14px', fontFamily: 'Century Gothic,sans-serif'}"
            @click="select('active')"
            class="clickable">
          Входящие
        </span>
        <span
            :style="{ color: selected === 'archive' ? '#000000' : '#7E7E7E', fontWeight: 'bold', fontSize: '14px', fontFamily: 'Century Gothic,sans-serif' }"
            @click="select('archive')"
            class="clickable">
          Исходящие
        </span>
      </header>
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
      selected: 'active',
      formData: {
        type: '',
        species: '',
        size: '',
        lighting: '',
        wateringFrequency: '',
        temperature: '',
        careLevel: '',
        description: '',
        city: '',
        price: null,
        image: '',
      },
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

    select(type) {
      this.selected = type;
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

.page-layout {
  display: flex;
}

#search-button {
  margin-left: 1%;
  width: 8%;
}

#create-ads {
  color: #89A758;
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