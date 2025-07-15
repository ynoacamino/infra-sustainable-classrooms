'use client';

import { useState } from 'react';
import {
  Plus,
  Edit,
  Trash2,
  Tag,
  Folder,
  Settings,
  BarChart3,
  Users,
  Eye,
} from 'lucide-react';
import { Button } from '@/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/ui/card';
import { Badge } from '@/ui/badge';
import { Separator } from '@/ui/separator';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/ui/dialog';
import type { VideoCategory, VideoTag } from '@/types/video_learning/models';
import { CreateCategoryForm } from '@/components/video_learning/forms/create-category-form';
import { CreateTagForm } from '@/components/video_learning/forms/create-tag-form';

interface VideoManagementProps {
  categories: VideoCategory[];
  tags: VideoTag[];
}

export function VideoManagement({
  categories: initialCategories,
  tags: initialTags,
}: VideoManagementProps) {
  const [categories, setCategories] =
    useState<VideoCategory[]>(initialCategories);
  const [tags, setTags] = useState<VideoTag[]>(initialTags);
  const [categoryDialogOpen, setCategoryDialogOpen] = useState(false);
  const [tagDialogOpen, setTagDialogOpen] = useState(false);

  const handleCategorySuccess = (category: VideoCategory) => {
    setCategories((prev) => [...prev, category]);
    setCategoryDialogOpen(false);
  };

  const handleTagSuccess = (tag: VideoTag) => {
    setTags((prev) => [...prev, tag]);
    setTagDialogOpen(false);
  };

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
      {/* Categories Management */}
      <Card>
        <CardHeader>
          <div className="flex items-center justify-between">
            <div>
              <CardTitle className="flex items-center gap-2">
                <Folder className="h-5 w-5" />
                Categories
              </CardTitle>
              <CardDescription>
                Manage video categories for better organization
              </CardDescription>
            </div>
            <Dialog
              open={categoryDialogOpen}
              onOpenChange={setCategoryDialogOpen}
            >
              <DialogTrigger asChild>
                <Button size="sm">
                  <Plus className="h-4 w-4 mr-2" />
                  Add Category
                </Button>
              </DialogTrigger>
              <DialogContent>
                <DialogHeader>
                  <DialogTitle>Create New Category</DialogTitle>
                  <DialogDescription>
                    Add a new category to organize your videos
                  </DialogDescription>
                </DialogHeader>
                <CreateCategoryForm
                  onSuccess={handleCategorySuccess}
                  onCancel={() => setCategoryDialogOpen(false)}
                />
              </DialogContent>
            </Dialog>
          </div>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            {categories.length === 0 ? (
              <div className="text-center py-8">
                <Folder className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
                <p className="text-sm text-muted-foreground">
                  No categories yet
                </p>
                <p className="text-xs text-muted-foreground">
                  Create your first category to get started
                </p>
              </div>
            ) : (
              categories.map((category) => (
                <div
                  key={category.id}
                  className="flex items-center justify-between p-3 border rounded-lg"
                >
                  <div className="flex items-center gap-3">
                    <Folder className="h-4 w-4 text-muted-foreground" />
                    <span className="font-medium">{category.name}</span>
                  </div>
                  <div className="flex items-center gap-1">
                    <Button variant="ghost" size="sm">
                      <Edit className="h-4 w-4" />
                    </Button>
                    <Button
                      variant="ghost"
                      size="sm"
                      className="text-destructive hover:text-destructive"
                    >
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </div>
                </div>
              ))
            )}
          </div>
        </CardContent>
      </Card>

      {/* Tags Management */}
      <Card>
        <CardHeader>
          <div className="flex items-center justify-between">
            <div>
              <CardTitle className="flex items-center gap-2">
                <Tag className="h-5 w-5" />
                Tags
              </CardTitle>
              <CardDescription>
                Create and manage tags for video classification
              </CardDescription>
            </div>
            <Dialog open={tagDialogOpen} onOpenChange={setTagDialogOpen}>
              <DialogTrigger asChild>
                <Button size="sm">
                  <Plus className="h-4 w-4 mr-2" />
                  Add Tag
                </Button>
              </DialogTrigger>
              <DialogContent>
                <DialogHeader>
                  <DialogTitle>Create New Tag</DialogTitle>
                  <DialogDescription>
                    Add a new tag to help classify your videos
                  </DialogDescription>
                </DialogHeader>
                <CreateTagForm
                  onSuccess={handleTagSuccess}
                  onCancel={() => setTagDialogOpen(false)}
                />
              </DialogContent>
            </Dialog>
          </div>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            {tags.length === 0 ? (
              <div className="text-center py-8">
                <Tag className="h-12 w-12 mx-auto mb-4 text-muted-foreground" />
                <p className="text-sm text-muted-foreground">No tags yet</p>
                <p className="text-xs text-muted-foreground">
                  Create your first tag to get started
                </p>
              </div>
            ) : (
              <div className="flex flex-wrap gap-2">
                {tags.map((tag) => (
                  <div key={tag.id} className="group relative">
                    <Badge variant="outline" className="pr-8">
                      {tag.name}
                    </Badge>
                    <div className="absolute right-1 top-1/2 -translate-y-1/2 opacity-0 group-hover:opacity-100 flex gap-1">
                      <Button variant="ghost" size="sm" className="h-4 w-4 p-0">
                        <Edit className="h-3 w-3" />
                      </Button>
                      <Button
                        variant="ghost"
                        size="sm"
                        className="h-4 w-4 p-0 text-destructive hover:text-destructive"
                      >
                        <Trash2 className="h-3 w-3" />
                      </Button>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
        </CardContent>
      </Card>

      {/* Content Analytics */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <BarChart3 className="h-5 w-5" />
            Content Analytics
          </CardTitle>
          <CardDescription>
            Overview of your content performance
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="grid grid-cols-2 gap-4">
              <div className="text-center p-4 bg-muted/50 rounded-lg">
                <div className="text-2xl font-bold text-primary">
                  {categories.length}
                </div>
                <div className="text-sm text-muted-foreground">Categories</div>
              </div>
              <div className="text-center p-4 bg-muted/50 rounded-lg">
                <div className="text-2xl font-bold text-primary">
                  {tags.length}
                </div>
                <div className="text-sm text-muted-foreground">Tags</div>
              </div>
            </div>

            <Separator />

            <div className="space-y-3">
              <h4 className="font-medium">Quick Actions</h4>
              <div className="grid grid-cols-1 gap-2">
                <Button variant="outline" size="sm" className="justify-start">
                  <Eye className="h-4 w-4 mr-2" />
                  View Analytics
                </Button>
                <Button variant="outline" size="sm" className="justify-start">
                  <Users className="h-4 w-4 mr-2" />
                  Audience Insights
                </Button>
                <Button variant="outline" size="sm" className="justify-start">
                  <Settings className="h-4 w-4 mr-2" />
                  Content Settings
                </Button>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
