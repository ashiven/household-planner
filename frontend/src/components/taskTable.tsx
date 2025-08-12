import * as React from "react";
import Table from "@mui/joy/Table";
import Button from "@mui/joy/Button";
import Box from "@mui/joy/Box";
import Modal from "@mui/joy/Modal";
import ModalDialog from "@mui/joy/ModalDialog";
import ModalClose from "@mui/joy/ModalClose";
import Typography from "@mui/joy/Typography";
import Input from "@mui/joy/Input";

interface Member {
  Name: string;
  Phonenumber: string;
}

interface Task {
  Name: string;
  Assignee: Member | null;
}

export default function TaskTable({
  id,
  tasks,
}: {
  id: string;
  tasks: Task[];
}) {
  const [tableData, setTableData] = React.useState<Task[]>(tasks);
  const [open, setOpen] = React.useState(false);
  const [editIndex, setEditIndex] = React.useState<number | null>(null);
  const [formData, setFormData] = React.useState<Task>({
    Name: "",
    Assignee: null,
  });

  React.useEffect(() => {
    setTableData(tasks);
  }, [tasks]);

  React.useEffect(() => {
    if (tableData.length > 0) {
      saveChanges();
    }
  }, [tableData]);

  const authorize = async (action: Function) => {
    const password = prompt("Bitte Passwort eingeben:");

    if (!password) {
      return;
    }

    const response = await fetch("/auth", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ password }),
    });

    if (response.ok) {
      action();
    } else {
      alert("Falsches Passwort.");
    }
  };

  const startEdit = (index: number) => {
    setEditIndex(index);
    setFormData(tableData[index]);
    setOpen(true);
  };

  const saveEdit = () => {
    if (editIndex !== null) {
      const newData = [...tableData];
      newData[editIndex] = formData;
      setTableData(newData);
    } else {
      const newData = [...tableData, formData];
      setTableData(newData);
    }
    setOpen(false);
    setEditIndex(null);
  };

  const deleteTask = (index: number) => {
    setTableData((prevData) => prevData.filter((_, i) => i !== index));
  };

  const startAdd = () => {
    setEditIndex(null);
    setFormData({ Name: "", Assignee: null });
    setOpen(true);
  };

  const saveChanges = async () => {
    try {
      console.log("Saving changes to tasks:", JSON.stringify(tableData));
      const response = await fetch(`/tasks/${id}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(tableData),
      });

      if (!response.ok) {
        throw new Error("Failed to save changes");
      }
    } catch (error) {
      console.error("Error saving tasks:", error);
    }
  };

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
          {tableData.map((task, index) => (
            <tr key={index}>
              <td>{task.Name}</td>
              <td>{task.Assignee?.Name ?? ""}</td>
              <td>
                <Box sx={{ display: "flex", gap: 2, flexWrap: "wrap" }}>
                  <Button
                    color="neutral"
                    onClick={() => authorize(() => startEdit(index))}
                  >
                    Bearbeiten
                  </Button>
                  <Button
                    variant="soft"
                    color="danger"
                    onClick={() => authorize(() => deleteTask(index))}
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
        <Button color="primary" onClick={() => authorize(startAdd)}>
          Hinzufügen
        </Button>
      </Box>

      <Modal open={open} onClose={() => setOpen(false)}>
        <ModalDialog>
          <ModalClose />
          <Typography level="h4" sx={{ mb: 2 }}>
            Aufgabe
          </Typography>
          <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
            <Input
              placeholder="Name"
              value={formData.Name}
              onChange={(e) =>
                setFormData({ ...formData, Name: e.target.value })
              }
            />
            <Button color="success" onClick={saveEdit}>
              Speichern
            </Button>
          </Box>
        </ModalDialog>
      </Modal>
    </Box>
  );
}
