import React, { useEffect, useState } from "react";
import axios from "axios";
import Confbox from "../components/ConfBox";

const Setconf = () => {
  const [config, setConfig] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const onClickHDFS = () => {
    fetchHdfs("http://localhost:8080/api/hdfs")
  }
  const onClickCORE = () => {
    fetchHdfs("http://localhost:8080/api/core")
  }

  const fetchHdfs = async (url) => {
    try {
      setError(null);
      setConfig(null);
      setLoading(true);

      const response = await axios.get(url);
      setConfig(response.data);
      
    } catch (error) {
      console.log(error)
      setError(error);
    }
    setLoading(false);
  };

  useEffect(() => {
    fetchHdfs("http://localhost:8080/api/hdfs");
  }, []);

  if (loading) return <div>Loading ...</div>;
  if (error) return <div>Error!!! Return Page</div>;
  if (!config) return null;


  return (
    <div>
      <h2>SetConf</h2>
      <button onClick={onClickHDFS}>HDFS config</button>
      <button onClick={onClickCORE}>CORE config</button>
      <Confbox conf={config}/>
    </div>
  );
};

export default Setconf;
