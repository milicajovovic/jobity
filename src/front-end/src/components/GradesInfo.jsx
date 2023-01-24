import axios from "axios";
import React, { useEffect, useState } from "react";
import { FaStar } from "react-icons/fa";

function GradesInfo({ employerID }) {
    const [grades, setGrades] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:3004/employer/" + employerID).then(res => {
            setGrades(res.data);
        });
        // eslint-disable-next-line
    }, []);

    return (
        <div>
            {grades.map(grade => (
                <div key={grade.ID}>
                    <hr />
                    {[...Array(grade.Grade)].map((v, i) => <FaStar key={i} />)}
                    <p>{grade.Comment}</p>
                </div>
            ))}
        </div>
    )
}

export default GradesInfo