import { render, h } from "preact";

import "./index.scss";

const App = () => (
  <div class="mt-10 text-3xl mx-auto max-w-6xl">
    <div>Name: dashboard</div>
    <div>Framework: preact</div>
    <div>Language: TypeScript</div>
    <div>CSS: Tailwind</div>
  </div>
);

render(<App />, document.getElementById("app"));
