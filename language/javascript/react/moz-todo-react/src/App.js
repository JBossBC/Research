import React, {useRef, useEffect, useState} from "react";
import Form from "./components/Form";
import FilterButton from "./components/FilterButton";
import Todo from "./components/Todo";
import {nanoid} from "nanoid";

function App(props) {
    const [filter, setFilter] = useState('All');
    const [tasks,setTasks]=useState(props.tasks);
    const FILTER_MAP={
        All:()=>true,
        Active:(task)=>!task.completed,
        Completed:(task)=>task.completed
    }
    const FILTER_NAMES = Object.keys(FILTER_MAP);

    function toggleTaskCompleted(id) {
        const updatedTasks=tasks.map((task)=>{
            if (id == task.id){
                return{...task,completed: !task.completed}
            }
            return task;
        });
        setTasks(updatedTasks)
    }
    function deleteTask(id) {
        const remainingTasks=tasks.filter((task)=>id!=task.id);
        setTasks(remainingTasks)
    }
    function editTask(id,newName){
        const editedTaskList=tasks.map((task)=>{
            if(id==task.id){
                return{...task,name:newName}
            }
            return task
        });
        setTasks(editedTaskList);
    }
    const filterList = FILTER_NAMES.map((name) => (
        <FilterButton
            key={name}
            name={name}
            isPressed={name === filter}
            setFilter={setFilter}
        />
    ));
    const taskList = tasks
        .filter(FILTER_MAP[filter])
        .map((task) => (
            <Todo
                id={task.id}
                name={task.name}
                completed={task.completed}
                key={task.id}
                toggleTaskCompleted={toggleTaskCompleted}
                deleteTask={deleteTask}
                editTask={editTask}
            />
        ));
    function usePrevious(value) {
        const ref = useRef();
        useEffect(() => {
            ref.current = value;
        });
        return ref.current;
    }
    const prevTaskLength=usePrevious(tasks.length);
   useEffect(()=>{
       if(tasks.length-prevTaskLength === -1){
             listHeadingRef.current.focus();
       }
   },[tasks.length,prevTaskLength])
    function addTask(name) {
        const newTask = { id: `todo-${nanoid()}`, name, completed: false };
        setTasks([...tasks, newTask]);
    }
    const tasksNoun = taskList.length !== 1 ? 'tasks' : 'task';
    const headingText = `${taskList.length} ${tasksNoun} remaining`;
    const listHeadingRef=useRef(null);
    return (
        <div className="todoapp stack-large">
            <h1>TodoMatic</h1>
            <Form addTask={addTask} />
            <div className="filters btn-group stack-exception">
                {filterList}
            </div>
            <h2 id="list-heading" ref={listHeadingRef} tabIndex="-1">{headingText}</h2>
            <ul
                role="list"
                className="todo-list stack-large stack-exception"
                aria-labelledby="list-heading"
            >
                {taskList}
            </ul>
        </div>
    );
}

export default App;
