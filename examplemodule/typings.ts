
import { Dispatch } from 'react';
import { mapActions } from './actions';

type EmptyActionType<T extends string> = {
	type: T;
};

type PayloadType<Payload, Optional extends boolean> = Optional extends true
? { payload?: Payload }
: { payload: Payload };

type ActionType<
	T extends string,
	Payload,
	Optional extends boolean = false
> = EmptyActionType<T> & PayloadType<Payload, Optional>;

export type IAction = 
	| ActionType<'update', Partial<IExamplemoduleState>>

type IActions = ReturnType<typeof mapActions> & {
	// Unmapped or actions from props here
};


export type IDispatch = Dispatch<IAction>;

export interface IExamplemodule {};
export interface IExamplemoduleState extends Record<string, any> {};

export interface IExamplemoduleContext {
	state: IExamplemoduleState;
	actions: IActions;
	dispatch: IDispatch;
}
