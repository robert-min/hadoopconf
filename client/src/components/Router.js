import { HashRouter as Router, Routes, Route } from "react-router-dom";
import Home from "../routes/Home";
import React from "react";
import Setconf from "../routes/Setconf";
import Download from "../routes/Download";
import Navigation from "./Navigation";

const AppRouter = () => {

    return (
        <Router>
            <Navigation />
            <Routes>
            <Route path="/" element={<Home />}></Route>
            <Route path="/setconf" element={<Setconf />}></Route>
            <Route path="/download" element={<Download />}></Route>
               
            </Routes>            
        </Router>
    );
}

export default AppRouter;