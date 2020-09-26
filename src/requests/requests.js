import axios from 'axios'

export default {
  catchErrors (error) {
    if (error.response) {
      return { data: null, err: error.response.data }
    } else if (error.request) {
      return { data: null, err: error.request }
    } else {
      return { data: null, err: error.message }
    }
  },
  submitForm (form) {
    return axios.post('/api/form', form, {}).then(resp => {
      return { data: resp.data, err: null }
    }).catch(error => this.catchErrors(error))
  }
}
