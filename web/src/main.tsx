import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "@/styles/globals.css";
import { Routes } from '@generouted/react-router';
import { ThemeProvider } from './components/theme-provider';

function RenderWrapper() {

  return (
  <StrictMode>
    <ThemeProvider defaultTheme='dark' storageKey='echo-rss-theme'>
      <Routes />
    </ThemeProvider>
  </StrictMode>
  )
}

createRoot(document.getElementById("root")!).render(
  <RenderWrapper />
);
