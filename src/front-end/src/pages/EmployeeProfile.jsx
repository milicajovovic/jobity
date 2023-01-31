import axios from "axios";
import React, { useEffect, useState } from "react";
import { Col, Container, Row, Card, ToastContainer, Toast, Form, Button } from "react-bootstrap";
import EmployeeMenu from "../components/EmployeeMenu";

function EmployeeProfile() {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    const [employee, setEmployee] = useState({});
    const [types, setTypes] = useState([]);
    const [typesSame, setTypesSame] = useState(true);
    const [skills, setSkills] = useState([]);
    const [skillsSame, setSkillsSame] = useState(true);
    const [checkedTypes] = useState([]);
    const [checkedSkills] = useState([]);
    const [noPdf, setNoPdf] = useState(false);

    useEffect(() => {
        const jwt = localStorage.getItem("jwt");
        const userId = localStorage.getItem("userId");
        axios.get("http://localhost:3007/employees/employee/" + userId, { headers: { Authorization: jwt } }).then(res => {
            setEmployee(res.data);
        });

        axios.get("http://localhost:3007/ads/jobTypes").then(res => {
            setTypes(res.data);
        });

        axios.get("http://localhost:3007/ads/requiredSkills").then(res => {
            setSkills(res.data);
        });
    }, []);

    const update = (event) => {
        event.preventDefault();

        let password = event.target.password.value;
        if (password === "Password") {
            password = employee.Password;
        }

        const updatedEmployee = {
            "ID": employee.ID,
            "Email": employee.Email,
            "Password": password,
            "FirstName": event.target.firstName.value,
            "LastName": event.target.lastName.value,
            "Birthday": employee.Birthday,
            "Education": event.target.education.value,
            "JobType": typesSame ? employee.JobType : checkedTypes,
            "Skills": skillsSame ? employee.Skills : checkedSkills,
            "CV": event.target.cv.value,
        };

        let jwt = localStorage.getItem("jwt");

        if (updatedEmployee.CV !== "") {
            let dto = {
                "PdfPath": updatedEmployee.CV.split("\\").pop(),
                "Email": updatedEmployee.Email,
                "Password": updatedEmployee.Password,
                "EmployeeID": employee.ID
            };

            axios.post("http://localhost:3007/employees/update/pdf", dto, { headers: { Authorization: jwt } }).then(res => {
                setMessage("successfully updated");
                setShow(true);
            }).catch((err) => {
                setMessage(err.response.data);
                setShow(true);
            });
        } else {
            setNoPdf(true);

            axios.post("http://localhost:3007/employees/update/form", updatedEmployee, { headers: { Authorization: jwt } }).then(res => {
                setMessage("successfully updated");
                setShow(true);
            }).catch((err) => {
                setMessage(err.response.data);
                setShow(true);
            });
        }
    };

    const typesChanged = (event) => {
        if (event.target.checked) {
            checkedTypes.push(event.target.id);
        } else {
            checkedTypes.splice(checkedTypes.indexOf(event.target.id));
        }
        setTypesSame(false);
    };

    const employesType = (type) => {
        if (employee.JobType !== undefined) {
            return employee.JobType.includes(type);
        }
        return false;
    };

    const skillsChanged = (event) => {
        if (event.target.checked) {
            checkedSkills.push(event.target.id);
        } else {
            checkedSkills.splice(checkedTypes.indexOf(event.target.id));
        }
        setSkillsSame(false);
    };

    const employesSkill = (skill) => {
        if (employee.Skills !== undefined) {
            return employee.Skills.includes(skill);
        }
        return false;
    };

    return (
        <Container fluid>
            <Row>
                <EmployeeMenu />
            </Row>
            <ToastContainer position="top-center" className="text-center p-3">
                <Toast onClose={() => setShow(false)} show={show} delay={3000} autohide>
                    <Toast.Body>{message}</Toast.Body>
                </Toast>
            </ToastContainer>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "40rem" }}>
                        <Card.Title className="text-center mt-3 mb-3" as="h3">My profile</Card.Title>
                        <Form onSubmit={update}>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Email</Form.Label>
                                    <Form.Control type="text" name="email" defaultValue={employee.Email} disabled />
                                </Col>
                                <Col>
                                    <Form.Label>Password</Form.Label>
                                    <Form.Control type="password" name="password" defaultValue="Password" required />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>First name</Form.Label>
                                    <Form.Control type="text" name="firstName" defaultValue={employee.FirstName} required={noPdf} />
                                </Col>
                                <Col>
                                    <Form.Label>Last name</Form.Label>
                                    <Form.Control type="text" name="lastName" defaultValue={employee.LastName} required={noPdf} />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>CV in PDF</Form.Label>
                                    <Form.Control type="file" name="cv" defaultValue={employee.CV} />
                                </Col>
                            </Row>
                            <Row className="mb-3">
                                <Col>
                                    <Form.Label>Education</Form.Label>
                                    <Form.Control as="textarea" rows={2} name="education" defaultValue={employee.Education} />
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
                                                defaultChecked={employesType(type)}
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
                                                defaultChecked={employesSkill(skill)}
                                            />
                                        ))}
                                    </div>
                                </Col>
                            </Row>
                            <Row>
                                <Col className="d-grid">
                                    <Button variant="primary" type="submit" className="mb-2">
                                        Save changes
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

export default EmployeeProfile