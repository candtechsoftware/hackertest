import HttpClient from './http-client';
import * as dotenv from 'dotenv';
dotenv.config();


export interface ProcessEnv {
    [key: string]: string | undefined;
}

const base: string = "https://jsonmock.hackerrank.com/api/countries/search"

class Country extends HttpClient {
    
  public constructor() {
    super(base);
  }

  getTotoalPages = async (s: string): Promise<number> =>{
    try {
      let pages = await this.instance.get(`?name=${s}`);
      return pages.data.total_pages; 
    } catch (err) {
      console.error(err); 
      return -1; 
    }
  }
    


  getCountry = async(s: string, p: number) => {
    let totalCountries: number = 0; 
    try {
    const totalPages = await this.getTotoalPages(s); 
    
  

    for (let i = 1; i <= totalPages; i++){
        let response = await this.instance.get(`?name=${s}&page=${i}`);
        let countries = response.data.data; 
        for (let country of countries ){
          if (country.population > p) totalCountries++;
        }

    }
    console.log("Number of countries that meet criteria: ", totalCountries);
    }catch(err) { 
      console.error(err);
    }  

  }
}

let c: Country = new Country();
c.getCountry("un", 100090);
