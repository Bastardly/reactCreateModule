
import React, { useReducer, useMemo } from 'react';
import { IExamplemodule, IExamplemoduleContext } from './typings';
import { ExamplemoduleReducer, initialState } from './reducer';
import { ExamplemoduleContext, useExamplemoduleContext } from './context';
import { mapActions } from './actions';


const TestComponent = () => {
	const { state, actions, dispatch } = useExamplemoduleContext();

	return <div>Hello World</div>;
}

export function Examplemodule(props: IExamplemodule) {
	const [state, dispatch] = useReducer(ExamplemoduleReducer , initialState);

	const value = useMemo((): IExamplemoduleContext => {
		return {
			state,
			actions: mapActions(dispatch),
			dispatch,
		};
	}, [state]);

	return (
		<ExamplemoduleContext.Provider value={ value }>
			<TestComponent />
		</ExamplemoduleContext.Provider>
	);
}
