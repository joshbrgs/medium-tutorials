import { createApp } from "vue";

import Counter from "./counter.vue";

export default (el) => {
   createApp(Counter).mount(el);
}
