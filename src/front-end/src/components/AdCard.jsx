import axios from "axios";
import React, { useEffect, useState } from "react";
import { Button, Card, Col, NavLink, Row } from "react-bootstrap";
import EmployerInfo from "../components/EmployerInfo";
import AdApplication from "./AdApplication";
import Confirmation from "./Confirmation";
import EditAd from "./EditAd";

function AdCard({ ad }) {
    const [employer, setEmployer] = useState({});
    const [showEmployer, setShowEmployer] = useState(false);
    const [employeeLogged, setEmployeeLogged] = useState(false);
    const [employerLogged, setEmployerLogged] = useState(false);
    const [showApplication, setShowApplication] = useState(false);
    const [question, setQuestion] = useState("");
    const [showConfirmation, setShowConfirmation] = useState(false);
    const [showEditAd, setShowEditAd] = useState(false);

    useEffect(() => {
        axios.get("http://localhost:3007/employers/employer/" + ad.EmployerID).then(res => {
            setEmployer(res.data);
        });

        if (localStorage.getItem("role") === "employee") {
            setEmployeeLogged(true);
        } else if (localStorage.getItem("role") === "employer") {
            setEmployerLogged(true);
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

    const deleteAd = () => {
        setQuestion("Are you sure you want to delete " + ad.Name + "?");
        setShowConfirmation(true);
    };

    return (
        <Card style={{ width: "40rem" }} className="mb-4">
            <Card.Header className="d-flex">
                <div>
                    <Card.Title>{ad.Name}</Card.Title>
                    <Card.Subtitle>{convertDate(ad.Posted)}</Card.Subtitle>
                </div>
                <div className="my-auto ms-auto">
                    {employerLogged ?
                        null
                        : <NavLink onClick={() => setShowEmployer(true)}>{employer.Name}</NavLink>
                    }
                </div>
            </Card.Header>
            <Card.Body>
                <Row>
                    <Col md="auto">
                        <Card.Text>{ad.Description}</Card.Text>
                    </Col>
                    <Col className="d-flex justify-content-end">
                        {employeeLogged ?
                            <Button onClick={() => setShowApplication(true)}>Apply</Button>
                            : null
                        }
                        {employerLogged ?
                            <div>
                                <Button className="me-2" onClick={() => setShowEditAd(true)}>Edit</Button>
                                <Button onClick={deleteAd}>Delete</Button>
                            </div>
                            : null
                        }
                    </Col>
                </Row>
            </Card.Body>
            <EmployerInfo show={showEmployer} onHide={() => setShowEmployer(false)} employer={employer} />
            <AdApplication show={showApplication} onHide={() => setShowApplication(false)} ad={ad} close={setShowApplication} />
            <Confirmation show={showConfirmation} onHide={() => setShowConfirmation(false)} question={question} id={ad.ID} action={"ads/delete"} close={setShowConfirmation} />
            <EditAd show={showEditAd} onHide={() => setShowEditAd(false)} ad={ad} close={setShowEditAd} />
        </Card>
    )
}

export default AdCard