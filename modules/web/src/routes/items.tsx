import { createFileRoute } from '@tanstack/react-router'
import * as React from 'react'
import FolderSidebar, { type FolderNode } from '../components/FolderSidebar'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { Select } from '../components/ui/select'
import {
  ChevronRight,
  FolderPlus,
  Grid,
  List,
  Plus,
  Search,
} from 'lucide-react'
import clsx from 'clsx'

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

// Sample items data
interface Item {
  id: string
  name: string
  image?: string
  quantity: number
  value: number
  folderId?: string
}

const sampleItems: Item[] = [
  {
    id: 'item-1',
    name: 'MacBook Pro 16"',
    image: 'https://via.placeholder.com/150',
    quantity: 5,
    value: 2499.99,
    folderId: '1-1-1',
  },
  {
    id: 'item-2',
    name: 'Dell XPS 15',
    image: 'https://via.placeholder.com/150',
    quantity: 3,
    value: 1899.99,
    folderId: '1-1-2',
  },
  {
    id: 'item-3',
    name: 'iPhone 15 Pro',
    image: 'https://via.placeholder.com/150',
    quantity: 10,
    value: 999.99,
    folderId: '1-2',
  },
  {
    id: 'item-4',
    name: 'Office Desk',
    image: 'https://via.placeholder.com/150',
    quantity: 8,
    value: 349.99,
    folderId: '2-1',
  },
]

type DisplayType = 'grid' | 'list'

function findFolderPath(
  folders: FolderNode[],
  targetId: string,
  currentPath: FolderNode[] = []
): FolderNode[] | null {
  for (const folder of folders) {
    const newPath = [...currentPath, folder]
    if (folder.id === targetId) {
      return newPath
    }
    if (folder.children) {
      const result = findFolderPath(folder.children, targetId, newPath)
      if (result) return result
    }
  }
  return null
}

