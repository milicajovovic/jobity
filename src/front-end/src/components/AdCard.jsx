import axios from "axios";
import React, { useEffect, useState } from "react";
import { Card, NavLink } from "react-bootstrap";
import EmployerInfo from "../components/EmployerInfo";

function AdCard({ ad }) {
    const [employer, setEmployer] = useState([]);
    const [showEmployer, setShowEmployer] = useState(false);

    useEffect(() => {
        axios.get("http://localhost:3003/employer/" + ad.EmployerID).then(res => {
            setEmployer(res.data);
        });
        // eslint-disable-next-line
    }, []);

    return (
        <Card style={{ width: "40rem" }} className="mb-4">
            <Card.Header className="d-flex flex-row">
                <div>
                    <Card.Title>{ad.Name}</Card.Title>
                    <Card.Subtitle>{ad.Posted}</Card.Subtitle>
                </div>
                <div className="my-auto ms-auto">
                    <NavLink onClick={() => setShowEmployer(true)}>{employer.Name}</NavLink>
                </div>
            </Card.Header>
            <Card.Body>
                <Card.Text>{ad.Description}</Card.Text>
            </Card.Body>
            <EmployerInfo show={showEmployer} onHide={() => setShowEmployer(false)} employer={employer} />
        </Card>
    )
}

export default AdCard