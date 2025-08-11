import Table from "@mui/joy/Table";

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
    <Table color="primary" variant="soft" aria-label="basic table">
      <thead>
        <tr>
          <th style={{ width: "40%" }}>Name</th>
          <th>Telefonnummer</th>
        </tr>
      </thead>
      <tbody>
        {members.map((member, _idx) => (
          <tr>
            <td>{member.Name}</td>
            <td>{member.Phonenumber}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
}

function TaskTable({ tasks }: { tasks: Task[] }) {
  return (
    <Table color="primary" variant="soft" aria-label="basic table">
      <thead>
        <tr>
          <th style={{ width: "40%" }}>Aufgabe</th>
          <th>Zugewiesen</th>
        </tr>
      </thead>
      <tbody>
        {tasks.map((task, _idx) => (
          <tr>
            <td>{task.Name}</td>
            <td>{task.Assignee?.Name ?? ""}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
}

export { MembersTable, TaskTable };
