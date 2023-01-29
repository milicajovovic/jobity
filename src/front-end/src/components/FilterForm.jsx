import axios from "axios";
import React, { useState, useEffect } from "react";
import { Form } from "react-bootstrap";

function FilterForm({ ads, changeAds }) {
    const [types, setTypes] = useState([]);
    const [skills, setSkills] = useState([]);
    const [checkedTypes] = useState([]);
    const [checkedSkills] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:3007/ads/jobTypes").then(res => {
            setTypes(res.data);
        });

        axios.get("http://localhost:3007/ads/requiredSkills").then(res => {
            setSkills(res.data);
        });
    }, []);

    const typesChanged = (event) => {
        if (event.target.checked) {
            checkedTypes.push(event.target.id);
        } else {
            checkedTypes.splice(checkedTypes.indexOf(event.target.id));
        }

        filterAds();
    };

    const skillsChanged = (event) => {
        if (event.target.checked) {
            checkedSkills.push(event.target.id);
        } else {
            checkedSkills.splice(checkedSkills.indexOf(event.target.id));
        }

        filterAds();
    };

    const filterAds = () => {
        let filteredAds = [];

        filteredAds = ads.filter(ad => checkedTypes.every(type => ad.JobType.includes(type)));
        filteredAds = filteredAds.filter(ad => checkedSkills.every(skill => ad.RequiredSkills.includes(skill)));

        changeAds(filteredAds);
    }

    return (
        <Form>
            <div key={`types-checkbox`} className="mb-3">
                {types.map((type) => (
                    <Form.Check
                        key={type}
                        id={type}
                        label={type}
                        inline
                        type="checkbox"
                        onChange={typesChanged}
                    />
                ))}
            </div>
            <div key={`skills-checkbox`} className="mb-3">
                {skills.map((skill) => (
                    <Form.Check
                        key={skill}
                        id={skill}
                        label={skill}
                        inline
                        type="checkbox"
                        onChange={skillsChanged}
                    />
                ))}
            </div>
        </Form>
    )
}

export default FilterForm