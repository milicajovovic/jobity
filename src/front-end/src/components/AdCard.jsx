import axios from "axios";
import React, { useEffect, useState } from "react";
import { Button, Card, NavLink } from "react-bootstrap";
import EmployerInfo from "../components/EmployerInfo";
import AdApplication from "./AdApplication";

function AdCard({ ad }) {
    const [employer, setEmployer] = useState({});
    const [showEmployer, setShowEmployer] = useState(false);
    const [employeeLogged, setEmployeeLogged] = useState(false);
    const [showApplication, setShowApplication] = useState(false);

    useEffect(() => {
        axios.get("http://localhost:3007/employers/employer/" + ad.EmployerID).then(res => {
            setEmployer(res.data);
        });

        if (localStorage.getItem("role") === "employee") {
            setEmployeeLogged(true);
        }
        // eslint-disable-next-line
    }, []);

    const convertDate = (posted) => {
        let date = new Date(posted);
        let year = date.getFullYear();
        let month = date.getMonth() + 1;
        let day = date.getDate();

        if (day < 10) {
            day = "0" + day;
        }
        if (month < 10) {
            month = "0" + month;
        }

        return day + '-' + month + '-' + year;
    };

    return (
        <Card style={{ width: "40rem" }} className="mb-4">
            <Card.Header className="d-flex">
                <div>
                    <Card.Title>{ad.Name}</Card.Title>
                    <Card.Subtitle>{convertDate(ad.Posted)}</Card.Subtitle>
                </div>
                <div className="my-auto ms-auto">
                    <NavLink onClick={() => setShowEmployer(true)}>{employer.Name}</NavLink>
                </div>
            </Card.Header>
            <Card.Body className="d-flex">
                <Card.Text>{ad.Description}</Card.Text>
                {employeeLogged ? <Button className="ms-auto" onClick={() => setShowApplication(true)}>Apply</Button> : null}
            </Card.Body>
            <EmployerInfo show={showEmployer} onHide={() => setShowEmployer(false)} employer={employer} />
            <AdApplication show={showApplication} onHide={() => setShowApplication(false)} ad={ad} close={setShowApplication} />
        </Card>
    )
}

export default AdCard