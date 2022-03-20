import React, {useState} from 'react';

const Counter = function (){
    const [state, setState] = useState(0)
    const [value, setValue] = useState('Текст в инпуте')

    function inc(){
        setState(state + 1)
    }

    function dec(){
        setState(state - 1)
    }

    return(
        <div>
            <h1>{state}</h1>
            <h1>{value}</h1>
            <input
                type="text"
                value={value}
                onChange={event => setValue(event.target.value)}
            />
            <button onClick={inc}>INC</button>
            <button onClick={dec}>DEC</button>
        </div>
    )
};

export default Counter;