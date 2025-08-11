import Table from "@mui/joy/Table";

function MembersTable(members: any) {
  return (
    <Table color="primary" variant="soft" aria-label="basic table">
      <thead>
        <tr>
          <th style={{ width: "40%" }}>Name</th>
          <th>Telefonnummer</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>Jannik</td>
          <td>2134902304</td>
        </tr>
        <tr>
          <td>Luka</td>
          <td>192030213</td>
        </tr>
      </tbody>
    </Table>
  );
}

function TaskTable(tasks: any) {
  return (
    <Table color="primary" variant="soft" aria-label="basic table">
      <thead>
        <tr>
          <th>Aufgabe</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>Putzen</td>
        </tr>
        <tr>
          <td>Wischen</td>
        </tr>
      </tbody>
    </Table>
  );
}

export { MembersTable, TaskTable };
