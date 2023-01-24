import axios from "axios";
import React, { useEffect, useState } from "react";
import { Col, Container, Row, Card, Button, Navbar } from "react-bootstrap";
import NavigationBar from "../components/NavigationBar";
import AdCard from "../components/AdCard";
import SearchForm from "../components/SearchForm";
import FilterForm from "../components/FilterForm";

function Home() {
    const [allAds, setAllAds] = useState([]);
    const [shownAds, setShownAds] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:3001/").then(res => {
            setAllAds(res.data);
            setShownAds(res.data);
        });
    }, []);

    const sortAds = () => {
        setShownAds([].concat(shownAds).sort((a, b) => new Date(b.Posted) - new Date(a.Posted)))
    }

    return (
        <Container fluid>
            <Row>
                <NavigationBar />
            </Row>
            <Row>
                <Navbar bg="light" className="searchBar">
                    <Container>
                        <SearchForm changeAds={setShownAds} />
                        <FilterForm ads={allAds} changeAds={setShownAds} />
                        <Button onClick={sortAds}>Sort by date</Button>
                    </Container>
                </Navbar>
            </Row>
            <Row className="d-flex justify-content-center h-100 pt-5">
                <Col md="auto">
                    {shownAds.length > 0 ?
                        shownAds.map(a => (
                            <AdCard ad={a} key={a.ID} />
                        ))
                        : <Card body style={{ width: "40rem" }} className="text-center">There is no ads for chosen criteria</Card>
                    }
                </Col>
            </Row>
        </Container>
    )
}

export default Home