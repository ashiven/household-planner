import * as React from "react";
import { createFileRoute } from "@tanstack/react-router";
import Sheet from "@mui/joy/Sheet";

import Accordion from "@mui/joy/Accordion";
import AccordionDetails from "@mui/joy/AccordionDetails";
import AccordionGroup from "@mui/joy/AccordionGroup";
import AccordionSummary from "@mui/joy/AccordionSummary";
import { useQuery } from "@tanstack/react-query";

import { MembersTable, TaskTable } from "../../components/table";

export const Route = createFileRoute("/")({
  component: Index,
});

async function fetchData() {
  let memberData, dailyData, weeklyData, monthlyData;
  const memberRes = await fetch("/members");
  const dailyRes = await fetch("/tasks/daily");
  const weeklyRes = await fetch("/tasks/weekly");
  const monthlyRes = await fetch("/tasks/monthly");

  try {
    memberData = await memberRes.json();
    dailyData = await dailyRes.json();
    weeklyData = await weeklyRes.json();
    monthlyData = await monthlyRes.json();
  } catch {
    memberData = [];
    dailyData = [];
    weeklyData = [];
    monthlyData = [];
  }

  const combinedResponse = {
    members: memberData,
    daily: dailyData,
    weekly: weeklyData,
    monthly: monthlyData,
  };
  return combinedResponse;
}

function Index() {
  const { isLoading, isError, data, error } = useQuery({
    queryKey: ["index"],
    queryFn: fetchData,
    initialData: {
      members: [],
      daily: [],
      weekly: [],
      monthly: [],
    },
  });

  if (isLoading) {
    return <span>Loading...</span>;
  }
  if (isError) {
    return <span>Error loading data: {error.message}</span>;
  }

  const [membersExpanded, setMembersExpanded] = React.useState(true);
  const [dailyExpanded, setDailyExpanded] = React.useState(true);
  const [weeklyExpanded, setWeeklyExpanded] = React.useState(true);
  const [monthlyExpanded, setMonthlyExpanded] = React.useState(true);

  return (
    <Sheet variant="soft" color="neutral" sx={{ p: 4, minHeight: "100vh" }}>
      <AccordionGroup size="lg">
        <Accordion
          expanded={membersExpanded}
          onChange={() => setMembersExpanded(!membersExpanded)}
        >
          <AccordionSummary>Mitglieder</AccordionSummary>
          <AccordionDetails>
            <MembersTable members={data.members} />
          </AccordionDetails>
        </Accordion>
        <Accordion
          expanded={dailyExpanded}
          onChange={() => setDailyExpanded(!dailyExpanded)}
        >
          <AccordionSummary>Tägliche Aufgaben</AccordionSummary>
          <AccordionDetails>
            <TaskTable tasks={data.daily} />
          </AccordionDetails>
        </Accordion>
        <Accordion
          expanded={weeklyExpanded}
          onChange={() => setWeeklyExpanded(!weeklyExpanded)}
        >
          <AccordionSummary>Wöchentliche Aufgaben</AccordionSummary>
          <AccordionDetails>
            <TaskTable tasks={data.weekly} />
          </AccordionDetails>
        </Accordion>
        <Accordion
          expanded={monthlyExpanded}
          onChange={() => setMonthlyExpanded(!monthlyExpanded)}
        >
          <AccordionSummary>Monatliche Aufgaben</AccordionSummary>
          <AccordionDetails>
            <TaskTable tasks={data.monthly} />
          </AccordionDetails>
        </Accordion>
      </AccordionGroup>
    </Sheet>
  );
}
