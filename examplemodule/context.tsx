
import { createContext, useContext } from 'react';
import { IExamplemoduleContext } from './typings';

export const ExamplemoduleContext = createContext<IExamplemoduleContext>(null);

export function useExamplemoduleContext(): IExamplemoduleContext {
	return useContext(ExamplemoduleContext);
}
