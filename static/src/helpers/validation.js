export const maxChar = (str, max) => str.length > max && true;
export const isEmpty = val => val.trim().length === 0;
export const isValidEmail = val =>
  new RegExp("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$").test(val);
