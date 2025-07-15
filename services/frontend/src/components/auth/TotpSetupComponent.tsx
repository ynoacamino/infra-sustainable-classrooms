'use client';

import { useState } from 'react';
import { QRCodeSVG } from 'qrcode.react';
import {
  CheckIcon,
  ClipboardIcon,
  DownloadIcon,
  ShieldCheckIcon,
  TriangleAlertIcon,
} from 'lucide-react';
import { Button } from '@/ui/button';
import { Link } from '@/ui/link';

interface TotpSetupComponentProps {
  totpUrl: string;
  backupCodes: string[];
  issuer: string;
}

export function TotpSetupComponent({
  totpUrl,
  backupCodes,
  issuer,
}: TotpSetupComponentProps) {
  const [urlCopied, setUrlCopied] = useState(false);
  const [codesCopied, setCodesCopied] = useState(false);

  const copyToClipboard = async (
    text: string,
    setCopied: (value: boolean) => void,
  ) => {
    try {
      await navigator.clipboard.writeText(text);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error('Failed to copy to clipboard:', err);
    }
  };

  const downloadBackupCodes = () => {
    const codesText = backupCodes.join('\n');
    const blob = new Blob(
      [
        `Backup Codes for ${issuer}\n` +
          `Generated on: ${new Date().toLocaleString()}\n\n` +
          `IMPORTANT: Save these codes in a secure location.\n` +
          `Each code can only be used once.\n\n` +
          `Backup Codes:\n` +
          codesText,
      ],
      { type: 'text/plain' },
    );

    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `backup-codes-${new Date().toISOString().split('T')[0]}.txt`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };

  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="text-center">
        <ShieldCheckIcon className="mx-auto h-12 w-12 text-green-500" />
        <h3 className="mt-4 text-lg font-medium text-gray-900">
          Almost Done! Set Up Your Authenticator App
        </h3>
        <p className="mt-2 text-sm text-gray-600">
          Scan the QR code or enter the setup key manually in your authenticator
          app
        </p>
      </div>

      {/* QR Code Section */}
      <div className="bg-gray-50 rounded-lg p-6">
        <div className="text-center space-y-4">
          <h4 className="text-md font-medium text-gray-900">
            Step 1: Scan QR Code
          </h4>

          <div className="flex justify-center">
            <div className="bg-white p-4 rounded-lg shadow-sm border">
              <QRCodeSVG
                value={totpUrl}
                size={200}
                level="M"
                className="mx-auto"
              />
            </div>
          </div>

          <p className="text-sm text-gray-600 max-w-md mx-auto">
            Open your authenticator app (Google Authenticator, Authy, 1Password,
            etc.) and scan this QR code to add your account.
          </p>
        </div>
      </div>

      {/* Manual Entry Section */}
      <div className="bg-blue-50 rounded-lg p-6">
        <h4 className="text-md font-medium text-gray-900 mb-4">
          Step 2: Or Enter Manually
        </h4>

        <div className="space-y-3">
          <div>
            <label className="block text-sm font-medium text-gray-700">
              Setup Key (if you can&apos;t scan the QR code)
            </label>
            <div className="mt-1 flex">
              <input
                type="text"
                readOnly
                value={totpUrl}
                className="flex-1 min-w-0 block w-full px-3 py-2 rounded-l-md border border-gray-300 bg-white text-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              />
              <button
                onClick={() => copyToClipboard(totpUrl, setUrlCopied)}
                className="inline-flex items-center px-3 py-2 border border-l-0 border-gray-300 rounded-r-md bg-gray-50 text-gray-500 text-sm hover:bg-gray-100 focus:outline-none focus:ring-1 focus:ring-indigo-500 focus:border-indigo-500"
              >
                {urlCopied ? (
                  <CheckIcon className="h-4 w-4 text-green-500" />
                ) : (
                  <ClipboardIcon className="h-4 w-4" />
                )}
              </button>
            </div>
          </div>

          <div className="text-xs text-gray-600 space-y-1">
            <p>
              <strong>Account:</strong>{' '}
              {new URL(totpUrl).searchParams.get('issuer')} (
              {new URL(totpUrl).pathname.split(':')[1]})
            </p>
            <p>
              <strong>Secret:</strong>{' '}
              {new URL(totpUrl).searchParams.get('secret')}
            </p>
          </div>
        </div>
      </div>

      {/* Backup Codes Section */}
      <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-6">
        <div className="flex items-start">
          <TriangleAlertIcon className="size-5 text-yellow-400 mt-0.5 mr-3 flex-shrink-0" />
          <div className="flex-1">
            <h4 className="text-md font-medium text-gray-900 mb-4">
              Step 3: Save Your Backup Codes
            </h4>

            <p className="text-sm text-gray-700 mb-4">
              These backup codes can be used to access your account if you lose
              your authenticator device.
              <strong className="text-red-600">
                {' '}
                Each code can only be used once.
              </strong>
            </p>

            <div className="bg-white rounded-md p-4 border">
              <div className="grid grid-cols-2 gap-2 text-sm font-mono">
                {backupCodes.map((code, index) => (
                  <div
                    key={index}
                    className="bg-gray-50 px-2 py-1 rounded text-center"
                  >
                    {code}
                  </div>
                ))}
              </div>
            </div>

            <div className="mt-4 flex space-x-3">
              <Button onClick={downloadBackupCodes}>
                <DownloadIcon className="h-4 w-4 mr-2" />
                Download Codes
              </Button>

              <Button
                onClick={() =>
                  copyToClipboard(backupCodes.join('\n'), setCodesCopied)
                }
                variant="secondary"
              >
                {codesCopied ? (
                  <CheckIcon className="h-4 w-4 mr-2 text-green-500" />
                ) : (
                  <ClipboardIcon className="h-4 w-4 mr-2" />
                )}
                Copy Codes
              </Button>
            </div>
          </div>
        </div>
      </div>

      {/* Instructions */}
      <div className="bg-gray-50 rounded-lg p-6">
        <h4 className="text-md font-medium text-gray-900 mb-4">
          What&apos;s Next?
        </h4>

        <ol className="list-decimal list-inside space-y-2 text-sm text-gray-700">
          <li>
            Add this account to your authenticator app using the QR code or
            setup key above
          </li>
          <li>
            Save your backup codes in a secure location (password manager,
            secure note, etc.)
          </li>
          <li>
            Test your setup by entering a code from your authenticator app
          </li>
          <li>Complete your registration process</li>
        </ol>

        <div className="mt-4 p-3 bg-blue-100 rounded-md">
          <p className="text-xs text-blue-800">
            <strong>Tip:</strong> Popular authenticator apps include Google
            Authenticator, Microsoft Authenticator, Authy, 1Password, Bitwarden,
            and LastPass Authenticator.
          </p>
        </div>
      </div>

      {/* Action Buttons */}
      <div className="flex justify-between pt-6 border-t">
        {/* TODO: Delete cookies except session token */}
        <Link href="/auth/register">Back to Registration</Link>
        <Link href="/auth/verify">Continue to Verification</Link>
      </div>
    </div>
  );
}
