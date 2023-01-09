import React from "react";
import { Link } from "react-router-dom";

const Navigation = () => (
  <div>
      <h2>Hadoop Config Help Web</h2>
    <nav>
      <ul>
        <li>
          <Link to="/setconf">Setconf</Link>
        </li>
        <li>
            <Link to="/download">Download</Link>
        </li>
      </ul>
    </nav>
  </div>
);

export default Navigation;
