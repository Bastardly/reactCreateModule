package main

import (
	"os"
	"text/template"
)

type IBaseTemplateData struct {
	ModuleName string
}

type IExtendedTemplateData struct {
	ModuleName string
	ComponentFileName string
}

type IDataTypes interface {
    IBaseTemplateData | IExtendedTemplateData
}

type ITemplateInfo struct {
	Path string
	FileName string
	TempateString string
}

type ITemplate[T IDataTypes] struct {
	Info ITemplateInfo
	Data T
}


func createTemplate[T IDataTypes](data ITemplate[T]) {
	// Reads our templateString as a new template
	t, err := template.New(data.Info.FileName).Parse(data.Info.TempateString)
	if err != nil {
		panic(err)
	}

	// Create our template as a new file
	file, err := os.Create(data.Info.Path + "/" + data.Info.FileName)
	if err != nil {
	  panic(err)
  }

  t.Execute(file, data.Data)
}

const typingsTemplate = `
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
	| ActionType<'update', Partial<I{{ .ModuleName}}State>>


export type IDispatch = Dispatch<IAction>;

export interface I{{ .ModuleName}} {};
export interface I{{ .ModuleName}}State extends Record<string, any> {};

export interface I{{ .ModuleName}}Context {
	state: I{{ .ModuleName}}State;
	actions: {};
	dispatch: IDispatch;
}
`


func createTypingsTemplate(path string) {
	data := ITemplate[IBaseTemplateData] {
		Info: ITemplateInfo{ path, "typings.ts",typingsTemplate},
		Data: IBaseTemplateData{moduleName},
	}

	createTemplate(data)
}


const componentTemplate = `
import React, { useReducer, useMemo } from 'react';
import { I{{ .ModuleName}}, I{{ .ModuleName}}Context } from './typings';
import { {{ .ModuleName}}Reducer, initialState } from './reducer';
import { {{ .ModuleName}}Context, use{{ .ModuleName}}Context } from './context';
import { mapActions } from './actions';


const TestComponent = () => {
	const { state, actions, dispatch } = use{{ .ModuleName}}Context();

	return <div>Hello World</div>;
}

export function {{ .ModuleName}}(props: I{{ .ModuleName}}) {
	const [state, dispatch] = useReducer({{ .ModuleName}}Reducer , initialState);

	const value = useMemo((): I{{ .ModuleName}}Context => {
		return {
			state,
			actions: mapActions(dispatch),
			dispatch,
		};
	}, [state]);

	return (
		<{{ .ModuleName}}Context.Provider value={ value }>
			<TestComponent />
		</{{ .ModuleName}}Context.Provider>
	);
}
`

func createComponentTemplate(path, componentFileName string) {
	data := ITemplate[IBaseTemplateData] {
		Info: ITemplateInfo{ path, componentFileName, componentTemplate},
		Data: IBaseTemplateData{moduleName},
	}

	createTemplate(data)
}

const contextTemplate = `
import { createContext, useContext } from 'react';
import { I{{ .ModuleName}}Context } from './typings';

export const {{ .ModuleName}}Context = createContext<I{{ .ModuleName}}Context>(null);

export function use{{ .ModuleName}}Context(): I{{ .ModuleName}}Context {
	return useContext({{ .ModuleName}}Context);
}
`

func createContextTemplate(path string) {
	data := ITemplate[IBaseTemplateData] {
		Info: ITemplateInfo{ path, "context.tsx", contextTemplate},
		Data: IBaseTemplateData{moduleName},
	}

	createTemplate(data)
}

const reducerTemplate = `
import { I{{ .ModuleName}}State, IAction } from './typings';

export const initialState: I{{ .ModuleName}}State = {};

export function {{ .ModuleName}}Reducer(state: I{{ .ModuleName}}State, action: IAction) {
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
`

func createReducerTemplate(path string) {
	data := ITemplate[IBaseTemplateData] {
		Info: ITemplateInfo{ path, "reducer.tsx", reducerTemplate},
		Data: IBaseTemplateData{moduleName},
	}

	createTemplate(data)
}

const indexTemplate = `
export { {{ .ModuleName}} } from './{{ .ComponentFileName}}';
`

func createIndexTemplate(path, componentFileName string) {
	data := ITemplate[IExtendedTemplateData] {
		Info: ITemplateInfo{ path, "index.ts", indexTemplate},
		Data: IExtendedTemplateData{moduleName, componentFileName},
	}

	createTemplate(data)
}

const actionsTemplate = `
import { IDispatch, I{{ .ModuleName}}State } from './typings';

export const mapActions = (dispatch: IDispatch) => ({
	update: (payload: Partial<I{{ .ModuleName}}State>) =>
		dispatch({ type: 'update', payload }),
});
`

func createActionsTemplate(path string) {
	data := ITemplate[IBaseTemplateData] {
		Info: ITemplateInfo{ path, "actions.ts", actionsTemplate},
		Data: IBaseTemplateData{moduleName},
	}

	createTemplate(data)
}