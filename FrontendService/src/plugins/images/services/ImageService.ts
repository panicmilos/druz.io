import axios from "axios";
import { useState } from "react";
import { IMAGES_SERVICE_URL } from "../imports";
import { Image } from "../models/Image";

export const useImageService = () => {

  const [imageService] = useState(new ImageService());

  return imageService;
}

export class ImageService {

  public ID: string;
  private baseUrl: string;

  constructor() {
    this.ID = 'ImageService';
    this.baseUrl = `${IMAGES_SERVICE_URL}/images`;
  }

  public async upload(image: any): Promise<Image> {
    return (await axios.post(`${this.baseUrl}`, { ...image })).data;
  }
  
}