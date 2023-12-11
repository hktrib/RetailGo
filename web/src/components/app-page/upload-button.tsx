import { useCallback, useState } from "react";

import type { FileWithPath } from "@uploadthing/react";
import { useDropzone } from "@uploadthing/react/hooks";
import { generateClientDropzoneAccept } from "uploadthing/client";
import { useUploadThing } from "@/lib/uploadthing";
import { cx } from "class-variance-authority";
import { Upload } from "lucide-react";

export default function UploadButton({
  endpoint,
  className,
}: {
  endpoint: "storeItemImage";
  className?: string;
}) {
  const [files, setFiles] = useState<File[]>([]);

  const onDrop = useCallback((acceptedFiles: FileWithPath[]) => {
    setFiles([acceptedFiles[0]]);
  }, []);

  // runs on client
  const { startUpload, permittedFileInfo } = useUploadThing(endpoint, {
    onClientUploadComplete: () => {
      // onUploadComplete data here
      alert("uploaded successfully!");
    },
    onUploadError: () => {
      alert("error occurred while uploading");
    },
    onUploadBegin: () => {
      alert("upload has begun");
    },
  });

  const fileTypes = permittedFileInfo?.config
    ? Object.keys(permittedFileInfo?.config)
    : [];

  const { getRootProps, getInputProps } = useDropzone({
    onDrop,
    accept: fileTypes ? generateClientDropzoneAccept(fileTypes) : undefined,
  });

  return (
    <div
      {...getRootProps()}
      className={cx(
        className,
        "flex h-48 items-center justify-center rounded-md border-2 border-dotted dark:border-zinc-800",
      )}
    >
      <div>
        {files.length > 0 ? (
          <button
            onClick={() => startUpload(files)}
            className="text-sm dark:bg-sky-500 dark:text-white"
          >
            Upload file
          </button>
        ) : (
          <>
            <input {...getInputProps()} />
            <div className="flex items-center gap-x-2">
              <Upload className="h-4 w-4 dark:text-zinc-300" />
              <span className="text-sm dark:text-zinc-200">Upload image</span>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
