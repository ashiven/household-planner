import { createRootRoute, Link, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

import Button from "@mui/joy/Button";

export const Route = createRootRoute({
  component: () => (
    <>
      <div className="p-2 flex gap-2">
        <Link to="/" className="[&.active]:font-bold">
          <Button variant="solid">Home</Button>
        </Link>{" "}
        <Link to="/about" className="[&.active]:font-bold">
          <Button variant="solid">About</Button>
        </Link>
      </div>
      <hr />
      <Outlet />
      <TanStackRouterDevtools />
    </>
  ),
});
