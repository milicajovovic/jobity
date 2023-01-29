import axios from "axios";
import React from "react";
import { Form, Button } from "react-bootstrap";

function SearchForm({ changeAds }) {
    const search = (event) => {
        event.preventDefault();

        var searchName = event.target.name.value
        if (searchName === "") {
            searchName = "Name"
        }
        var searchDescription = event.target.description.value
        if (searchDescription === "") {
            searchDescription = "Description"
        }

        axios.get("http://localhost:3007/ads/search/" + searchName + "/" + searchDescription).then(res => {
            changeAds(res.data);
        });
    }

    return (
        <Form onSubmit={search} className="d-flex">
            <Form.Control type="text" name="name" placeholder="Name" className="me-2" />
            <Form.Control type="text" name="description" placeholder="Description" className="me-2" />
            <Button type="submit">Search</Button>
        </Form>
    )
}

export default SearchForm