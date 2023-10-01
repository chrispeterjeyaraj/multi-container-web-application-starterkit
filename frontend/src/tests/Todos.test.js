import { render, waitFor, screen } from "@testing-library/react";
import Todos from "../features/Todos";
import axios from "axios";

jest.mock("axios");

const dummyTodos = [
  {
    id: 1,
    name: "todo 1",
    iscompleted: false,
  },
  {
    id: 2,
    name: "todo 2",
    iscompleted: false,
  },
  {
    id: 3,
    name: "todo 3",
    iscompleted: false,
  },
];

test("todos list", async () => {
  axios.get.mockResolvedValue({ data: dummyTodos });
  render(<Todos />);

  const todoList = await waitFor(() => screen.findAllByTestId("todo"));

  expect(todoList).toHaveLength(3);
});
