import CounterExample from 'components/counter-example'
import FetchData from 'components/fetch-data'
import HomePage from 'components/home-page'
import FamilyPage from 'components/family/family.page'
import PeoplePage from 'components/people/people.page'

export const routes = [
  { name: 'home', path: '/', component: HomePage, display: 'Dashboard', icon: 'tachometer-alt' },
  { name: 'family', path: '/family', component: FamilyPage, display: 'Family', icon: 'home' },
  { name: 'people', path: '/people', component: PeoplePage, display: 'People', icon: 'users' },
  { name: 'counter', path: '/counter', component: CounterExample, display: 'Counter', icon: 'graduation-cap' },
  { name: 'fetch-data', path: '/fetch-data', component: FetchData, display: 'Fetch data', icon: 'list' }
]
