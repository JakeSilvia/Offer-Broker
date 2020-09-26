<template>
  <div>
    <div class="section">
      <div class="container">
        <h1 class="title is-1 has-text-centered has-text-primary">
            OfferBroker
        </h1>
        <div class="container columns">
          <div class="column is-one-fifth"></div>
          <div class="column is-three-fifths">
            <b-field required class="pb-3" label="Email">
              <b-input placeholder="JohnD@gmail.com" rounded v-model="email" type="email"/>
            </b-field>
            <b-field class="pb-3" label="Search">
              <b-input expanded placeholder="iPhone X..." rounded v-model="search"/>
            </b-field>
            <b-field class="pb-3" label="Radius">
              <b-input expanded rounded type="number" v-model.number="radius"/>
            </b-field>
            <b-field groupd class="pb-3" label="Price">
              <b-input expanded rounded type="number" placeholder="Min" v-model.number="minPrice"/>
              <b-input expanded rounded type="number" placeholder="Max" v-model.number="maxPrice"/>
            </b-field>
            <div class="pt-4">
              <b-button type="is-primary" expanded :loading="isLoading" @click="submit">Submit</b-button>
            </div>
          </div>
          <div class="column is-one-fifth"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  // @ is an alias to /src

  export default {
    name: 'Home',
    data() {
      return {
        email: '',
        search: '',
        radius: null,
        maxPrice: null,
        minPrice: null,
        isLoading: false
      };
    },
    methods: {
      submit() {
        let err = this.validate()
        if (err != null) {
          this.$notification('is-danger', err)
          return
        }

        this.isLoading = true
        let form = {
          email: this.email,
          search: this.search,
          min_price: this.minPrice,
          max_price: this.maxPrice,
          radius: this.radius,
        }

        this.$ajax.submitForm(form).then(resp => {
          this.isLoading = false
          if (resp.err != null) {
            this.$notification('is-danger', resp.err)
            return
          }

          this.$notification('is-success', 'Thank you!')
        })
      },
      validate () {
        if (this.email === '') {
          return 'Please enter email.'
        }
        if (this.search === '') {
          return 'Please enter Search.'
        }

        return null
      }
    }
  };
</script>
