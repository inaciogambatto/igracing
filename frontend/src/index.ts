import {name} from './create';
import axios from 'axios';

console.log(name);

const getMessage = async () => {
  try {
    const response = await axios.get('http://localhost:8080/api/helloworld');
    console.log(response.data);
  } catch (error) {
    console.error('Erro ao fazer a requisição:', error);
  }
};

getMessage();