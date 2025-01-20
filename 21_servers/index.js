// initialization 
const express = require("express");
const port = 3001; 
const app = express(); 
// middlewares 
app.use(express.json())
app.use(express.urlencoded({extended:true}))
// get request 1 
app.get('/', (req,res)=>{
    res.status(200).send("Welcome to go lang website!")
});
// get request 2 
app.get('/get',(req,res)=>{
    res.status(200).json({message:"Hello from Priya"})
    // res.status(200).send("This is the welcome message of the custom get path")
});
// post request 1 
app.post('/post',(req,res)=>{
    const jsonBody = req.body
    res.status(200).send(jsonBody)
}); 
// post request 2 
app.post('/postform',(req,res)=>{
    res.status(200).send(JSON.stringify(req.body))
})

// app starting 
app.listen(port,()=>{
    console.log(`Example of a web server running at http://localhost:${port}`);
}); 