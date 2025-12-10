import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export function readFile(filename) {
  const filePath = path.join(__dirname, '../go/files', filename);
  const content = fs.readFileSync(filePath, 'utf-8');
  return content.split('\n').map(line => line.trimEnd());
}
