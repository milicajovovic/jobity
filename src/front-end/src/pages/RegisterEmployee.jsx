import React, { useState, useEffect } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { ToastContainer, Toast, Form, Button, Container, Row, Col, Card } from "react-bootstrap";
import NavigationBar from "../components/NavigationBar";

function RegisterEmployee() {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    const [types, setTypes] = useState([]);
    const [skills, setSkills] = useState([]);
    const [checkedTypes] = useState([]);
    const [checkedSkills] = useState([]);
    const [noPdf, setNoPdf] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        axios.get("http://localhost:3001/jobTypes").then(res => {
            setTypes(res.data);
        });

        axios.get("http://localhost:3001/requiredSkills").then(res => {
            setSkills(res.data);
        });
    }, []);

    const register = (event) => {
        event.preventDefault();

        const newEmployee = {
            "Email": event.target.email.value,
            "Password": event.target.password.value,
            "FirstName": event.target.firstName.value,
            "LastName": event.target.lastName.value,
            "Birthday": new Date(event.target.birthday.value).toISOString(),
            "Education": event.target.education.value,
            "JobType": checkedTypes,
            "Skills": checkedSkills,
            "CV": event.target.cv.value,
            "ProfilePicture": event.target.profilePicture.value,
        };

        console.log(newEmployee.Birthday)
        if (newEmployee.CV !== "") {
            let dto = {
                "PdfPath": newEmployee.CV.split("\\").pop(),
                "Password": newEmployee.Password,
            };

            axios.post("http://localhost:3002/register/pdf", dto).then(res => {
                navigate("/home/employee");
            }).catch((err) => {
                setMessage(err.response.data);
                setShow(true);
            });
        } else {
            setNoPdf(true);

            if (validEmail(newEmployee.Email)) {
                axios.post("http://localhost:3002/register/form", newEmployee).then(res => {
                    navigate("/home/employee");
                }).catch((err) => {
                    setMessage(err.response.data);
                    setShow(true);
                });
            } else {
                setMessage("email is not valid");
                setShow(true);
            }
        }        
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

    const validEmail = (email) => {
        let emailRegex = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/;
        if (email.match(emailRegex)) {
            return true;
        }
        return false;
    };

    return (
        <Container fluid>
            <ToastContainer position="top-center" className="text-center p-3">
                <Toast onClose={() => setShow(false)} show={show} delay={3000} autohide>
                    <Toast.Body>{message}</Toast.Body>
                </Toast>
            </ToastContainer>
            <Row>
                <NavigationBar />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "40rem" }}>
                        <Card.Title className="text-center mt-3 mb-3" as="h3">Register as employee</Card.Title>
                        <Form onSubmit={register}>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Email</Form.Label>
                                    <Form.Control type="text" name="email" placeholder="Enter email" required={noPdf} />
                                </Col>
                                <Col>
                                    <Form.Label>Password</Form.Label>
                                    <Form.Control type="password" name="password" placeholder="Enter password" required />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>First name</Form.Label>
                                    <Form.Control type="text" name="firstName" placeholder="Enter first name" required={noPdf} />
                                </Col>
                                <Col>
                                    <Form.Label>Last name</Form.Label>
                                    <Form.Control type="text" name="lastName" placeholder="Enter last name" required={noPdf} />
                                </Col>
                                <Col>
                                    <Form.Label>Birthday</Form.Label>
                                    <Form.Control type="date" name="birthday" required={noPdf} />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Profile picture</Form.Label>
                                    <Form.Control type="file" name="profilePicture" />
                                </Col>
                                <Col>
                                    <Form.Label>CV in PDF</Form.Label>
                                    <Form.Control type="file" name="cv" />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Education</Form.Label>
                                    <Form.Control as="textarea" rows={2} name="education" placeholder="Enter education" />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Desired job types</Form.Label>
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
                                    <Form.Label>Personal skills</Form.Label>
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
                                        Sign in
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

export default RegisterEmployee