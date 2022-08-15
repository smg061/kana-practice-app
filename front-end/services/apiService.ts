import { Kana } from "../models/models"

export class ApiService { 

    constructor(private baseUrl: string) {}

    async getAll<T>(): Promise<T[]> {
        const response  = await fetch(this.baseUrl)
        const data: T[] = await response.json()
        return data
    }

    async getAllKana(): Promise<Kana[] > {
        const response  = await fetch(`${this.baseUrl}/kana`)
        const data: Kana[] = await response.json()
        return data
    }
}



