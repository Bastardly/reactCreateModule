
import { IDispatch, IExamplemoduleState } from './typings';

export const mapActions = (dispatch: IDispatch) => ({
	update: (payload: Partial<IExamplemoduleState>) =>
		dispatch({ type: 'update', payload }),
});
