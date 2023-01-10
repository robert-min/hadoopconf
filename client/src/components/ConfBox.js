import React from "react";

const Confbox = ({conf}) => {

  return (
    <ul>
      {conf.map((c) => (
        <li key={c.name}>
          {c.name}
          <p>{c.description}</p>
        </li>
      ))}
    </ul>
  );
};


export default Confbox