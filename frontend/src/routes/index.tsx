import { createFileRoute } from "@tanstack/react-router";
import Sheet from "@mui/joy/Sheet";

import Accordion from "@mui/joy/Accordion";
import AccordionDetails from "@mui/joy/AccordionDetails";
import AccordionGroup from "@mui/joy/AccordionGroup";
import AccordionSummary from "@mui/joy/AccordionSummary";

export const Route = createFileRoute("/")({
  component: Index,
});

function Index() {
  return (
    <Sheet variant="soft" color="primary" sx={{ p: 4 }}>
      <AccordionGroup size="lg">
        <Accordion>
          <AccordionSummary>Members</AccordionSummary>
          <AccordionDetails>Content</AccordionDetails>
        </Accordion>
        <Accordion>
          <AccordionSummary>Daily Tasks</AccordionSummary>
          <AccordionDetails>Content</AccordionDetails>
        </Accordion>
        <Accordion>
          <AccordionSummary>Weekly Tasks</AccordionSummary>
          <AccordionDetails>Content</AccordionDetails>
        </Accordion>
        <Accordion>
          <AccordionSummary>Monthly Tasks</AccordionSummary>
          <AccordionDetails>Content</AccordionDetails>
        </Accordion>
      </AccordionGroup>
    </Sheet>
  );
}
