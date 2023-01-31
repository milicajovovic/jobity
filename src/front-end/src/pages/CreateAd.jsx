import React, { useState, useEffect } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { ToastContainer, Toast, Form, Button, Container, Row, Col, Card } from "react-bootstrap";
import EmployerMenu from "../components/EmployerMenu";

function CreateAd() {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    const [types, setTypes] = useState([]);
    const [skills, setSkills] = useState([]);
    const [checkedTypes] = useState([]);
    const [checkedSkills] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        axios.get("http://localhost:3007/ads/jobTypes").then(res => {
            setTypes(res.data);
        });

        axios.get("http://localhost:3007/ads/requiredSkills").then(res => {
            setSkills(res.data);
        });
    }, []);

    const create = (event) => {
        event.preventDefault();

        let jwt = localStorage.getItem("jwt");
        let userId = localStorage.getItem("userId");

        const newAd = {
            "Name": event.target.name.value,
            "EmployerID": parseInt(userId),
            "Description": event.target.description.value,
            "Posted": (new Date()).toISOString(),
            "JobType": checkedTypes,
            "RequiredSkills": checkedSkills
        };

        axios.post("http://localhost:3007/ads/create", newAd, { headers: { Authorization: jwt } }).then(res => {
            navigate("/employer/home");
        }).catch((err) => {
            setMessage(err.response.data);
            setShow(true);
        });
    }

    const typesChanged = (event) => {
        if (event.target.checked) {
            checkedTypes.push(event.target.id);
        } else {
            checkedTypes.splice(checkedTypes.indexOf(event.target.id));
        }
    };

    const skillsChanged = (event) => {
        if (event.target.checked) {
            checkedSkills.push(event.target.id);
        } else {
            checkedSkills.splice(checkedTypes.indexOf(event.target.id));
        }
    };

    return (
        <Container fluid>
            <ToastContainer position="top-center" className="text-center p-3">
                <Toast onClose={() => setShow(false)} show={show} delay={3000} autohide>
                    <Toast.Body>{message}</Toast.Body>
                </Toast>
            </ToastContainer>
            <Row>
                <EmployerMenu />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "40rem" }}>
                        <Card.Title className="text-center mt-3 mb-3" as="h3">Create new ad</Card.Title>
                        <Form onSubmit={create}>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Name</Form.Label>
                                    <Form.Control type="text" name="name" placeholder="Enter name" required />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Description</Form.Label>
                                    <Form.Control as="textarea" rows={2} name="description" placeholder="Enter description" required />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Job types</Form.Label>
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
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Required skills</Form.Label>
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
                                </Col>
                            </Row>
                            <Row>
                                <Col className="d-grid">
                                    <Button variant="primary" type="submit" className="mb-2">
                                        Create
                                    </Button>
                                </Col>
                            </Row>
                        </Form>
                    </Card>
                </Col>
            </Row>
        </Container>
    )
}

export default CreateAd