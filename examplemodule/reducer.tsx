
import { IExamplemoduleState, IAction } from './typings';

export const initialState: IExamplemoduleState = {};

export function ExamplemoduleReducer(state: IExamplemoduleState, action: IAction) {
	switch (action.type) {
	case 'update': {
		return {
			...state,
			...action.payload,
		}
	}
		default: {
			return state
		}
	}
}
