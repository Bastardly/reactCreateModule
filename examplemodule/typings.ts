
import { Dispatch } from 'react';

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


export type IDispatch = Dispatch<IAction>;

export interface IExamplemodule {};
export interface IExamplemoduleState extends Record<string, any> {};

export interface IExamplemoduleContext {
	state: IExamplemoduleState;
	actions: {};
	dispatch: IDispatch;
}
