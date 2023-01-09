import React, { useEffect, useState } from "react";
import axios from "axios";

const Setconf = () => {
  const [hdfs, setHdfs] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchHdfs = async () => {
    try {
      setError(null);
      setHdfs(null);
      setLoading(true);

      const response = await axios.get("http://localhost:8080/api/hdfs");
      setHdfs(response.data);
    } catch (error) {
      setError(error);
    }
    setLoading(false)
  };
  useEffect(() => {fetchHdfs();}, []);

  if (loading) return <div>Loading ...</div>;
  if (error) return <div>Error!!! Return Page</div>;
  if (!hdfs) return null;

  return (
    <div>
      <ul>
        {hdfs.map(h => (
            <li key={h.id}>
                {h.name} : {h.description}
            </li>
        ))}
      </ul>
      
    </div>
  );
};

export default Setconf;
