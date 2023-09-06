import { initializeApp } from "firebase/app";
import { getAuth, GoogleAuthProvider } from "firebase/auth";

const firebaseConfig = {};

export const firebaseAuth = () => {
    return getAuth(initializeApp(firebaseConfig));
};

export const googleProvider = () => {
    const provider = new GoogleAuthProvider();
    provider.addScope("https://www.googleapis.com/auth/calendar.events");

    return provider;
};
