import React, {useState,useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.css'
import {Button,Card,Row,Col} from 'react-bootstrap'

const Entry=({entryData,setChangeIngredient,deleteSingleEntry,setChangeEntry})=>{
    return(
        <Card>
            <Row>
                <Col>Dish:{entryData !== undefined && entryData.dish}</Col>
                <Col>Ingredients:{entryData !== undefined && entryData.ingredients}</Col>
                <Col>Calories:{entryData !== undefined && entryData.calories}</Col>
                <Col>Fat:{entryData !== undefined && entryData.fat}</Col>
                <Col> <Button onClick={() => deleteSingleEntry(entryData._id)} delete entry > </Button> </Col>
                <Col> <Button onClick={() => setChangeIngredient()} change ingredients > </Button></Col>
                <Col> <Button onClick={() => setChangeEntry()} change entry > </Button> </Col>
            </Row>
        </Card>
    )
}