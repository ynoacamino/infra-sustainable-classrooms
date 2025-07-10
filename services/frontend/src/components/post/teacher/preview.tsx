import { marked } from 'marked';
import markedShiki from 'marked-shiki';
import '@/modules/post/lib/setupCopy.css';
import { useCallback, useEffect, useState } from 'react';
import { codeToHtml } from 'shiki';
import { setupCopy } from '@/lib/post/setupCopy';

export default function Preview({ content }: { content: string }) {
  const [parsedContent, setParsedContent] = useState<string>('');

  const parse = useCallback(async () => {
    const mark = await marked
      .setOptions({
        async: true,
      })
      .use(
        markedShiki({
          async highlight(code, lang) {
            return await codeToHtml(code, {
              lang,
              theme: 'github-dark',
            });
          },
          container: `<figure class="highlighted-code">
      <button class="btn-copy">Copy</button>
      %s
      </figure>`,
        }),
      )
      .parse(content);
    setParsedContent(mark);
  }, [content]);

  useEffect(() => {
    parse().catch((err) => console.log(err));
  }, [content, parse]);

  useEffect(() => {
    setupCopy();
  }, [parsedContent]);

  return (
    <div
      className="prose text-foreground w-full max-w-none"
      dangerouslySetInnerHTML={{ __html: parsedContent }}
    ></div>
  );
}
