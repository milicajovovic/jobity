import axios from "axios";
import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Col, Container, Row, Card, Button, Table } from "react-bootstrap";
import AdminMenu from "../components/AdminMenu";
import Confirmation from "../components/Confirmation";

function AdminEmployees() {
    const [employees, setEmployees] = useState([]);
    const [question, setQuestion] = useState("");
    const [id, setId] = useState();
    const [action, setAction] = useState("");
    const [show, setShow] = useState(false);

    useEffect(() => {
        axios.get("http://localhost:3007/employees", { headers: { Authorization: localStorage.getItem("jwt") } }).then(res => {
            setEmployees(res.data);
        });
    }, [employees]);

    const convertDate = (birthday) => {
        let date = new Date(birthday);
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

    const blockEmployee = (employee) => {
        setQuestion("Are you sure you want to block " + employee.FirstName + " " + employee.LastName + "?");
        setId(employee.ID);
        setAction("employees/block");
        setShow(true);
    };

    const deleteEmployee = (employee) => {
        setQuestion("Are you sure you want to delete " + employee.FirstName + " " + employee.LastName + "?");
        setId(employee.ID);
        setAction("employees/delete");
        setShow(true);
    };

    return (
        <Container fluid>
            <Row>
                <AdminMenu />
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    <Card body style={{ width: "50rem" }}>
                        <Card.Title className="text-center mt-3" as="h3">Employees</Card.Title>
                        <Card.Body className="text-center">
                            <Table striped bordered>
                                <thead>
                                    <tr>
                                        <th>Email</th>
                                        <th>First name</th>
                                        <th>Last name</th>
                                        <th>Birthday</th>
                                        <th>Block</th>
                                        <th>Delete</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {employees.map(employee => (
                                        <tr key={employee.ID}>
                                            <td>{employee.Email}</td>
                                            <td>{employee.FirstName}</td>
                                            <td>{employee.LastName}</td>
                                            <td>{convertDate(employee.Birthday)}</td>
                                            <td>
                                                <Button onClick={() => blockEmployee(employee)} disabled={employee.Blocked || employee.Deleted}>
                                                    Block
                                                </Button>
                                            </td>
                                            <td>
                                                <Button onClick={() => deleteEmployee(employee)} disabled={employee.Deleted}>
                                                    Delete
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
            <Confirmation show={show} onHide={() => setShow(false)} question={question} id={id} action={action} close={setShow}/>
        </Container>
    )
}

export default AdminEmployees