export const Token = '7d45d757a974e302bc54';
export const LocalStorageName = 'token';

// not practical but makes this demo easy to run
export function login() {
  window.localStorage.setItem(LocalStorageName, Token);
  window.location.href = "/";
}

export function logout() {
  window.localStorage.removeItem(LocalStorageName);
  window.location.href = "/login";
}

export function getToken() {
  return window.localStorage.getItem(LocalStorageName);
}
