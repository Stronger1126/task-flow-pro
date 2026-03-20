'use strict';

const express = require('express');
const router = express.Router();
const TaskController = require('../controllers/taskController');

// Create a new task
router.post('/tasks', TaskController.create);

// Get all tasks
router.get('/tasks', TaskController.getAll);

// Get a specific task by ID
router.get('/tasks/:id', TaskController.getById);

// Update a specific task by ID
router.put('/tasks/:id', TaskController.update);

// Delete a specific task by ID
router.delete('/tasks/:id', TaskController.delete);

module.exports = router;