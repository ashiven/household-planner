import Table from "@mui/joy/Table";
import Button from "@mui/joy/Button";
import Box from "@mui/joy/Box";

interface Member {
  Name: string;
  Phonenumber: string;
}

interface Task {
  Name: string;
  Assignee: Member | null;
}

function MembersTable({ members }: { members: Member[] }) {
  return (
    <Box sx={{ mb: 2 }}>
      <Table color="primary" variant="soft" aria-label="basic table">
        <thead>
          <tr>
            <th style={{ width: "40%" }}>Name</th>
            <th>Telefonnummer</th>
            <th>Aktionen</th>
          </tr>
        </thead>
        <tbody>
          {members.map((member, _idx) => (
            <tr>
              <td>{member.Name}</td>
              <td>{member.Phonenumber}</td>
              <td>
                <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
                  <Button color="neutral" onClick={function () {}}>
                    Bearbeiten
                  </Button>
                  <Button
                    variant="soft"
                    color="danger"
                    onClick={function () {}}
                  >
                    Löschen
                  </Button>
                </Box>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
      <Box sx={{ marginTop: 2 }}>
        <Button color="primary" onClick={function () {}}>
          Hinzufügen
        </Button>
      </Box>
    </Box>
  );
}

function TaskTable({ tasks }: { tasks: Task[] }) {
  return (
    <Box sx={{ mb: 2 }}>
      <Table color="primary" variant="soft" aria-label="basic table">
        <thead>
          <tr>
            <th style={{ width: "40%" }}>Aufgabe</th>
            <th>Zugewiesen</th>
            <th>Aktionen</th>
          </tr>
        </thead>
        <tbody>
          {tasks.map((task, _idx) => (
            <tr>
              <td>{task.Name}</td>
              <td>{task.Assignee?.Name ?? ""}</td>
              <td>
                <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
                  <Button color="neutral" onClick={function () {}}>
                    Bearbeiten
                  </Button>
                  <Button
                    variant="soft"
                    color="danger"
                    onClick={function () {}}
                  >
                    Löschen
                  </Button>
                </Box>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
      <Box sx={{ marginTop: 2 }}>
        <Button color="primary" onClick={function () {}}>
          Hinzufügen
        </Button>
      </Box>
    </Box>
  );
}

export { MembersTable, TaskTable };
