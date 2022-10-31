import { useEffect, useState, createContext, useContext } from "react";
import { Snackbar } from "@material-ui/core";

export const FloatingContext = createContext();

export const FloatingContextProvider = ({ children }) => {
	const [state, setState] = useState({
		open: false,
		text: "",
	});

	const handlers = {
		openSnackBar: (text) => {
			setState({
				open: true,
				text,
			});
		},
	};

	const onClose = () => {
		setState((state) => ({
			...state,
			open: false,
		}));
	};

	return (
		<FloatingContext.Provider value={handlers}>
			<Snackbar
				anchorOrigin={{
					vertical: "bottom",
					horizontal: "center",
				}}
				open={Boolean(state.open)}
				onClose={onClose}
				autoHideDuration={6000}
				message={state.text}
			/>

			{children}
		</FloatingContext.Provider>
	);
};