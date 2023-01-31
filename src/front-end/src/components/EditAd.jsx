import axios from "axios";
import React, { useState, useEffect } from "react";
import { ToastContainer, Toast, Form, Button, Container, Row, Col } from "react-bootstrap";
import Modal from 'react-bootstrap/Modal';

function EditAd(props) {
    const [show, setShow] = useState(false);
    const [message, setMessage] = useState("");
    const [types, setTypes] = useState([]);
    const [typesSame, setTypesSame] = useState(true);
    const [skills, setSkills] = useState([]);
    const [skillsSame, setSkillsSame] = useState(true);
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

    const edit = (event) => {
        event.preventDefault();

        const updatedAd = {
            "ID": props.ad.ID,
            "Name": event.target.name.value,
            "EmployerID": props.ad.EmployerID,
            "Description": event.target.description.value,
            "Posted": props.ad.Posted,
            "JobType": typesSame ? props.ad.JobType : checkedTypes,
            "RequiredSkills": skillsSame ? props.ad.RequiredSkills : checkedSkills
        };

        let jwt = localStorage.getItem("jwt");
        axios.post("http://localhost:3007/ads/update", updatedAd, { headers: { Authorization: jwt } }).then(res => {
            setMessage("successfully edited");
            setShow(true);
        }).catch((err) => {
            setMessage(err.response.data);
            setShow(true);
        });

        props.close(false);
    }

    const typesChanged = (event) => {
        if (event.target.checked) {
            checkedTypes.push(event.target.id);
        } else {
            checkedTypes.splice(checkedTypes.indexOf(event.target.id));
        }
        setTypesSame(false);
    };

    const adsType = (type) => {
        if (props.ad.JobType !== undefined) {
            return props.ad.JobType.includes(type);
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

    const adsSkill = (skill) => {
        if (props.ad.RequiredSkills !== undefined) {
            return props.ad.RequiredSkills.includes(skill);
        }
        return false;
    };

    return (
        <Container>
            <ToastContainer position="top-center" className="text-center p-3">
                <Toast onClose={() => setShow(false)} show={show} delay={3000} autohide>
                    <Toast.Body>{message}</Toast.Body>
                </Toast>
            </ToastContainer>
            <Modal
                {...props}
                aria-labelledby="contained-modal-title-vcenter"
                centered
            >
                <Modal.Header closeButton>
                    <Modal.Title>Edit - {props.ad.Name}</Modal.Title>
                </Modal.Header>

                <Form onSubmit={edit}>
                    <Modal.Body>
                        <Row className="mb-3">
                            <Col>
                                <Form.Label>Name</Form.Label>
                                <Form.Control type="text" name="name" defaultValue={props.ad.Name} required />
                            </Col>
                        </Row>
                        <Row className="mb-3">
                            <Col>
                                <Form.Label>Description</Form.Label>
                                <Form.Control as="textarea" rows={2} name="description" defaultValue={props.ad.Description} required />
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
                                            defaultChecked={adsType(type)}
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
                                            defaultChecked={adsSkill(skill)}
                                        />
                                    ))}
                                </div>
                            </Col>
                        </Row>
                    </Modal.Body>
                    <Modal.Footer>
                        <Col className="d-grid">
                            <Button type="submit" className="mb-2">
                                Save changes
                            </Button>
                        </Col>
                    </Modal.Footer>
                </Form>
            </Modal>
        </Container>
    );
}

export default EditAd