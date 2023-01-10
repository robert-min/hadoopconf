import React, { useEffect, useState } from "react";
import axios from "axios";
import Confbox from "../components/ConfBox";
import styles from "./Setconf.module.css";
import configform from "./configform.json"

const Setconf = () => {
  const [config, setConfig] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [text, setText] = useState("")


  const onClickHDFS = (event) => {
    event.preventDefault();
    fetchHdfs("http://localhost:8080/api/hdfs")
  }
  const onClickCORE = (event) => {
    event.preventDefault();
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

  const textChange = (event) => {
    setText(event.target.value);
  }



  return (
    <>
    <div className={styles.configbox}>
      <h2>SetConf</h2>
      <button onClick={onClickHDFS}>HDFS config</button>
      <button onClick={onClickCORE}>CORE config</button>
      <Confbox conf={config}/>
    </div>
    <div className={styles.typingbox}>
      <form className={styles.typingform}>
      <pre>{configform.front}</pre>
        <textarea className={styles.typingarea} type="text" value={text} onChange={e => textChange(e)} placeholder="write here" />
        <pre>{configform.back}</pre>
      </form>

    </div>
    </>
  );
};

export default Setconf;
