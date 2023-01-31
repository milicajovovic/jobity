import axios from "axios";
import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Col, Container, Row, Card, Button, Table, NavLink } from "react-bootstrap";
import EmployerMenu from "../components/EmployerMenu";
import EmployeeInfo from "../components/EmployeeInfo";
import AdInfo from "../components/AdInfo";
import DeclineApplication from "../components/DeclineApplication";
import SetInterview from "../components/SetInterview";

function EmployerInterviews() {
    const [applications, setApplications] = useState([]);
    const [showEmployee, setShowEmployee] = useState(false);
    const [employee, setEmployee] = useState({});
    const [showAd, setShowAd] = useState(false);
    const [ad, setAd] = useState({});
    const [showInterview, setShowInterview] = useState(false);
    const [showDecision, setShowDecision] = useState(false);
    const [application, setApplication] = useState({});

    useEffect(() => {
        let jwt = localStorage.getItem("jwt");
        let userId = localStorage.getItem("userId");
        axios.get("http://localhost:3007/applications/interviews/" + userId, { headers: { Authorization: jwt } }).then(res => {
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

    const changeInterview = (application) => {
        setApplication(application);
        setShowInterview(true);
    };

    const decide = (application) => {
        setApplication(application);
        setShowDecision(true);
    };

    const convertDate = (interview) => {
        let date = new Date(interview);
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
        <Container fluid>
            <Row>
                <EmployerMenu />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "50rem" }}>
                        <Card.Title className="text-center mt-3" as="h3">Interviews</Card.Title>
                        <Card.Body className="text-center">
                            <Table striped bordered>
                                <thead>
                                    <tr>
                                        <th>Ad</th>
                                        <th>Employee</th>
                                        <th>Interview</th>
                                        <th>Decision</th>
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
                                                <NavLink onClick={() => changeInterview(application)}>{convertDate(application.Interview)}</NavLink>
                                            </td>
                                            <td>
                                                <Button onClick={() => decide(application)}>
                                                    Decide
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
            <SetInterview show={showInterview} onHide={() => setShowInterview(false)} application={application} close={setShowInterview} />
            <DeclineApplication show={showDecision} onHide={() => setShowDecision(false)} application={application} close={setShowDecision} />
        </Container>
    )
}

export default EmployerInterviews