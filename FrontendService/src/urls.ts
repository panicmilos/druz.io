const env = (window as any)._env_;

export const API_URL = env.SAME_HOST === 'true' ? 
  `${env.API_PROTOCOL}://${window.location.hostname}:${env.API_PORT}${env.API_PATH}` 
  :
  `${env.API_PROTOCOL}://${env.API_HOSTNAME}:${env.API_PORT}${env.API_PATH}`

export const USERS_SERVICE_URL = `http://localhost:8000`;
export const USER_RELATIONS_SERVICE_URL = `http://localhost:8001`;
export const POSTS_SERVICE_URL = `http://localhost:8004`;