function ItemsPage() {
  const [selectedFolderId, setSelectedFolderId] = React.useState<
    string | undefined
  >()
  const [searchQuery, setSearchQuery] = React.useState('')
  const [sortBy, setSortBy] = React.useState('name')
  const [displayType, setDisplayType] = React.useState<DisplayType>('grid')

  const breadcrumbPath = selectedFolderId
    ? findFolderPath(sampleFolders, selectedFolderId) || []
    : []

  const filteredItems = sampleItems.filter(
    (item) =>
      (!selectedFolderId || item.folderId === selectedFolderId) &&
      item.name.toLowerCase().includes(searchQuery.toLowerCase())
  )

  const sortedItems = [...filteredItems].sort((a, b) => {
    if (sortBy === 'name') return a.name.localeCompare(b.name)
    if (sortBy === 'quantity') return b.quantity - a.quantity
    if (sortBy === 'value') return b.value - a.value
    return 0
  })

  return (
    <div className="flex h-screen">
      <FolderSidebar
        folders={sampleFolders}
        selectedId={selectedFolderId}
        onSelect={setSelectedFolderId}
      />
      <div className="flex-1 flex flex-col overflow-hidden">
        {/* Header Section */}
        <div className="border-b border-gray-200 bg-white p-6">
          {/* Breadcrumb and Action Buttons */}
          <div className="flex items-center justify-between mb-6">
            {/* Breadcrumb */}
            <div className="flex items-center gap-2 text-sm flex-1 min-w-0">
              {breadcrumbPath.length > 0 ? (
                <>
                  {breadcrumbPath.map((folder, index) => (
                    <React.Fragment key={folder.id}>
                      <button
                        onClick={() => setSelectedFolderId(folder.id)}
                        className="font-medium hover:text-gray-600 truncate"
                      >
                        {folder.name}
                      </button>
                      {index < breadcrumbPath.length - 1 && (
                        <ChevronRight className="h-4 w-4 flex-shrink-0 text-gray-400" />
                      )}
                    </React.Fragment>
                  ))}
                </>
              ) : (
                <span className="font-medium text-gray-900">All Items</span>
              )}
            </div>

            {/* Action Buttons */}
            <div className="flex items-center gap-3 ml-4">
              <Button variant="outline" className="gap-2">
                <Plus className="h-4 w-4" />
                Add Item
              </Button>
              <Button variant="outline" className="gap-2">
                <FolderPlus className="h-4 w-4" />
                Add Folder
              </Button>
            </div>
          </div>

          {/* Search and Filters */}
          <div className="flex items-center gap-4">
            {/* Search Bar */}
            <div className="relative flex-1">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-gray-500" />
              <Input
                type="text"
                placeholder="Search items..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="pl-10"
              />
            </div>

            {/* Sort Dropdown */}
            <Select
              value={sortBy}
              onChange={(e) => setSortBy(e.target.value)}
              className="w-40"
            >
              <option value="name">Name</option>
              <option value="quantity">Quantity</option>
              <option value="value">Value</option>
            </Select>

            {/* Display Type Toggle */}
            <div className="flex border border-gray-300 rounded-md overflow-hidden">
              <button
                onClick={() => setDisplayType('grid')}
                className={clsx(
                  'p-2 hover:bg-gray-100',
                  displayType === 'grid' && 'bg-gray-100'
                )}
              >
                <Grid className="h-4 w-4" />
              </button>
              <button
                onClick={() => setDisplayType('list')}
                className={clsx(
                  'p-2 hover:bg-gray-100',
                  displayType === 'list' && 'bg-gray-100'
                )}
              >
                <List className="h-4 w-4" />
              </button>
            </div>
          </div>
        </div>

        {/* Items Grid/List */}
        <div className="flex-1 overflow-y-auto p-6 bg-gray-50">
          {sortedItems.length > 0 ? (
            <div
              className={clsx(
                displayType === 'grid'
                  ? 'grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4'
                  : 'flex flex-col gap-3'
              )}
            >
              {sortedItems.map((item) => (
                <ItemCard
                  key={item.id}
                  item={item}
                  displayType={displayType}
                />
              ))}
            </div>
          ) : (
            <div className="flex items-center justify-center h-64 text-gray-500">
              {searchQuery
                ? 'No items found matching your search'
                : 'No items in this folder'}
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

function ItemCard({
  item,
  displayType,
}: {
  item: Item
  displayType: DisplayType
}) {
  if (displayType === 'list') {
    return (
      <div className="bg-white border border-gray-200 rounded-lg p-4 hover:shadow-md transition-shadow flex items-center gap-4">
        <div className="w-16 h-16 flex-shrink-0 bg-gray-100 rounded-md overflow-hidden">
          {item.image ? (
            <img
              src={item.image}
              alt={item.name}
              className="w-full h-full object-cover"
            />
          ) : (
            <div className="w-full h-full flex items-center justify-center text-gray-400">
              <Package className="h-8 w-8" />
            </div>
          )}
        </div>
        <div className="flex-1 min-w-0">
          <h3 className="font-medium text-gray-900 truncate">{item.name}</h3>
          <div className="flex items-center gap-4 mt-1 text-sm text-gray-600">
            <span>Qty: {item.quantity}</span>
            <span>Total: ${(item.value * item.quantity).toLocaleString()}</span>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="bg-white border border-gray-200 rounded-lg overflow-hidden hover:shadow-md transition-shadow">
      <div className="aspect-square bg-gray-100 overflow-hidden">
        {item.image ? (
          <img
            src={item.image}
            alt={item.name}
            className="w-full h-full object-cover"
          />
        ) : (
          <div className="w-full h-full flex items-center justify-center text-gray-400">
            <Package className="h-16 w-16" />
          </div>
        )}
      </div>
      <div className="p-4">
        <h3 className="font-medium text-gray-900 truncate mb-2">
          {item.name}
        </h3>
        <div className="text-sm text-gray-600 space-y-1">
          <p>Quantity: {item.quantity}</p>
          <p className="font-medium text-gray-900">
            Total: ${(item.value * item.quantity).toLocaleString()}
          </p>
        </div>
      </div>
    </div>
  )
}

function Package({ className }: { className?: string }) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
      className={className}
    >
      <path d="M16.5 9.4 7.55 4.24" />
      <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z" />
      <polyline points="3.29 7 12 12 20.71 7" />
      <line x1="12" x2="12" y1="22" y2="12" />
    </svg>
  )
}
