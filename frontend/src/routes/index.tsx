import { createFileRoute } from "@tanstack/react-router";
import * as React from "react";
import Button from "@mui/joy/Button";
import Stack from "@mui/joy/Stack";

export const Route = createFileRoute("/")({
  component: Index,
});

function Index() {
  return (
    <Stack spacing={3} sx={{ justifyContent: "center", alignItems: "center" }}>
      <Button variant="solid">Members</Button>
      <Button variant="solid">Daily Tasks</Button>
      <Button variant="solid">Weekly Tasks</Button>
      <Button variant="solid">Monthly Tasks</Button>
    </Stack>
  );
}
