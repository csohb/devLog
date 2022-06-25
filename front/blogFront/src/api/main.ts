import reqresApi from './common';

export function fetchProfileList() {
  return reqresApi.get('/main/introduce');
}
