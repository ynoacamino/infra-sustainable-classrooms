import ModeSwitch, {
  ModeSwitchSkeleton,
} from '@/components/post/admin/mode-switch';
import { useOnePost } from '@/hooks/post/use-posts';
import { useViewOrEdit } from '@/hooks/post/use-view-or-edit';
import { useEffect } from 'react';

export default function EditPost({ postId }: { postId: string }) {
  const { content, setContent, setViewOrEdit, viewOrEdit } = useViewOrEdit();
  const { isLoading, post } = useOnePost({ postId });

  useEffect(() => {
    if (!isLoading && post && post.content) {
      setContent(post.content);
    }
  }, [isLoading, post, setContent]);

  if (isLoading || !post) {
    return <ModeSwitchSkeleton />;
  }

  return (
    <ModeSwitch
      viewOrEdit={viewOrEdit}
      setViewOrEdit={setViewOrEdit}
      content={content}
      setContent={setContent}
    />
  );
}
