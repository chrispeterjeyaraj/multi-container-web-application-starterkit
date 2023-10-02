import React from 'react';
import { render, waitFor } from '@testing-library/react';
import Todos from '../features/Todos';
import { getRequest } from '../app/AxiosClient'; // Mock AxiosClient

// Mock AxiosClient function
jest.mock('../app/AxiosClient', () => ({
  getRequest: jest.fn(),
}));

describe('Todos Component', () => {
  it('should render a list of todos', async () => {
    const mockData = [
      { ID: 1, name: 'Todo 1' },
      { ID: 2, name: 'Todo 2' },
    ];

    // Mock the getRequest function to resolve with the mock data
    getRequest.mockResolvedValue({ data: mockData });

    // Render the component
    const { getByText } = render(<Todos />);

    // Wait for the component to load and render
    await waitFor(() => {
      // Assert that each todo name is present in the rendered content
      mockData.forEach((todo) => {
        expect(getByText(todo.name)).toBeInTheDocument();
      });
    });
  });
});
