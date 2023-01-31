import axios from "axios";
import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Col, Container, Row, Card, Button, Table } from "react-bootstrap";
import EmployerMenu from "../components/EmployerMenu";
import EmployeeInfo from "../components/EmployeeInfo";
import AdInfo from "../components/AdInfo";
import DeclineApplication from "../components/DeclineApplication";
import SetInterview from "../components/SetInterview";

function EmployerApplications() {
    const [applications, setApplications] = useState([]);
    const [showEmployee, setShowEmployee] = useState(false);
    const [employee, setEmployee] = useState({});
    const [showAd, setShowAd] = useState(false);
    const [ad, setAd] = useState({});
    const [showAccept, setShowAccept] = useState(false);
    const [showDecline, setShowDecline] = useState(false);    
    const [application, setApplication] = useState({});

    useEffect(() => {
        let jwt = localStorage.getItem("jwt");
        let userId = localStorage.getItem("userId");
        axios.get("http://localhost:3007/applications/employer/" + userId, { headers: { Authorization: jwt } }).then(res => {
            setApplications(res.data);
        });
    }, [applications]);

    const viewAd = (id) => {
        let jwt = localStorage.getItem("jwt");
        axios.get("http://localhost:3007/ads/ad/" + id, { headers: { Authorization: jwt } }).then(res => {
            setAd(res.data);
        });
        setShowAd(true);
    };

    const viewEmployee = (id) => {
        let jwt = localStorage.getItem("jwt");
        axios.get("http://localhost:3007/employees/employee/" + id, { headers: { Authorization: jwt } }).then(res => {
            setEmployee(res.data);
        });
        setShowEmployee(true);
    };

    const accept = (application) => {
        setApplication(application);
        setShowAccept(true);
    };

    const decline = (application) => {
        setApplication(application);
        setShowDecline(true);
    };

    return (
        <Container fluid>
            <Row>
                <EmployerMenu />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "50rem" }}>
                        <Card.Title className="text-center mt-3" as="h3">Applications</Card.Title>
                        <Card.Body className="text-center">
                            <Table striped bordered>
                                <thead>
                                    <tr>
                                        <th>Ad</th>
                                        <th>Employee</th>
                                        <th>Accept</th>
                                        <th>Decline</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {applications.map(application => (
                                        <tr key={application.ID}>
                                            <td>
                                                <Button onClick={() => viewAd(application.AdID)}>
                                                    View
                                                </Button>
                                            </td>
                                            <td>
                                                <Button onClick={() => viewEmployee(application.EmployeeID)}>
                                                    View
                                                </Button>
                                            </td>
                                            <td>
                                                <Button onClick={() => accept(application)}>
                                                    Accept
                                                </Button>
                                            </td>
                                            <td>
                                                <Button onClick={() => decline(application)}>
                                                    Decline
                                                </Button>
                                            </td>
                                        </tr>
                                    ))}
                                </tbody>
                            </Table>
                        </Card.Body>
                    </Card>
                </Col>
            </Row>
            <AdInfo show={showAd} onHide={() => setShowAd(false)} ad={ad} />
            <EmployeeInfo show={showEmployee} onHide={() => setShowEmployee(false)} employee={employee} />
            <SetInterview show={showAccept} onHide={() => setShowAccept(false)} application={application} close={setShowAccept} />
            <DeclineApplication show={showDecline} onHide={() => setShowDecline(false)} application={application} close={setShowDecline} />
        </Container>
    )
}

export default EmployerApplications