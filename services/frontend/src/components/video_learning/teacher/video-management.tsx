'use client';

import { useState, useEffect } from 'react';
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
import { Input } from '@/ui/input';
import { Label } from '@/ui/label';
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
import { videoLearningService } from '@/services/video_learning/service';
import { cookies } from 'next/headers';
import type { VideoCategory, VideoTag } from '@/types/video_learning/models';
import { Skeleton } from '@/ui/skeleton';
import { toast } from 'sonner';

export function VideoManagement() {
  const [categories, setCategories] = useState<VideoCategory[]>([]);
  const [tags, setTags] = useState<VideoTag[]>([]);
  const [loading, setLoading] = useState(true);
  const [newCategoryName, setNewCategoryName] = useState('');
  const [newTagName, setNewTagName] = useState('');
  const [isCreatingCategory, setIsCreatingCategory] = useState(false);
  const [isCreatingTag, setIsCreatingTag] = useState(false);
  const [categoryDialogOpen, setCategoryDialogOpen] = useState(false);
  const [tagDialogOpen, setTagDialogOpen] = useState(false);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setLoading(true);
      const service = await videoLearningService(cookies());

      const [categoriesResult, tagsResult] = await Promise.all([
        service.getAllCategories(),
        service.getAllTags(),
      ]);

      if (categoriesResult.success) {
        setCategories(categoriesResult.data);
      }

      if (tagsResult.success) {
        setTags(tagsResult.data);
      }
    } catch (error) {
      console.error('Failed to load data:', error);
      toast.error('Failed to load categories and tags');
    } finally {
      setLoading(false);
    }
  };

  const handleCreateCategory = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!newCategoryName.trim()) {
      toast.error('Please enter a category name');
      return;
    }

    setIsCreatingCategory(true);
    try {
      const service = await videoLearningService(cookies());
      const result = await service.createCategory({
        name: newCategoryName.trim(),
      });

      if (result.success) {
        setCategories((prev) => [...prev, result.data]);
        setNewCategoryName('');
        setCategoryDialogOpen(false);
        toast.success('Category created successfully');
      } else {
        toast.error('Failed to create category');
      }
    } catch (error) {
      console.error('Failed to create category:', error);
      toast.error('An error occurred while creating category');
    } finally {
      setIsCreatingCategory(false);
    }
  };

  const handleCreateTag = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!newTagName.trim()) {
      toast.error('Please enter a tag name');
      return;
    }

    setIsCreatingTag(true);
    try {
      const service = await videoLearningService(cookies());
      const result = await service.createTag({
        name: newTagName.trim(),
      });

      if (result.success) {
        setTags((prev) => [...prev, result.data]);
        setNewTagName('');
        setTagDialogOpen(false);
        toast.success('Tag created successfully');
      } else {
        toast.error('Failed to create tag');
      }
    } catch (error) {
      console.error('Failed to create tag:', error);
      toast.error('An error occurred while creating tag');
    } finally {
      setIsCreatingTag(false);
    }
  };

  const getRandomUsageCount = () => Math.floor(Math.random() * 50) + 1;

  if (loading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {Array.from({ length: 4 }).map((_, i) => (
          <Skeleton key={i} className="h-64 w-full" />
        ))}
      </div>
    );
  }

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
                <form onSubmit={handleCreateCategory} className="space-y-4">
                  <div>
                    <Label htmlFor="category-name">Category Name</Label>
                    <Input
                      id="category-name"
                      value={newCategoryName}
                      onChange={(e) => setNewCategoryName(e.target.value)}
                      placeholder="Enter category name"
                      disabled={isCreatingCategory}
                    />
                  </div>
                  <div className="flex justify-end gap-2">
                    <Button
                      type="button"
                      variant="outline"
                      onClick={() => setCategoryDialogOpen(false)}
                      disabled={isCreatingCategory}
                    >
                      Cancel
                    </Button>
                    <Button type="submit" disabled={isCreatingCategory}>
                      {isCreatingCategory ? 'Creating...' : 'Create Category'}
                    </Button>
                  </div>
                </form>
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
                    <Badge variant="secondary" className="text-xs">
                      {getRandomUsageCount()} videos
                    </Badge>
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
                <form onSubmit={handleCreateTag} className="space-y-4">
                  <div>
                    <Label htmlFor="tag-name">Tag Name</Label>
                    <Input
                      id="tag-name"
                      value={newTagName}
                      onChange={(e) => setNewTagName(e.target.value)}
                      placeholder="Enter tag name"
                      disabled={isCreatingTag}
                    />
                  </div>
                  <div className="flex justify-end gap-2">
                    <Button
                      type="button"
                      variant="outline"
                      onClick={() => setTagDialogOpen(false)}
                      disabled={isCreatingTag}
                    >
                      Cancel
                    </Button>
                    <Button type="submit" disabled={isCreatingTag}>
                      {isCreatingTag ? 'Creating...' : 'Create Tag'}
                    </Button>
                  </div>
                </form>
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

      {/* Management Tools */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Settings className="h-5 w-5" />
            Management Tools
          </CardTitle>
          <CardDescription>
            Tools to help manage your video content
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="p-4 border rounded-lg">
              <h4 className="font-medium mb-2">Bulk Operations</h4>
              <p className="text-sm text-muted-foreground mb-3">
                Apply actions to multiple videos at once
              </p>
              <div className="flex gap-2">
                <Button variant="outline" size="sm">
                  Bulk Edit
                </Button>
                <Button variant="outline" size="sm">
                  Bulk Delete
                </Button>
              </div>
            </div>

            <div className="p-4 border rounded-lg">
              <h4 className="font-medium mb-2">Import/Export</h4>
              <p className="text-sm text-muted-foreground mb-3">
                Manage your content data
              </p>
              <div className="flex gap-2">
                <Button variant="outline" size="sm">
                  Export Data
                </Button>
                <Button variant="outline" size="sm">
                  Import Videos
                </Button>
              </div>
            </div>

            <div className="p-4 border rounded-lg">
              <h4 className="font-medium mb-2">Moderation</h4>
              <p className="text-sm text-muted-foreground mb-3">
                Review and moderate content
              </p>
              <div className="flex gap-2">
                <Button variant="outline" size="sm">
                  Review Queue
                </Button>
                <Button variant="outline" size="sm">
                  Reports
                </Button>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
