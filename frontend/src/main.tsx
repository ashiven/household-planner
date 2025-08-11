import { StrictMode } from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createRouter } from "@tanstack/react-router";
import "@fontsource/inter";
import { routeTree } from "./routeTree.gen";
import { CssVarsProvider } from "@mui/joy/styles";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { GlobalStyles } from "@mui/joy";

const router = createRouter({ routeTree });

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const rootElement = document.getElementById("root")!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <StrictMode>
      <QueryClientProvider client={new QueryClient()}>
        <CssVarsProvider defaultMode="dark">
          <GlobalStyles
            styles={{
              html: { margin: 0, padding: 0 },
              body: { margin: 0, padding: 0 },
            }}
          />
          <RouterProvider router={router} />
        </CssVarsProvider>
      </QueryClientProvider>
    </StrictMode>,
  );
}
