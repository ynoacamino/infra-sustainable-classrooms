'use client';

import React, { useRef } from 'react';
import Jodit from 'jodit-react';

interface JoditEditorProps {
  placeholder?: string;
  content: string;
  setContent: (content: string) => void;
}

export default function JoditEditor({
  placeholder,
  content,
  setContent,
}: JoditEditorProps) {
  const editor = useRef(null);

  const config = {
    readonly: false,
    placeholder: placeholder || 'Start typings...',
    language: 'es',
  };

  return (
    <Jodit
      ref={editor}
      value={content}
      config={config}
      tabIndex={1}
      onBlur={(newContent) => setContent(newContent)}
    />
  );
}
