import { createFileRoute } from '@tanstack/react-router'
import * as React from 'react'
import FolderSidebar, { type FolderNode } from '../components/FolderSidebar'

export const Route = createFileRoute('/items')({
  component: ItemsPage,
})

// Sample folder data - replace with your actual data fetching logic
const sampleFolders: FolderNode[] = [
  {
    id: '1',
    name: 'Electronics',
    children: [
      {
        id: '1-1',
        name: 'Laptops',
        children: [
          { id: '1-1-1', name: 'Gaming' },
          { id: '1-1-2', name: 'Business' },
        ],
      },
      {
        id: '1-2',
        name: 'Phones',
      },
    ],
  },
  {
    id: '2',
    name: 'Furniture',
    children: [
      { id: '2-1', name: 'Office' },
      { id: '2-2', name: 'Home' },
    ],
  },
  {
    id: '3',
    name: 'Supplies',
  },
]

function ItemsPage() {
  const [selectedFolderId, setSelectedFolderId] = React.useState<
    string | undefined
  >()

  return (
    <div className="flex h-screen">
      <FolderSidebar
        folders={sampleFolders}
        selectedId={selectedFolderId}
        onSelect={setSelectedFolderId}
      />
      <div className="flex-1 p-6 overflow-y-auto">
        <div className="max-w-7xl mx-auto">
          <h1 className="text-3xl font-bold mb-6">All Items</h1>
          {selectedFolderId ? (
            <div className="bg-gray-50 border border-gray-200 rounded-lg p-4">
              <p className="text-sm text-gray-600">
                Selected folder: <span className="font-medium">{selectedFolderId}</span>
              </p>
              <p className="text-sm text-gray-500 mt-2">
                Items in this folder will be displayed here
              </p>
            </div>
          ) : (
            <div className="bg-gray-50 border border-gray-200 rounded-lg p-4">
              <p className="text-sm text-gray-500">
                Select a folder to view its items
              </p>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
