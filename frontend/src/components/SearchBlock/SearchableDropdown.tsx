import { Button } from '@/components/ui/button'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Input } from '@/components/ui/input'
import { fuzzyMatch } from '@/lib/fuzzySearch'
import { useCallback, useState } from 'react'

interface SearchableDropdownProps {
  items: string[]
  value: string
  onChange: (value: string) => void
  placeholder: string
  label: string
  disabled?: boolean
  isLoading?: boolean
}

export function SearchableDropdown({
  items,
  value,
  onChange,
  placeholder,
  label,
  disabled,
  isLoading
}: SearchableDropdownProps) {
  const [isOpen, setIsOpen] = useState(false)
  const [search, setSearch] = useState('')

  const filteredItems = items.filter(item => 
    fuzzyMatch(item.toLowerCase(), search.toLowerCase())
  )

  const handleSelect = useCallback((item: string) => {
    onChange(item)
    setIsOpen(false)
    setSearch('')
  }, [onChange])

  return (
    <div className="flex-1 space-y-2">
      <label className="text-sm font-medium text-gray-300">
        {label}
      </label>
      <DropdownMenu open={isOpen} onOpenChange={setIsOpen}>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            className="w-full justify-start bg-gray-700 border-gray-600 text-white"
            disabled={disabled || isLoading}
          >
            {value || placeholder}
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-[--radix-dropdown-menu-trigger-width] bg-gray-700 border-gray-600">
          <div className="p-2">
            <Input
              placeholder="Search..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className="bg-gray-800 border-gray-600"
            />
          </div>
          <div className="max-h-[300px] overflow-y-auto">
            {filteredItems.map((item) => (
              <DropdownMenuItem
                key={item}
                onSelect={() => handleSelect(item)}
                className="text-white hover:bg-gray-600"
              >
                {item}
              </DropdownMenuItem>
            ))}
          </div>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
